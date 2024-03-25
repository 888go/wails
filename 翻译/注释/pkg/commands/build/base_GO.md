
<原文开始>
// BaseBuilder is the common builder struct
<原文结束>

# <翻译开始>
// BaseBuilder 是通用构建器结构体
# <翻译结束>


<原文开始>
// NewBaseBuilder creates a new BaseBuilder
<原文结束>

# <翻译开始>
// NewBaseBuilder 创建一个新的 BaseBuilder
# <翻译结束>


<原文开始>
// SetProjectData sets the project data for this builder
<原文结束>

# <翻译开始>
// SetProjectData 为该构建器设置项目数据
# <翻译结束>


<原文开始>
// if file doesn't exist, ignore
<原文结束>

# <翻译开始>
// 如果文件不存在，则忽略
# <翻译结束>


<原文开始>
// Loop over all but 1 bytes
<原文结束>

# <翻译开始>
// 遍历除最后1个字节外的所有字节
# <翻译结束>


<原文开始>
// CleanUp does post-build housekeeping
<原文结束>

# <翻译开始>
// CleanUp 进行构建后清理工作
# <翻译结束>


<原文开始>
		// Delete file. We ignore errors because these files will be overwritten
		// by the next build anyway.
<原文结束>

# <翻译开始>
// 删除文件。我们忽略错误，因为这些文件无论如何将在下次构建时被覆盖。
# <翻译结束>


<原文开始>
// If we have a single argument, just return it
<原文结束>

# <翻译开始>
// 如果我们有一个单一的参数，直接返回它
# <翻译结束>


<原文开始>
// If an argument contains a space, quote it
<原文结束>

# <翻译开始>
// 如果参数中包含空格，则对该参数进行引号引用
# <翻译结束>


<原文开始>
// If we aren't using the standard compiler, add it to the filename
<原文结束>

# <翻译开始>
// 如果我们没有使用标准编译器，将其添加到文件名中
# <翻译结束>


<原文开始>
// Parse the `go version` output. EG: `go version go1.16 windows/amd64`
<原文结束>

# <翻译开始>
// 解析`go version`命令的输出结果。例如：`go version go1.16 windows/amd64`
# <翻译结束>


<原文开始>
// CompileProject compiles the project
<原文结束>

# <翻译开始>
// CompileProject 编译项目
# <翻译结束>


<原文开始>
// Check if the runtime wrapper exists
<原文结束>

# <翻译开始>
// 检查运行时包装器是否存在
# <翻译结束>


<原文开始>
// Default go build command
<原文结束>

# <翻译开始>
// 默认的Go构建命令
# <翻译结束>


<原文开始>
// Add better debugging flags
<原文结束>

# <翻译开始>
// 添加更好的调试标志
# <翻译结束>


<原文开始>
// Add webview2 strategy if we have it
<原文结束>

# <翻译开始>
// 如果我们有webview2策略，则添加它
# <翻译结束>


<原文开始>
// This mode allows you to debug a production build (not dev build)
<原文结束>

# <翻译开始>
// 此模式允许您调试生产构建（非开发构建）
# <翻译结束>


<原文开始>
// This options allows you to enable devtools in production build (not dev build as it's always enabled there)
<原文结束>

# <翻译开始>
// 这个选项允许你在生产构建中启用开发者工具（在开发构建中它始终是启用的，所以无需在此设置）
# <翻译结束>


<原文开始>
// Add the output type build tag
<原文结束>

# <翻译开始>
// 添加输出类型构建标签
# <翻译结束>


<原文开始>
// Get application build directory
<原文结束>

# <翻译开始>
// 获取应用程序构建目录
# <翻译结束>


<原文开始>
	// Add CGO flags
	// TODO: Remove this as we don't generate headers any more
	// We use the project/build dir as a temporary place for our generated c headers
<原文结束>

