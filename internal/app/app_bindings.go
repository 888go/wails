//go:build bindings

package app

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/leaanthony/gosod"
	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend/runtime/wrapper"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/project"
	"github.com/888go/wails/pkg/options"
)


// ff:
func (a *App) Run() error {

	// 创建绑定豁免 - 丑陋的解决方案。肯定有更优的方法
	bindingExemptions := []interface{}{
		a.options.X绑定启动前函数,
		a.options.X绑定应用退出函数,
		a.options.X绑定DOM就绪函数,
		a.options.X绑定应用关闭前函数,
	}

	// Check for CLI Flags
	bindingFlags := flag.NewFlagSet("bindings", flag.ContinueOnError)

	var tsPrefixFlag *string
	var tsPostfixFlag *string
	var tsOutputTypeFlag *string

	tsPrefix := os.Getenv("tsprefix")
	if tsPrefix == "" {
		tsPrefixFlag = bindingFlags.String("tsprefix", "", "Prefix for generated typescript entities")
	}

	tsSuffix := os.Getenv("tssuffix")
	if tsSuffix == "" {
		tsPostfixFlag = bindingFlags.String("tssuffix", "", "Suffix for generated typescript entities")
	}

	tsOutputType := os.Getenv("tsoutputtype")
	if tsOutputType == "" {
		tsOutputTypeFlag = bindingFlags.String("tsoutputtype", "", "Output type for generated typescript entities (classes|interfaces)")
	}

	_ = bindingFlags.Parse(os.Args[1:])
	if tsPrefixFlag != nil {
		tsPrefix = *tsPrefixFlag
	}
	if tsPostfixFlag != nil {
		tsSuffix = *tsPostfixFlag
	}
	if tsOutputTypeFlag != nil {
		tsOutputType = *tsOutputTypeFlag
	}

	appBindings := binding.NewBindings(a.logger, a.options.X绑定调用方法, bindingExemptions, IsObfuscated(), a.options.X绑定常量枚举)

	appBindings.SetTsPrefix(tsPrefix)
	appBindings.SetTsSuffix(tsSuffix)
	appBindings.SetOutputType(tsOutputType)

	err := generateBindings(appBindings)
	if err != nil {
		return err
	}
	return nil
}

// CreateApp 创建应用！

// ff:
// appoptions:
func CreateApp(appoptions *options.App) (*App, error) {
	// Set up logger
	myLogger := logger.New(appoptions.X日志记录器)
	myLogger.SetLogLevel(appoptions.X日志级别)

	result := &App{
		logger:  myLogger,
		options: appoptions,
	}

	return result, nil

}

func generateBindings(bindings *binding.Bindings) error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	projectConfig, err := project.Load(cwd)
	if err != nil {
		return err
	}

	wailsjsbasedir := filepath.Join(projectConfig.GetWailsJSDir(), "wailsjs")

	runtimeDir := filepath.Join(wailsjsbasedir, "runtime")
	_ = os.RemoveAll(runtimeDir)
	extractor := gosod.New(wrapper.RuntimeWrapper)
	err = extractor.Extract(runtimeDir, nil)
	if err != nil {
		return err
	}

	goBindingsDir := filepath.Join(wailsjsbasedir, "go")
	err = os.RemoveAll(goBindingsDir)
	if err != nil {
		return err
	}
	_ = fs.MkDirs(goBindingsDir)

	err = bindings.GenerateGoBindings(goBindingsDir)
	if err != nil {
		return err
	}

	return fs.SetPermissions(wailsjsbasedir, 0755)
}
