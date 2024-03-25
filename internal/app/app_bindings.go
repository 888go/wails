//go:build bindings

package app

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend/runtime/wrapper"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/project"
	"github.com/888go/wails/pkg/options"
	"github.com/leaanthony/gosod"
)

// ff:运行
func (a *App) Run() error {

	// 创建绑定豁免 - 丑陋的解决方案。肯定有更优的方法
	bindingExemptions := []interface{}{
		a.options.OnStartup,
		a.options.OnShutdown,
		a.options.OnDomReady,
		a.options.OnBeforeClose,
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

	appBindings := binding.NewBindings(a.logger, a.options.Bind, bindingExemptions, IsObfuscated(), a.options.EnumBind)

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
	myLogger := logger.New(appoptions.Logger)
	myLogger.SetLogLevel(appoptions.LogLevel)

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
