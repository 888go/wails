package build

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/shlex"
	"github.com/pterm/pterm"
	"github.com/samber/lo"

	"github.com/888go/wails/internal/staticanalysis"
	"github.com/888go/wails/pkg/commands/bindings"

	"github.com/888go/wails/internal/fs"

	"github.com/888go/wails/internal/shell"

	"github.com/888go/wails/internal/project"
	"github.com/888go/wails/pkg/clilogger"
)

// Mode 是用于表示构建模式的类型
type Mode int

const (
	// Dev mode
	Dev Mode = iota
	// Production mode
	Production
	// Debug build
	Debug
)

// Options 包含所有构建选项以及项目数据
type Options struct {
	LDFlags           string               // 可选的标志，用于传递给链接器
	UserTags          []string             // 传递给 Go 编译器的标签
	Logger            *clilogger.CLILogger // 所有输出都发送到日志器
	OutputType        string               // EG: 示例，桌面，服务器...
	Mode              Mode                 // release or dev
	Devtools          bool                 // 在生产环境中启用开发工具
	ProjectData       *project.Project     // The project data
	Pack              bool                 // 在构建后为应用程序创建一个包
	Platform          string               // 需要构建的目标平台
	Arch              string               // 此处用于构建的架构
	Compiler          string               // 需要使用的编译器命令
	SkipModTidy       bool                 // 在编译前跳过 mod tidy
	IgnoreFrontend    bool                 // 表示前端无需构建
	IgnoreApplication bool                 // 表示应用程序无需构建
	OutputFile        string               // 重写输出文件名
	BinDirectory      string               // 使用该目录来写入构建的应用程序
	CleanBinDirectory bool                 // 表示在构建之前是否应清理 bin 输出目录
	CompiledBinary    string               // 完全限定的编译后二进制文件路径
	KeepAssets        bool                 // 保留生成的资产/文件
	Verbosity         int                  // 详细程度等级 (0 - 静默模式, 1 - 默认模式, 2 - 详细模式)
	Compress          bool                 // 压缩最终的二进制文件
	CompressFlags     string               // Flags to pass to UPX
	WebView2Strategy  string               // WebView2 安装程序策略
	RunDelve          bool                 // 表示在构建后是否应运行 delve
	WailsJSDir        string               // 用于生成wailsjs模块的目录
	ForceBuild        bool                 // Force
	BundleName        string               // Bundlename for Mac
	TrimPath          bool                 // Use Go's trimpath compiler flag
	RaceDetector      bool                 // Build with Go's race detector
	WindowsConsole    bool                 // 表示应保留Windows控制台
	Obfuscated        bool                 // 表示绑定的方法应被混淆
	GarbleArgs        string               // Garble函数的参数
	SkipBindings      bool                 // 跳过绑定生成
}

// Build the project!
func X构建项目(选项 *Options) (string, error) {
	// Extract logger
	outputLogger := 选项.Logger

	// Get working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// wails js dir
	选项.WailsJSDir = 选项.ProjectData.GetWailsJSDir()

	// Set build directory
	选项.BinDirectory = filepath.Join(选项.ProjectData.GetBuildDir(), "bin")

	// Save the project type
	选项.ProjectData.OutputType = 选项.OutputType

	// Create builder
	var builder Builder

	switch 选项.OutputType {
	case "desktop":
		builder = newDesktopBuilder(选项)
	case "dev":
		builder = newDesktopBuilder(选项)
	default:
		return "", fmt.Errorf("cannot build assets for output type %s", 选项.ProjectData.OutputType)
	}

	// 设置我们的清理方法
	defer builder.CleanUp()

	// Initialise Builder
	builder.SetProjectData(选项.ProjectData)

	hookArgs := map[string]string{
		"${platform}": 选项.Platform + "/" + 选项.Arch,
	}

	for _, hook := range []string{选项.Platform + "/" + 选项.Arch, 选项.Platform + "/*", "*/*"} {
		if err := execPreBuildHook(outputLogger, 选项, hook, hookArgs); err != nil {
			return "", err
		}
	}

	// 如果嵌入式目录不存在，则创建它们
	if err := CreateEmbedDirectories(cwd, 选项); err != nil {
		return "", err
	}

	// Generate bindings
	if !选项.SkipBindings {
		err = GenerateBindings(选项)
		if err != nil {
			return "", err
		}
	}

	if !选项.IgnoreFrontend {
		err = builder.BuildFrontend(outputLogger)
		if err != nil {
			return "", err
		}
	}

	compileBinary := ""
	if !选项.IgnoreApplication {
		compileBinary, err = execBuildApplication(builder, 选项)
		if err != nil {
			return "", err
		}

		hookArgs["${bin}"] = compileBinary
		for _, hook := range []string{选项.Platform + "/" + 选项.Arch, 选项.Platform + "/*", "*/*"} {
			if err := execPostBuildHook(outputLogger, 选项, hook, hookArgs); err != nil {
				return "", err
			}
		}

	}
	return compileBinary, nil
}

