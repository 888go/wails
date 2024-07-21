
<原文开始>
// Project holds the data related to a Wails project
<原文结束>

# <翻译开始>
// Project 结构体持有与Wails项目相关联的数据
# <翻译结束>


<原文开始>
// Commands used in `wails dev`
<原文结束>

# <翻译开始>
// 以下命令在 `wails dev` 中使用
# <翻译结束>


<原文开始>
// The url of the external wails dev server. If this is set, this server is used for the frontend. Default ""
<原文结束>

# <翻译开始>
// 外部Wails开发服务器的URL。如果设置了这个值，那么将使用这个服务器作为前端服务。默认为 ""
# <翻译结束>


<原文开始>
// Directory to generate the API Module
<原文结束>

# <翻译开始>
// 用于生成API模块的目录
# <翻译结束>


<原文开始>
// The path to the project directory
<原文结束>

# <翻译开始>
// 项目目录的路径
# <翻译结束>


<原文开始>
// The type of application. EG: Desktop, Server, etc
<原文结束>

# <翻译开始>
// 应用程序的类型。例如：桌面应用、服务器应用等
# <翻译结束>


<原文开始>
// RunNonNativeBuildHooks will run build hooks though they are defined for a GOOS which is not equal to the host os
<原文结束>

# <翻译开始>
// RunNonNativeBuildHooks 将运行构建钩子，即使这些钩子是为与主机操作系统不同的 GOOS 环境定义的。
# <翻译结束>


<原文开始>
	// Build hooks for different targets, the hooks are executed in the following order
	// Key: GOOS/GOARCH - Executed at build level before/after a build of the specific platform and arch
	// Key: GOOS/*      - Executed at build level before/after a build of the specific platform
	// Key: */*         - Executed at build level before/after a build
	// The following keys are not yet supported.
	// Key: GOOS        - Executed at platform level before/after all builds of the specific platform
	// Key: *           - Executed at platform level before/after all builds of a platform
	// Key: [empty]     - Executed at global level before/after all builds of all platforms
<原文结束>

# <翻译开始>
	// 为不同目标构建钩子，这些钩子按照以下顺序执行：
	// Key: GOOS/GOARCH - 在特定平台和架构的构建级别上，在构建前后执行
	// Key: GOOS/*      - 在特定平台的构建级别上，在构建前后执行
	// Key: */*         - 在构建级别上，在所有构建前后执行
	// 下列键目前还不支持。
	// Key: GOOS        - 在特定平台级别上，在该平台的所有构建之前/之后执行
	// Key: *           - 在平台级别上，在所有平台的所有构建之前/之后执行
	// Key: [空]        - 在全局级别上，在所有平台的所有构建之前/之后执行
# <翻译结束>


<原文开始>
// The application information
<原文结束>

# <翻译开始>
// 应用程序信息
# <翻译结束>


<原文开始>
// Fully qualified filename
<原文结束>

# <翻译开始>
// 完全限定文件名
# <翻译结束>


<原文开始>
// The debounce time for hot-reload of the built-in dev server. Default 100
<原文结束>

# <翻译开始>
// 内置开发服务器热重载的防抖时间，默认为100
// （注：debounce 时间通常是指在连续触发事件后，会等待一段固定的时间再去执行回调函数，用于限制函数在一定时间段内只能被执行一次，从而避免短时间内大量无用计算或网络请求。这里的“热重载”一般指的是当代码发生变化时，自动重新加载并应用更改到运行中的程序。）
# <翻译结束>


<原文开始>
// The address to bind the wails dev server to. Default "localhost:34115"
<原文结束>

# <翻译开始>
// 绑定wails开发服务器的地址。默认为 "localhost:34115"
# <翻译结束>


<原文开始>
// Arguments that are forwared to the application in dev mode
<原文结束>

# <翻译开始>
// 在开发模式下传递给应用程序的参数
# <翻译结束>


<原文开始>
// Create default name if not given
<原文结束>

# <翻译开始>
// 如果未提供名称，则创建默认名称
# <翻译结束>


<原文开始>
// Author stores details about the application author
<原文结束>

# <翻译开始>
// Author 用于存储应用程序作者的详细信息
# <翻译结束>


<原文开始>
// Parse the given JSON data into a Project struct
<原文结束>

# <翻译开始>
// 将给定的JSON数据解析为一个Project结构体
# <翻译结束>


<原文开始>
// Load the project from the current working directory
<原文结束>

# <翻译开始>
// 从当前工作目录加载项目
# <翻译结束>

