package dev

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/samber/lo"
	"github.com/888go/wails/cmd/wails/flags"
	"github.com/888go/wails/cmd/wails/internal/gomod"
	"github.com/888go/wails/cmd/wails/internal/logutils"
	"golang.org/x/mod/semver"

	"github.com/888go/wails/pkg/commands/buildtags"

	"github.com/google/shlex"

	"github.com/pkg/browser"

	"github.com/fsnotify/fsnotify"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/process"
	"github.com/888go/wails/pkg/clilogger"
	"github.com/888go/wails/pkg/commands/build"
)

const (
	viteMinVersion = "v3.0.0"
)

func sliceToMap(input []string) map[string]struct{} {
	result := map[string]struct{}{}
	for _, value := range input {
		result[value] = struct{}{}
	}
	return result
}

// Application 在开发模式下运行应用程序

// ff:
// logger:
// f:
func Application(f *flags.Dev, logger *clilogger.CLILogger) error {
	cwd := lo.Must(os.Getwd())

	// 更新go.mod文件以使用当前wails版本
	err := gomod.SyncGoMod(logger, !f.NoSyncGoMod)
	if err != nil {
		return err
	}

	if !f.SkipModTidy {
		// 运行 go mod tidy 以确保我们的依赖是最新的
		err = runCommand(cwd, false, f.Compiler, "mod", "tidy")
		if err != nil {
			return err
		}
	}

	buildOptions := f.GenerateBuildOptions()
	buildOptions.Logger = logger

	userTags, err := buildtags.X解析(f.Tags)
	if err != nil {
		return err
	}

	buildOptions.UserTags = userTags

	projectConfig := f.ProjectConfig()

	// Setup signal handler
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM)
	exitCodeChannel := make(chan int, 1)

	// 如果有请求，则构建前端，但忽略构建应用程序本身。
	ignoreFrontend := buildOptions.IgnoreFrontend
	if !ignoreFrontend {
		buildOptions.IgnoreApplication = true
		if _, err := build.X构建项目(buildOptions); err != nil {
			return err
		}
		buildOptions.IgnoreApplication = false
	}

	legacyUseDevServerInsteadofCustomScheme := false
	// 前端：开发：监视器命令。
	frontendDevAutoDiscovery := projectConfig.IsFrontendDevServerURLAutoDiscovery()
	if command := projectConfig.DevWatcherCommand; command != "" {
		closer, devServerURL, devServerViteVersion, err := runFrontendDevWatcherCommand(projectConfig.GetFrontendDir(), command, frontendDevAutoDiscovery)
		if err != nil {
			return err
		}
		if devServerURL != "" {
			projectConfig.FrontendDevServerURL = devServerURL
			f.FrontendDevServerURL = devServerURL
		}
		defer closer()

		if devServerViteVersion != "" && semver.Compare(devServerViteVersion, viteMinVersion) < 0 {
			logutils.LogRed("Please upgrade your Vite Server to at least '%s' future Wails versions will require at least Vite '%s'", viteMinVersion, viteMinVersion)
			time.Sleep(3 * time.Second)
			legacyUseDevServerInsteadofCustomScheme = true
		}
	} else if frontendDevAutoDiscovery {
		return fmt.Errorf("unable to auto discover frontend:dev:serverUrl without a frontend:dev:watcher command, please either set frontend:dev:watcher or remove the auto discovery from frontend:dev:serverUrl")
	}

	// 只针对应用程序进行首次构建
	logger.X日志输出并换行("Building application for development...")
	buildOptions.IgnoreFrontend = true
	debugBinaryProcess, appBinary, err := restartApp(buildOptions, nil, f, exitCodeChannel, legacyUseDevServerInsteadofCustomScheme)
	buildOptions.IgnoreFrontend = ignoreFrontend || f.FrontendDevServerURL != ""
	if err != nil {
		return err
	}
	defer func() {
		if err := killProcessAndCleanupBinary(debugBinaryProcess, appBinary); err != nil {
			logutils.LogDarkYellow("Unable to kill process and cleanup binary: %s", err)
		}
	}()

	// open browser
	if f.Browser {
		err = browser.OpenURL(f.DevServerURL().String())
		if err != nil {
			return err
		}
	}

	logutils.LogGreen("Using DevServer URL: %s", f.DevServerURL())
	if f.FrontendDevServerURL != "" {
		logutils.LogGreen("Using Frontend DevServer URL: %s", f.FrontendDevServerURL)
	}
	logutils.LogGreen("Using reload debounce setting of %d milliseconds", f.Debounce)

	// 在终端中显示开发服务器URL，3秒后
	go func() {
		time.Sleep(3 * time.Second)
		logutils.LogGreen("\n\nTo develop in the browser and call your bound Go methods from Javascript, navigate to: %s", f.DevServerURL())
	}()

	// 监听变更并触发 restartApp() 函数
	debugBinaryProcess, err = doWatcherLoop(cwd, buildOptions, debugBinaryProcess, f, exitCodeChannel, quitChannel, f.DevServerURL(), legacyUseDevServerInsteadofCustomScheme)
	if err != nil {
		return err
	}

	// 如果当前程序正在运行，则终止程序并移除开发版二进制文件
	if err := killProcessAndCleanupBinary(debugBinaryProcess, appBinary); err != nil {
		return err
	}

	// 重置进程和二进制文件，以便defer语句能够识别并将其视为空操作（nop）。
	debugBinaryProcess = nil
	appBinary = ""

	logutils.LogGreen("Development mode exited")

	return nil
}