func CreateEmbedDirectories(cwd string, buildOptions *Options) error {
	path := cwd
	if buildOptions.ProjectData != nil {
		path = buildOptions.ProjectData.Path
	}
	embedDetails, err := staticanalysis.GetEmbedDetails(path)
	if err != nil {
		return err
	}

	for _, embedDetail := range embedDetails {
		fullPath := embedDetail.GetFullPath()
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			err := os.MkdirAll(fullPath, 0o755)
			if err != nil {
				return err
			}
			f, err := os.Create(filepath.Join(fullPath, "gitkeep"))
			if err != nil {
				return err
			}
			_ = f.Close()
		}
	}

	return nil
}

func fatal(message string) {
	printer := pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.FatalMessageStyle,
		Prefix: pterm.Prefix{
			Style: &pterm.ThemeDefault.FatalPrefixStyle,
			Text:  " FATAL ",
		},
	}
	printer.Println(message)
	os.Exit(1)
}

func printBulletPoint(text string, args ...any) {
	item := pterm.BulletListItem{
		Level: 2,
		Text:  text,
	}
	t, err := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{item}).Srender()
	if err != nil {
		fatal(err.Error())
	}
	t = strings.Trim(t, "\n\r")
	pterm.Printf(t, args...)
}

func GenerateBindings(buildOptions *Options) error {
	obfuscated := buildOptions.Obfuscated
	if obfuscated {
		printBulletPoint("Generating obfuscated bindings: ")
		buildOptions.UserTags = append(buildOptions.UserTags, "obfuscated")
	} else {
		printBulletPoint("Generating bindings: ")
	}

	if buildOptions.ProjectData.Bindings.TsGeneration.OutputType == "" {
		buildOptions.ProjectData.Bindings.TsGeneration.OutputType = "classes"
	}

	// Generate Bindings
	output, err := bindings.GenerateBindings(bindings.Options{
		Compiler:     buildOptions.Compiler,
		Tags:         buildOptions.UserTags,
		GoModTidy:    !buildOptions.SkipModTidy,
		TsPrefix:     buildOptions.ProjectData.Bindings.TsGeneration.Prefix,
		TsSuffix:     buildOptions.ProjectData.Bindings.TsGeneration.Suffix,
		TsOutputType: buildOptions.ProjectData.Bindings.TsGeneration.OutputType,
	})
	if err != nil {
		return err
	}

	if buildOptions.Verbosity == VERBOSE {
		pterm.Info.Println(output)
	}

	pterm.Println("Done.")

	return nil
}

func execBuildApplication(builder Builder, options *Options) (string, error) {
// 如果我们正在为Windows系统构建，那么在编译之前我们需要生成资源包。
// 这将在项目根目录下生成一个.syso文件
	if options.Pack && options.Platform == "windows" {
		printBulletPoint("Generating application assets: ")
		err := packageApplicationForWindows(options)
		if err != nil {
			return "", err
		}
		pterm.Println("Done.")

		// 当我们完成时，我们将需要移除syso文件
		defer func() {
			err := os.Remove(filepath.Join(options.ProjectData.Path, options.ProjectData.Name+"-res.syso"))
			if err != nil {
				fatal(err.Error())
			}
		}()
	}

	// 编译应用程序
	printBulletPoint("Compiling application: ")

	if options.Platform == "darwin" && options.Arch == "universal" {
		outputFile := builder.OutputFilename(options)
		amd64Filename := outputFile + "-amd64"
		arm64Filename := outputFile + "-arm64"

		// Build amd64 first
		options.Arch = "amd64"
		options.OutputFile = amd64Filename
		options.CleanBinDirectory = false
		if options.Verbosity == VERBOSE {
			pterm.Println("Building AMD64 Target: " + filepath.Join(options.BinDirectory, options.OutputFile))
		}
		err := builder.CompileProject(options)
		if err != nil {
			return "", err
		}
		// Build arm64
		options.Arch = "arm64"
		options.OutputFile = arm64Filename
		options.CleanBinDirectory = false
		if options.Verbosity == VERBOSE {
			pterm.Println("Building ARM64 Target: " + filepath.Join(options.BinDirectory, options.OutputFile))
		}
		err = builder.CompileProject(options)

		if err != nil {
			return "", err
		}
		// Run lipo
		if options.Verbosity == VERBOSE {
			pterm.Println(fmt.Sprintf("Running lipo: lipo -create -output %s %s %s", outputFile, amd64Filename, arm64Filename))
		}
		_, stderr, err := shell.RunCommand(options.BinDirectory, "lipo", "-create", "-output", outputFile, amd64Filename, arm64Filename)
		if err != nil {
			return "", fmt.Errorf("%s - %s", err.Error(), stderr)
		}
		// Remove temp binaries
		err = fs.DeleteFile(filepath.Join(options.BinDirectory, amd64Filename))
		if err != nil {
			return "", err
		}
		err = fs.DeleteFile(filepath.Join(options.BinDirectory, arm64Filename))
		if err != nil {
			return "", err
		}
		options.ProjectData.OutputFilename = outputFile
		options.CompiledBinary = filepath.Join(options.BinDirectory, outputFile)
	} else {
		err := builder.CompileProject(options)
		if err != nil {
			return "", err
		}
	}

	if runtime.GOOS == "darwin" {
		// 删除隔离属性
		if _, err := os.Stat(options.CompiledBinary); os.IsNotExist(err) {
			return "", fmt.Errorf("compiled binary does not exist at path: %s", options.CompiledBinary)
		}
		stdout, stderr, err := shell.RunCommand(options.BinDirectory, "xattr", "-rc", options.CompiledBinary)
		if err != nil {
			return "", fmt.Errorf("%s - %s", err.Error(), stderr)
		}
		if options.Verbosity == VERBOSE && stdout != "" {
			pterm.Info.Println(stdout)
		}
	}

	pterm.Println("Done.")

	// 非Windows系统下，我们是否需要打包应用？
	if options.Pack && options.Platform != "windows" {

		printBulletPoint("Packaging application: ")

		// TODO: 允许跨平台构建
		err := packageProject(options, runtime.GOOS)
		if err != nil {
			return "", err
		}
		pterm.Println("Done.")
	}

	if options.Platform == "windows" {
		const nativeWebView2Loader = "native_webview2loader"

		tags := options.UserTags
		if lo.Contains(tags, nativeWebView2Loader) {
			message := "You are using the legacy native WebView2Loader. This loader will be deprecated in the near future. Please report any bugs related to the new loader: https://github.com/wailsapp/wails/issues/2004"
			pterm.Warning.Println(message)
		} else {
			tags = append(tags, nativeWebView2Loader)
			message := fmt.Sprintf("Wails is now using the new Go WebView2Loader. If you encounter any issues with it, please report them to https://github.com/wailsapp/wails/issues/2004. You could also use the old legacy loader with `-tags %s`, but keep in mind this will be deprecated in the near future.", strings.Join(tags, ","))
			pterm.Info.Println(message)
		}
	}

	if options.Platform == "darwin" && (options.Mode == Debug || options.Devtools) {
		pterm.Warning.Println("This darwin build contains the use of private APIs. This will not pass Apple's AppStore approval process. Please use it only as a test build for testing and debug purposes.")
	}

	return options.CompiledBinary, nil
}

