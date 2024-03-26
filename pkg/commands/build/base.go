package build

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/pterm/pterm"

	"github.com/888go/wails/internal/system"

	"github.com/leaanthony/gosod"
	"github.com/888go/wails/internal/frontend/runtime/wrapper"

	"github.com/pkg/errors"

	"github.com/leaanthony/slicer"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/project"
	"github.com/888go/wails/internal/shell"
	"github.com/888go/wails/pkg/clilogger"
)

const (
	VERBOSE int = 2
)

// BaseBuilder 是通用构建器结构体
type BaseBuilder struct {
	filesToDelete slicer.StringSlicer
	projectData   *project.Project
	options       *Options
}

// NewBaseBuilder 创建一个新的 BaseBuilder
func NewBaseBuilder(options *Options) *BaseBuilder {
	result := &BaseBuilder{
		options: options,
	}
	return result
}

// SetProjectData 为该构建器设置项目数据
func (b *BaseBuilder) SetProjectData(projectData *project.Project) {
	b.projectData = projectData
}

func (b *BaseBuilder) addFileToDelete(filename string) {
	if !b.options.KeepAssets {
		b.filesToDelete.Add(filename)
	}
}

func (b *BaseBuilder) fileExists(path string) bool {
	// 如果文件不存在，则忽略
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func (b *BaseBuilder) convertFileToIntegerString(filename string) (string, error) {
	rawData, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return b.convertByteSliceToIntegerString(rawData), nil
}

func (b *BaseBuilder) convertByteSliceToIntegerString(data []byte) string {
	// Create string builder
	var result strings.Builder

	if len(data) > 0 {

		// 遍历除最后1个字节外的所有字节
		for i := 0; i < len(data)-1; i++ {
			result.WriteString(fmt.Sprintf("%v,", data[i]))
		}

		result.WriteString(strconv.FormatUint(uint64(data[len(data)-1]), 10))
	}

	return result.String()
}

// CleanUp 进行构建后清理工作
func (b *BaseBuilder) CleanUp() {
	// Delete all the files
	b.filesToDelete.Each(func(filename string) {
		// 如果文件不存在，则忽略
		if !b.fileExists(filename) {
			return
		}

// 删除文件。我们忽略错误，因为这些文件无论如何将在下次构建时被覆盖。
		_ = os.Remove(filename)
	})
}

func commandPrettifier(args []string) string {
	// 如果我们有一个单一的参数，直接返回它
	if len(args) == 1 {
		return args[0]
	}
	// 如果参数中包含空格，则对该参数进行引号引用
	for i, arg := range args {
		if strings.Contains(arg, " ") {
			args[i] = fmt.Sprintf("\"%s\"", arg)
		}
	}
	return strings.Join(args, " ")
}

func (b *BaseBuilder) OutputFilename(options *Options) string {
	outputFile := options.OutputFile
	if outputFile == "" {
		target := strings.TrimSuffix(b.projectData.OutputFilename, ".exe")
		if b.projectData.OutputType != "desktop" {
			target += "-" + b.projectData.OutputType
		}
		// 如果我们没有使用标准编译器，将其添加到文件名中
		if options.Compiler != "go" {
			// 解析`go version`命令的输出结果。例如：`go version go1.16 windows/amd64`
			stdout, _, err := shell.RunCommand(".", options.Compiler, "version")
			if err != nil {
				return ""
			}
			versionSplit := strings.Split(stdout, " ")
			if len(versionSplit) == 4 {
				target += "-" + versionSplit[2]
			}
		}
		switch b.options.Platform {
		case "windows":
			outputFile = target + ".exe"
		case "darwin", "linux":
			if b.options.Arch == "" {
				b.options.Arch = runtime.GOARCH
			}
			outputFile = fmt.Sprintf("%s-%s-%s", target, b.options.Platform, b.options.Arch)
		}

	}
	return outputFile
}

// CompileProject 编译项目
func (b *BaseBuilder) CompileProject(options *Options) error {
	// 检查运行时包装器是否存在
	err := generateRuntimeWrapper(options)
	if err != nil {
		return err
	}

	verbose := options.Verbosity == VERBOSE
	// Run go mod tidy first
	if !options.SkipModTidy {
		cmd := exec.Command(options.Compiler, "mod", "tidy")
		cmd.Stderr = os.Stderr
		if verbose {
			println("")
			cmd.Stdout = os.Stdout
		}
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	commands := slicer.String()

	compiler := options.Compiler
	if options.Obfuscated {
		if !shell.CommandExists("garble") {
			return fmt.Errorf("the 'garble' command was not found. Please install it with `go install mvdan.cc/garble@latest`")
		} else {
			compiler = "garble"
			if options.GarbleArgs != "" {
				commands.AddSlice(strings.Split(options.GarbleArgs, " "))
			}
			options.UserTags = append(options.UserTags, "obfuscated")
		}
	}

	// 默认的Go构建命令
	commands.Add("build")

	// 添加更好的调试标志
	if options.Mode == Dev || options.Mode == Debug {
		commands.Add("-gcflags")
		commands.Add("all=-N -l")
	}

	if options.ForceBuild {
		commands.Add("-a")
	}

	if options.TrimPath {
		commands.Add("-trimpath")
	}

	if options.RaceDetector {
		commands.Add("-race")
	}

	var tags slicer.StringSlicer
	tags.Add(options.OutputType)
	tags.AddSlice(options.UserTags)

	// 如果我们有webview2策略，则添加它
	if options.WebView2Strategy != "" {
		tags.Add(options.WebView2Strategy)
	}

	if options.Mode == Production || options.Mode == Debug {
		tags.Add("production")
	}
	// 此模式允许您调试生产构建（非开发构建）
	if options.Mode == Debug {
		tags.Add("debug")
	}

	// 这个选项允许你在生产构建中启用开发者工具（在开发构建中它始终是启用的，所以无需在此设置）
	if options.Devtools {
		tags.Add("devtools")
	}

	if options.Obfuscated {
		tags.Add("obfuscated")
	}

	tags.Deduplicate()

	// 添加输出类型构建标签
	commands.Add("-tags")
	commands.Add(tags.Join(","))

	// LDFlags
	ldflags := slicer.String()
	if options.LDFlags != "" {
		ldflags.Add(options.LDFlags)
	}

	if options.Mode == Production {
		ldflags.Add("-w", "-s")
		if options.Platform == "windows" && !options.WindowsConsole {
			ldflags.Add("-H windowsgui")
		}
	}

	ldflags.Deduplicate()

	if ldflags.Length() > 0 {
		commands.Add("-ldflags")
		commands.Add(ldflags.Join(" "))
	}

	// 获取应用程序构建目录
	appDir := options.BinDirectory
	if options.CleanBinDirectory {
		err = cleanBinDirectory(options)
		if err != nil {
			return err
		}
	}

	// Set up output filename
	outputFile := b.OutputFilename(options)
	compiledBinary := filepath.Join(appDir, outputFile)
	commands.Add("-o")
	commands.Add(compiledBinary)

	options.CompiledBinary = compiledBinary

	// Build the application
	cmd := exec.Command(compiler, commands.AsSlice()...)
	cmd.Stderr = os.Stderr
	if verbose {
		pterm.Info.Println("Build command:", compiler, commandPrettifier(commands.AsSlice()))
		cmd.Stdout = os.Stdout
	}
	// Set the directory
	cmd.Dir = b.projectData.Path

// 添加CGO标志
// TODO: 移除这一行，因为我们现在已经不再生成头文件了
// 我们使用项目/构建目录作为临时位置来存放我们生成的C语言头文件
	buildBaseDir, err := fs.RelativeToCwd("build")
	if err != nil {
		return err
	}

	cmd.Env = os.Environ() // inherit env

	if options.Platform != "windows" {
		// 使用shell.UpsertEnv，以免覆盖用户自定义的CGO_CFLAGS环境变量
		cmd.Env = shell.UpsertEnv(cmd.Env, "CGO_CFLAGS", func(v string) string {
			if options.Platform == "darwin" {
				if v != "" {
					v += " "
				}
				v += "-mmacosx-version-min=10.13"
			}
			return v
		})
		// 使用shell.UpsertEnv，这样我们不会覆盖用户的CGO_CXXFLAGS环境变量
		cmd.Env = shell.UpsertEnv(cmd.Env, "CGO_CXXFLAGS", func(v string) string {
			if v != "" {
				v += " "
			}
			v += "-I" + buildBaseDir
			return v
		})

		cmd.Env = shell.UpsertEnv(cmd.Env, "CGO_ENABLED", func(v string) string {
			return "1"
		})
		if options.Platform == "darwin" {
// 确定版本以便链接到更新的框架
// 为什么CGO没有这个选项呢？！？！
			info, err := system.GetInfo()
			if err != nil {
				return err
			}
			versionSplit := strings.Split(info.OS.Version, ".")
			majorVersion, err := strconv.Atoi(versionSplit[0])
			if err != nil {
				return err
			}
			addUTIFramework := majorVersion >= 11
			// 设置Mac SDK的最低版本为10.13
			cmd.Env = shell.UpsertEnv(cmd.Env, "CGO_LDFLAGS", func(v string) string {
				if v != "" {
					v += " "
				}
				if addUTIFramework {
					v += "-framework UniformTypeIdentifiers "
				}
				v += "-mmacosx-version-min=10.13"

				return v
			})
		}
	}

	cmd.Env = shell.UpsertEnv(cmd.Env, "GOOS", func(v string) string {
		return options.Platform
	})

	cmd.Env = shell.UpsertEnv(cmd.Env, "GOARCH", func(v string) string {
		return options.Arch
	})

	if verbose {
		printBulletPoint("Environment:", strings.Join(cmd.Env, " "))
	}

	// Run command
	err = cmd.Run()
	cmd.Stderr = os.Stderr

	// 如果我们有错误，则格式化该错误
	if err != nil {
		if options.Platform == "darwin" {
			output, _ := cmd.CombinedOutput()
			stdErr := string(output)
			if strings.Contains(err.Error(), "ld: framework not found UniformTypeIdentifiers") ||
				strings.Contains(stdErr, "ld: framework not found UniformTypeIdentifiers") {
				pterm.Warning.Println(`
NOTE: It would appear that you do not have the latest Xcode cli tools installed.
Please reinstall by doing the following:
  1. Remove the current installation located at "xcode-select -p", EG: sudo rm -rf /Library/Developer/CommandLineTools
  2. Install latest Xcode tools: xcode-select --install`)
			}
		}
		return err
	}

	if !options.Compress {
		return nil
	}

	printBulletPoint("Compressing application: ")

	// 我们是否安装了 upx？
	if !shell.CommandExists("upx") {
		pterm.Warning.Println("Warning: Cannot compress binary: upx not found")
		return nil
	}

	args := []string{"--best", "--no-color", "--no-progress", options.CompiledBinary}

	if options.CompressFlags != "" {
		args = strings.Split(options.CompressFlags, " ")
		args = append(args, options.CompiledBinary)
	}

	if verbose {
		pterm.Info.Println("upx", strings.Join(args, " "))
	}

	output, err := exec.Command("upx", args...).Output()
	if err != nil {
		return errors.Wrap(err, "Error during compression:")
	}
	pterm.Println("Done.")
	if verbose {
		pterm.Info.Println(string(output))
	}

	return nil
}

func generateRuntimeWrapper(options *Options) error {
	if options.WailsJSDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		options.WailsJSDir = filepath.Join(cwd, "frontend")
	}
	wrapperDir := filepath.Join(options.WailsJSDir, "wailsjs", "runtime")
	_ = os.RemoveAll(wrapperDir)
	extractor := gosod.New(wrapper.RuntimeWrapper)
	err := extractor.Extract(wrapperDir, nil)
	if err != nil {
		return err
	}

	return nil
}

// NpmInstall 在给定的目录中运行 "npm install"
func (b *BaseBuilder) NpmInstall(sourceDir string, verbose bool) error {
	return b.NpmInstallUsingCommand(sourceDir, "npm install", verbose)
}

// NpmInstallUsingCommand 在指定的npm项目目录中运行给定的安装命令
func (b *BaseBuilder) NpmInstallUsingCommand(sourceDir string, installCommand string, verbose bool) error {
	packageJSON := filepath.Join(sourceDir, "package.json")

	// 检查 package.json 是否存在
	if !fs.FileExists(packageJSON) {
		// 没有package.json，无需安装
		return nil
	}

	install := false

	// 获取package.json的MD5校验和
	packageJSONMD5 := fs.MustMD5File(packageJSON)

	// 检查是否需要执行npm install
	packageChecksumFile := filepath.Join(sourceDir, "package.json.md5")
	if fs.FileExists(packageChecksumFile) {
		// Compare checksums
		storedChecksum := fs.MustLoadString(packageChecksumFile)
		if storedChecksum != packageJSONMD5 {
			fs.MustWriteString(packageChecksumFile, packageJSONMD5)
			install = true
		}
	} else {
		install = true
		fs.MustWriteString(packageChecksumFile, packageJSONMD5)
	}

	// 如果node_modules不存在，则进行安装
	nodeModulesDir := filepath.Join(sourceDir, "node_modules")
	if !fs.DirExists(nodeModulesDir) {
		install = true
	}

	// 检查是否为强制安装
	if b.options.ForceBuild {
		install = true
	}

	// Shortcut installation
	if !install {
		if verbose {
			pterm.Println("Skipping npm install")
		}
		return nil
	}

	// 将InstallCommand拆分并执行
	cmd := strings.Split(installCommand, " ")
	stdout, stderr, err := shell.RunCommand(sourceDir, cmd[0], cmd[1:]...)
	if verbose || err != nil {
		for _, l := range strings.Split(stdout, "\n") {
			pterm.Printf("    %s\n", l)
		}
		for _, l := range strings.Split(stderr, "\n") {
			pterm.Printf("    %s\n", l)
		}
	}

	return err
}

// NpmRun在指定的目录中执行npm目标
func (b *BaseBuilder) NpmRun(projectDir, buildTarget string, verbose bool) error {
	stdout, stderr, err := shell.RunCommand(projectDir, "npm", "run", buildTarget)
	if verbose || err != nil {
		for _, l := range strings.Split(stdout, "\n") {
			pterm.Printf("    %s\n", l)
		}
		for _, l := range strings.Split(stderr, "\n") {
			pterm.Printf("    %s\n", l)
		}
	}
	return err
}

// NpmRunWithEnvironment 在指定的目录下，使用给定的环境变量执行npm目标
func (b *BaseBuilder) NpmRunWithEnvironment(projectDir, buildTarget string, verbose bool, envvars []string) error {
	cmd := shell.CreateCommand(projectDir, "npm", "run", buildTarget)
	cmd.Env = append(os.Environ(), envvars...)
	var stdo, stde bytes.Buffer
	cmd.Stdout = &stdo
	cmd.Stderr = &stde
	err := cmd.Run()
	if verbose || err != nil {
		for _, l := range strings.Split(stdo.String(), "\n") {
			pterm.Printf("    %s\n", l)
		}
		for _, l := range strings.Split(stde.String(), "\n") {
			pterm.Printf("    %s\n", l)
		}
	}
	return err
}

// BuildFrontend 执行针对前端目录的 `npm build` 命令
func (b *BaseBuilder) BuildFrontend(outputLogger *clilogger.CLILogger) error {
	verbose := b.options.Verbosity == VERBOSE

	frontendDir := b.projectData.GetFrontendDir()
	if !fs.DirExists(frontendDir) {
		return fmt.Errorf("frontend directory '%s' does not exist", frontendDir)
	}

	// 检查 wails.json 中是否提供了 'InstallCommand'
	installCommand := b.projectData.InstallCommand
	if b.projectData.OutputType == "dev" {
		installCommand = b.projectData.GetDevInstallerCommand()
	}
	if installCommand == "" {
		// No - don't install
		printBulletPoint("No Install command. Skipping.")
		pterm.Println("")
	} else {
		// Do install if needed
		printBulletPoint("Installing frontend dependencies: ")
		if verbose {
			pterm.Println("")
			pterm.Info.Println("Install command: '" + installCommand + "'")
		}
		if err := b.NpmInstallUsingCommand(frontendDir, installCommand, verbose); err != nil {
			return err
		}
		outputLogger.X日志输出并换行("Done.")
	}

	// 检查是否存在构建命令
	buildCommand := b.projectData.BuildCommand
	if b.projectData.OutputType == "dev" {
		buildCommand = b.projectData.GetDevBuildCommand()
	}
	if buildCommand == "" {
		printBulletPoint("No Build command. Skipping.")
		pterm.Println("")
		// No - ignore
		return nil
	}

	printBulletPoint("Compiling frontend: ")
	cmd := strings.Split(buildCommand, " ")
	if verbose {
		pterm.Println("")
		pterm.Info.Println("Build command: '" + buildCommand + "'")
	}
	stdout, stderr, err := shell.RunCommand(frontendDir, cmd[0], cmd[1:]...)
	if verbose || err != nil {
		for _, l := range strings.Split(stdout, "\n") {
			pterm.Printf("    %s\n", l)
		}
		for _, l := range strings.Split(stderr, "\n") {
			pterm.Printf("    %s\n", l)
		}
	}
	if err != nil {
		return err
	}

	pterm.Println("Done.")
	return nil
}