func killProcessAndCleanupBinary(process *process.Process, binary string) error {
	if process != nil && process.Running {
		if err := process.Kill(); err != nil {
			return err
		}
	}

	if binary != "" {
		err := os.Remove(binary)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	return nil
}

func runCommand(dir string, exitOnError bool, command string, args ...string) error {
	logutils.LogGreen("Executing: " + command + " " + strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(string(output))
		println(err.Error())
		if exitOnError {
			os.Exit(1)
		}
		return err
	}
	return nil
}

// runFrontendDevWatcherCommand 将会在接收到相应命令时执行 `frontend:dev:watcher` 命令，例如：`npm run dev`
func runFrontendDevWatcherCommand(frontendDirectory string, devCommand string, discoverViteServerURL bool) (func(), string, string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	scanner := NewStdoutScanner()
	cmdSlice := strings.Split(devCommand, " ")
	cmd := exec.CommandContext(ctx, cmdSlice[0], cmdSlice[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = scanner
	cmd.Dir = frontendDirectory
	setParentGID(cmd)

	if err := cmd.Start(); err != nil {
		cancel()
		return nil, "", "", fmt.Errorf("unable to start frontend DevWatcher: %w", err)
	}

	var viteServerURL string
	if discoverViteServerURL {
		select {
		case serverURL := <-scanner.ViteServerURLChan:
			viteServerURL = serverURL
		case <-time.After(time.Second * 10):
			cancel()
			return nil, "", "", errors.New("failed to find Vite server URL")
		}
	}

	viteVersion := ""
	select {
	case version := <-scanner.ViteServerVersionC:
		viteVersion = version

	case <-time.After(time.Second * 5):
		// 那么，这很可能不是vite在运行
	}

	logutils.LogGreen("Running frontend DevWatcher command: '%s'", devCommand)
	var wg sync.WaitGroup
	wg.Add(1)

	const (
		stateRunning   int32 = 0
		stateCanceling int32 = 1
		stateStopped   int32 = 2
	)
	state := stateRunning
	go func() {
		if err := cmd.Wait(); err != nil {
			wasRunning := atomic.CompareAndSwapInt32(&state, stateRunning, stateStopped)
			if err.Error() != "exit status 1" && wasRunning {
				logutils.LogRed("Error from DevWatcher '%s': %s", devCommand, err.Error())
			}
		}
		atomic.StoreInt32(&state, stateStopped)
		wg.Done()
	}()

	return func() {
		if atomic.CompareAndSwapInt32(&state, stateRunning, stateCanceling) {
			killProc(cmd, devCommand)
		}
		cancel()
		wg.Wait()
	}, viteServerURL, viteVersion, nil
}

// restartApp 当文件发生更改时，执行应用程序的实际重建工作
func restartApp(buildOptions *build.Options, debugBinaryProcess *process.Process, f *flags.Dev, exitCodeChannel chan int, legacyUseDevServerInsteadofCustomScheme bool) (*process.Process, string, error) {
	appBinary, err := build.X构建项目(buildOptions)
	println()
	if err != nil {
		logutils.LogRed("Build error - " + err.Error())

		msg := "Continuing to run current version"
		if debugBinaryProcess == nil {
			msg = "No version running, build will be retriggered as soon as changes have been detected"
		}
		logutils.LogDarkYellow(msg)
		return nil, "", nil
	}

	// 如果需要的话，杀死已存在的二进制文件
	if debugBinaryProcess != nil {
		killError := debugBinaryProcess.Kill()

		if killError != nil {
			buildOptions.Logger.X日志输出并停止("Unable to kill debug binary (PID: %d)!", debugBinaryProcess.PID())
		}

		debugBinaryProcess = nil
	}

	// parse appargs if any
	args, err := shlex.Split(f.AppArgs)
	if err != nil {
		buildOptions.Logger.X日志输出并停止("Unable to parse appargs: %s", err.Error())
	}

	// 根据实际情况设置环境变量
	os.Setenv("loglevel", f.LogLevel)
	os.Setenv("assetdir", f.AssetDir)
	os.Setenv("devserver", f.DevServer)
	os.Setenv("frontenddevserverurl", f.FrontendDevServerURL)

	// 使用正确的参数启动新的二进制文件
	newProcess := process.NewProcess(appBinary, args...)
	err = newProcess.Start(exitCodeChannel)
	if err != nil {
		// Remove binary
		if fs.FileExists(appBinary) {
			deleteError := fs.DeleteFile(appBinary)
			if deleteError != nil {
				buildOptions.Logger.X日志输出并停止("Unable to delete app binary: " + appBinary)
			}
		}
		buildOptions.Logger.X日志输出并停止("Unable to start application: %s", err.Error())
	}

	return newProcess, appBinary, nil
}

// doWatcherLoop 是主监视循环，在dev处于活动状态时运行
func doWatcherLoop(cwd string, buildOptions *build.Options, debugBinaryProcess *process.Process, f *flags.Dev, exitCodeChannel chan int, quitChannel chan os.Signal, devServerURL *url.URL, legacyUseDevServerInsteadofCustomScheme bool) (*process.Process, error) {
	// 创建项目文件观察器
	watcher, err := initialiseWatcher(cwd)
	if err != nil {
		logutils.LogRed("Unable to create filesystem watcher. Reloads will not occur.")
		return nil, err
	}

	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(watcher)

	logutils.LogGreen("Watching (sub)/directory: %s", cwd)

	// Main Loop
	extensionsThatTriggerARebuild := sliceToMap(strings.Split(f.Extensions, ","))
	var dirsThatTriggerAReload []string
	for _, dir := range strings.Split(f.ReloadDirs, ",") {
		if dir == "" {
			continue
		}
		thePath, err := filepath.Abs(dir)
		if err != nil {
			logutils.LogRed("Unable to expand reloadDir '%s': %s", dir, err)
			continue
		}
		dirsThatTriggerAReload = append(dirsThatTriggerAReload, thePath)
		err = watcher.Add(thePath)
		if err != nil {
			logutils.LogRed("Unable to watch path: %s due to error %v", thePath, err)
		} else {
			logutils.LogGreen("Watching (sub)/directory: %s", thePath)
		}
	}

	quit := false
	interval := time.Duration(f.Debounce) * time.Millisecond
	timer := time.NewTimer(interval)
	rebuild := false
	reload := false
	assetDir := ""
	changedPaths := map[string]struct{}{}

	// 如果我们正在使用外部开发服务器，前端部分的重新加载可以被跳过，或者如果用户请求这样做
	skipAssetsReload := f.FrontendDevServerURL != "" || f.NoReload

	assetDirURL := joinPath(devServerURL, "/wails/assetdir")
	reloadURL := joinPath(devServerURL, "/wails/reload")
	for !quit {
		// reload := false
		select {
		case exitCode := <-exitCodeChannel:
			if exitCode == 0 {
				quit = true
			}
		case err := <-watcher.Errors:
			logutils.LogDarkYellow(err.Error())
		case item := <-watcher.Events:
			isEligibleFile := func(fileName string) bool {
				// 遍历所有文件模式
				ext := filepath.Ext(fileName)
				if ext != "" {
					ext = ext[1:]
					if _, exists := extensionsThatTriggerARebuild[ext]; exists {
						return true
					}
				}
				return false
			}

			// 处理写操作
			if item.Op&fsnotify.Write == fsnotify.Write {
				// Ignore directories
				itemName := item.Name
				if fs.DirExists(itemName) {
					continue
				}

				if isEligibleFile(itemName) {
					rebuild = true
					timer.Reset(interval)
					continue
				}

				for _, reloadDir := range dirsThatTriggerAReload {
					if strings.HasPrefix(itemName, reloadDir) {
						reload = true
						break
					}
				}

				if !reload {
					changedPaths[filepath.Dir(itemName)] = struct{}{}
				}

				timer.Reset(interval)
			}

			// 处理新创建的文件系统条目
			if item.Op&fsnotify.Create == fsnotify.Create {
				// 如果这是一个文件夹，将其添加到我们的监视列表中
				if fs.DirExists(item.Name) {
					// node_modules 是被禁止的！
					if !strings.Contains(item.Name, "node_modules") {
						err := watcher.Add(item.Name)
						if err != nil {
							buildOptions.Logger.X日志输出并停止("%s", err.Error())
						}
						logutils.LogGreen("Added new directory to watcher: %s", item.Name)
					}
				} else if isEligibleFile(item.Name) {
// 处理新文件的创建。
// 注意：在某些平台上，对文件的更新表现为 REMOVE（删除）-> CREATE（创建），而非 WRITE（写入），因此这不仅包括新建文件，
// 还包括对现有文件的更新操作。
					rebuild = true
					timer.Reset(interval)
					continue
				}
			}
		case <-timer.C:
			if rebuild {
				rebuild = false
				if f.NoGoRebuild {
					logutils.LogGreen("[Rebuild triggered] skipping due to flag -nogorebuild")
				} else {
					logutils.LogGreen("[Rebuild triggered] files updated")
					// Try and build the app

					newBinaryProcess, _, err := restartApp(buildOptions, debugBinaryProcess, f, exitCodeChannel, legacyUseDevServerInsteadofCustomScheme)
					if err != nil {
						logutils.LogRed("Error during build: %s", err.Error())
						continue
					}
					// 如果我们有一个新的进程，保存其配置
					if newBinaryProcess != nil {
						debugBinaryProcess = newBinaryProcess
					}
				}
			}

			if !skipAssetsReload && len(changedPaths) != 0 {
				if assetDir == "" {
					resp, err := http.Get(assetDirURL)
					if err != nil {
						logutils.LogRed("Error during retrieving assetdir: %s", err.Error())
					} else {
						content, err := io.ReadAll(resp.Body)
						if err != nil {
							logutils.LogRed("Error reading assetdir from devserver: %s", err.Error())
						} else {
							assetDir = string(content)
						}
						resp.Body.Close()
					}
				}

				if assetDir != "" {
					for thePath := range changedPaths {
						if strings.HasPrefix(thePath, assetDir) {
							reload = true
							break
						}
					}
				} else if len(dirsThatTriggerAReload) == 0 {
					logutils.LogRed("Reloading couldn't be triggered: Please specify -assetdir or -reloaddirs")
				}
			}
			if reload {
				reload = false
				_, err := http.Get(reloadURL)
				if err != nil {
					logutils.LogRed("Error during refresh: %s", err.Error())
				}
			}
			changedPaths = map[string]struct{}{}
		case <-quitChannel:
			logutils.LogGreen("\nCaught quit")
			quit = true
		}
	}
	return debugBinaryProcess, nil
}

func joinPath(url *url.URL, subPath string) string {
	u := *url
	u.Path = path.Join(u.Path, subPath)
	return u.String()
}