func execPreBuildHook(outputLogger *clilogger.CLILogger, options *Options, hookIdentifier string, argReplacements map[string]string) error {
	preBuildHook := options.ProjectData.PreBuildHooks[hookIdentifier]
	if preBuildHook == "" {
		return nil
	}

	return executeBuildHook(outputLogger, options, hookIdentifier, argReplacements, preBuildHook, "pre")
}

func execPostBuildHook(outputLogger *clilogger.CLILogger, options *Options, hookIdentifier string, argReplacements map[string]string) error {
	postBuildHook := options.ProjectData.PostBuildHooks[hookIdentifier]
	if postBuildHook == "" {
		return nil
	}

	return executeBuildHook(outputLogger, options, hookIdentifier, argReplacements, postBuildHook, "post")
}

func executeBuildHook(_ *clilogger.CLILogger, options *Options, hookIdentifier string, argReplacements map[string]string, buildHook string, hookName string) error {
	if !options.ProjectData.RunNonNativeBuildHooks {
		if hookIdentifier == "" {
			// That's the global hook
		} else {
			platformOfHook := strings.Split(hookIdentifier, "/")[0]
			if platformOfHook == "*" {
				// 这没问题，我们还没有为钩子指定特定的平台
			} else if platformOfHook == runtime.GOOS {
				// 此钩子用于宿主平台
			} else {
				// 跳过非原生的钩子
				printBulletPoint(fmt.Sprintf("Non native build hook '%s': Skipping.", hookIdentifier))
				return nil
			}
		}
	}

	printBulletPoint("Executing %s build hook '%s': ", hookName, hookIdentifier)
	args, err := shlex.Split(buildHook)
	if err != nil {
		return fmt.Errorf("could not parse %s build hook command: %w", hookName, err)
	}
	for i, arg := range args {
		newArg := argReplacements[arg]
		if newArg == "" {
			continue
		}
		args[i] = newArg
	}

	if options.Verbosity == VERBOSE {
		pterm.Info.Println(strings.Join(args, " "))
	}

	if !fs.DirExists(options.BinDirectory) {
		if err := fs.MkDirs(options.BinDirectory); err != nil {
			return fmt.Errorf("could not create target directory: %s", err.Error())
		}
	}

	stdout, stderr, err := shell.RunCommand(options.BinDirectory, args[0], args[1:]...)
	if options.Verbosity == VERBOSE {
		pterm.Info.Println(stdout)
	}
	if err != nil {
		return fmt.Errorf("%s - %s", err.Error(), stderr)
	}
	pterm.Println("Done.")

	return nil
}
