//go:build dev

package app

import (
	"context"
	"embed"
	"flag"
	"fmt"
	iofs "io/fs"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/888go/wails/pkg/assetserver"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend/desktop"
	"github.com/888go/wails/internal/frontend/devserver"
	"github.com/888go/wails/internal/frontend/dispatcher"
	"github.com/888go/wails/internal/frontend/runtime"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/menumanager"
	pkglogger "github.com/888go/wails/pkg/logger"
	"github.com/888go/wails/pkg/options"
)


// ff:
func (a *App) Run() error {
	err := a.frontend.Run(a.ctx)
	a.frontend.RunMainLoop()
	a.frontend.WindowClose()
	if a.shutdownCallback != nil {
		a.shutdownCallback(a.ctx)
	}
	return err
}

// CreateApp 创建应用！

// ff:
// appoptions:
func CreateApp(appoptions *options.App) (*App, error) {
	var err error

	ctx := context.Background()
	ctx = context.WithValue(ctx, "debug", true)
	ctx = context.WithValue(ctx, "devtoolsEnabled", true)

	// Set up logger
	myLogger := logger.New(appoptions.X日志记录器)
	myLogger.SetLogLevel(appoptions.X日志级别)

	// Check for CLI Flags
	devFlags := flag.NewFlagSet("dev", flag.ContinueOnError)

	var assetdirFlag *string
	var devServerFlag *string
	var frontendDevServerURLFlag *string
	var loglevelFlag *string

	assetdir := os.Getenv("assetdir")
	if assetdir == "" {
		assetdirFlag = devFlags.String("assetdir", "", "Directory to serve assets")
	}

	devServer := os.Getenv("devserver")
	if devServer == "" {
		devServerFlag = devFlags.String("devserver", "", "Address to bind the wails dev server to")
	}

	frontendDevServerURL := os.Getenv("frontenddevserverurl")
	if frontendDevServerURL == "" {
		frontendDevServerURLFlag = devFlags.String("frontenddevserverurl", "", "URL of the external frontend dev server")
	}

	loglevel := os.Getenv("loglevel")
	if loglevel == "" {
		loglevelFlag = devFlags.String("loglevel", "debug", "Loglevel to use - Trace, Debug, Info, Warning, Error")
	}

	// 如果我们没有在环境变量中获得资产目录（assetdir）
	if assetdir == "" {
		// 解析参数，但如果使用-appargs传递应用参数，则忽略错误。
		_ = devFlags.Parse(os.Args[1:])
		if assetdirFlag != nil {
			assetdir = *assetdirFlag
		}
		if devServerFlag != nil {
			devServer = *devServerFlag
		}
		if frontendDevServerURLFlag != nil {
			frontendDevServerURL = *frontendDevServerURLFlag
		}
		if loglevelFlag != nil {
			loglevel = *loglevelFlag
		}
	}

	assetConfig, err := assetserver.BuildAssetServerConfig(appoptions)
	if err != nil {
		return nil, err
	}

	if assetConfig.X静态资源 == nil && frontendDevServerURL != "" {
		myLogger.Warning("No AssetServer.Assets has been defined but a frontend DevServer, the frontend DevServer will not be used.")
		frontendDevServerURL = ""
		assetdir = ""
	}

	if frontendDevServerURL != "" {
		_, port, err := net.SplitHostPort(devServer)
		if err != nil {
			return nil, fmt.Errorf("unable to determine port of DevServer: %s", err)
		}

		ctx = context.WithValue(ctx, "assetserverport", port)

		ctx = context.WithValue(ctx, "frontenddevserverurl", frontendDevServerURL)

		externalURL, err := url.Parse(frontendDevServerURL)
		if err != nil {
			return nil, err
		}

		if externalURL.Host == "" {
			return nil, fmt.Errorf("Invalid frontend:dev:serverUrl missing protocol scheme?")
		}

		waitCb := func() { myLogger.Debug("Waiting for frontend DevServer '%s' to be ready", externalURL) }
		if !checkPortIsOpen(externalURL.Host, time.Minute, waitCb) {
			myLogger.Error("Timeout waiting for frontend DevServer")
		}

		handler := assetserver.NewExternalAssetsHandler(myLogger, assetConfig, externalURL)
		assetConfig.X静态资源 = nil
		assetConfig.X请求处理器 = handler
		assetConfig.X中间件 = nil

		myLogger.Info("Serving assets from frontend DevServer URL: %s", frontendDevServerURL)
	} else {
		if assetdir == "" {
			// 如果未定义assetdir，尝试从项目根目录和资源文件系统推断它。
			assetdir, err = tryInferAssetDirFromFS(assetConfig.X静态资源)
			if err != nil {
				return nil, fmt.Errorf("unable to infer the AssetDir from your Assets fs.FS: %w", err)
			}
		}

		if assetdir != "" {
			// 如果需要，让我们覆盖从磁盘上服务的资产
			absdir, err := filepath.Abs(assetdir)
			if err != nil {
				return nil, err
			}

			myLogger.Info("Serving assets from disk: %s", absdir)
			assetConfig.X静态资源 = os.DirFS(absdir)

			ctx = context.WithValue(ctx, "assetdir", assetdir)
		}
	}

	// 将已弃用的选项迁移到新的 AssetServer 选项
	appoptions.Assets弃用 = nil
	appoptions.AssetsHandler弃用 = nil
	appoptions.X绑定http请求 = &assetConfig

	if devServer != "" {
		ctx = context.WithValue(ctx, "devserver", devServer)
	}

	if loglevel != "" {
		level, err := pkglogger.X字符串到日志级别(loglevel)
		if err != nil {
			return nil, err
		}
		myLogger.SetLogLevel(level)
	}

	// 将日志器附加到上下文中
	ctx = context.WithValue(ctx, "logger", myLogger)
	ctx = context.WithValue(ctx, "buildtype", "dev")

	// Preflight checks
	err = PreflightChecks(appoptions, myLogger)
	if err != nil {
		return nil, err
	}

	// Merge default options
	options.MergeDefaults(appoptions)

	var menuManager *menumanager.Manager

	// 处理应用程序菜单
	if appoptions.X菜单 != nil {
		// 创建菜单管理器
		menuManager = menumanager.NewManager()
		err = menuManager.SetApplicationMenu(appoptions.X菜单)
		if err != nil {
			return nil, err
		}
	}

	// 创建绑定豁免 - 丑陋的解决方案。肯定有更优的方法
	bindingExemptions := []interface{}{
		appoptions.X绑定启动前函数,
		appoptions.X绑定应用退出函数,
		appoptions.X绑定DOM就绪函数,
		appoptions.X绑定应用关闭前函数,
	}
	appBindings := binding.NewBindings(myLogger, appoptions.X绑定调用方法, bindingExemptions, false, appoptions.X绑定常量枚举)

	eventHandler := runtime.NewEvents(myLogger)
	ctx = context.WithValue(ctx, "events", eventHandler)
	messageDispatcher := dispatcher.NewDispatcher(ctx, myLogger, appBindings, eventHandler, appoptions.X错误格式化)

	// 创建前端并注册到事件处理器
	desktopFrontend := desktop.NewFrontend(ctx, appoptions, myLogger, appBindings, messageDispatcher)
	appFrontend := devserver.NewFrontend(ctx, appoptions, myLogger, appBindings, messageDispatcher, menuManager, desktopFrontend)
	eventHandler.AddFrontend(appFrontend)
	eventHandler.AddFrontend(desktopFrontend)

	ctx = context.WithValue(ctx, "frontend", appFrontend)
	result := &App{
		ctx:              ctx,
		frontend:         appFrontend,
		logger:           myLogger,
		menuManager:      menuManager,
		startupCallback:  appoptions.X绑定启动前函数,
		shutdownCallback: appoptions.X绑定应用退出函数,
		debug:            true,
		devtoolsEnabled:  true,
	}

	result.options = appoptions

	return result, nil

}

func tryInferAssetDirFromFS(assets iofs.FS) (string, error) {
	if _, isEmbedFs := assets.(embed.FS); !isEmbedFs {
		// 我们只为嵌入式文件系统（embed.FS）的资源推断assetdir
		return "", nil
	}

	path, err := fs.FindPathToFile(assets, "index.html")
	if err != nil {
		return "", err
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(filepath.Join(path, "index.html")); err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf(
				"inferred assetdir '%s' does not exist or does not contain an 'index.html' file, "+
					"please specify it with -assetdir or set it in wails.json",
				path)
		}
		return "", err
	}

	return path, nil
}

func checkPortIsOpen(host string, timeout time.Duration, waitCB func()) (ret bool) {
	if timeout == 0 {
		timeout = time.Minute
	}

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, _ := net.DialTimeout("tcp", host, 2*time.Second)
		if conn != nil {
			conn.Close()
			return true
		}

		waitCB()
		time.Sleep(1 * time.Second)
	}
	return false
}