# <翻译开始>
// 添加CGO标志
// TODO: 移除这一行，因为我们现在已经不再生成头文件了
// 我们使用项目/构建目录作为临时位置来存放我们生成的C语言头文件
# <翻译结束>


<原文开始>
// Use shell.UpsertEnv so we don't overwrite user's CGO_CFLAGS
<原文结束>

# <翻译开始>
// 使用shell.UpsertEnv，以免覆盖用户自定义的CGO_CFLAGS环境变量
# <翻译结束>


<原文开始>
// Use shell.UpsertEnv so we don't overwrite user's CGO_CXXFLAGS
<原文结束>

# <翻译开始>
// 使用shell.UpsertEnv，这样我们不会覆盖用户的CGO_CXXFLAGS环境变量
# <翻译结束>


<原文开始>
			// Determine version so we can link to newer frameworks
			// Why doesn't CGO have this option?!?!
<原文结束>

# <翻译开始>
// 确定版本以便链接到更新的框架
// 为什么CGO没有这个选项呢？！？！
# <翻译结束>


<原文开始>
// Set the minimum Mac SDK to 10.13
<原文结束>

# <翻译开始>
// 设置Mac SDK的最低版本为10.13
# <翻译结束>


<原文开始>
// Format error if we have one
<原文结束>

# <翻译开始>
// 如果我们有错误，则格式化该错误
# <翻译结束>


<原文开始>
// Do we have upx installed?
<原文结束>

# <翻译开始>
// 我们是否安装了 upx？
# <翻译结束>


<原文开始>
// NpmInstall runs "npm install" in the given directory
<原文结束>

# <翻译开始>
// NpmInstall 在给定的目录中运行 "npm install"
# <翻译结束>


<原文开始>
// NpmInstallUsingCommand runs the given install command in the specified npm project directory
<原文结束>

# <翻译开始>
// NpmInstallUsingCommand 在指定的npm项目目录中运行给定的安装命令
# <翻译结束>


<原文开始>
// Check package.json exists
<原文结束>

# <翻译开始>
// 检查 package.json 是否存在
# <翻译结束>


<原文开始>
// No package.json, no install
<原文结束>

# <翻译开始>
// 没有package.json，无需安装
# <翻译结束>


<原文开始>
// Get the MD5 sum of package.json
<原文结束>

# <翻译开始>
// 获取package.json的MD5校验和
# <翻译结束>


<原文开始>
// Check whether we need to npm install
<原文结束>

# <翻译开始>
// 检查是否需要执行npm install
# <翻译结束>


<原文开始>
// Install if node_modules doesn't exist
<原文结束>

# <翻译开始>
// 如果node_modules不存在，则进行安装
# <翻译结束>


<原文开始>
// check if forced install
<原文结束>

# <翻译开始>
// 检查是否为强制安装
# <翻译结束>


<原文开始>
// Split up the InstallCommand and execute it
<原文结束>

# <翻译开始>
// 将InstallCommand拆分并执行
# <翻译结束>


<原文开始>
// NpmRun executes the npm target in the provided directory
<原文结束>

# <翻译开始>
// NpmRun在指定的目录中执行npm目标
# <翻译结束>


<原文开始>
// NpmRunWithEnvironment executes the npm target in the provided directory, with the given environment variables
<原文结束>

# <翻译开始>
// NpmRunWithEnvironment 在指定的目录下，使用给定的环境变量执行npm目标
# <翻译结束>


<原文开始>
// BuildFrontend executes the `npm build` command for the frontend directory
<原文结束>

# <翻译开始>
// BuildFrontend 执行针对前端目录的 `npm build` 命令
# <翻译结束>


<原文开始>
// Check there is an 'InstallCommand' provided in wails.json
<原文结束>

# <翻译开始>
// 检查 wails.json 中是否提供了 'InstallCommand'
# <翻译结束>


<原文开始>
// Check if there is a build command
<原文结束>

# <翻译开始>
// 检查是否存在构建命令
# <翻译结束>

