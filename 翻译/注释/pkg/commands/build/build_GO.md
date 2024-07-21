
<原文开始>
// Mode is the type used to indicate the build modes
<原文结束>

# <翻译开始>
// Mode 是用于表示构建模式的类型
# <翻译结束>


<原文开始>
// Options contains all the build options as well as the project data
<原文结束>

# <翻译开始>
// Options 包含所有构建选项以及项目数据
# <翻译结束>


<原文开始>
// Optional flags to pass to linker
<原文结束>

# <翻译开始>
// 可选的标志，用于传递给链接器
# <翻译结束>


<原文开始>
// Tags to pass to the Go compiler
<原文结束>

# <翻译开始>
// 传递给 Go 编译器的标签
# <翻译结束>


<原文开始>
// All output to the logger
<原文结束>

# <翻译开始>
// 所有输出都发送到日志器
# <翻译结束>


<原文开始>
// EG: desktop, server....
<原文结束>

# <翻译开始>
// EG: 示例，桌面，服务器...
# <翻译结束>


<原文开始>
// Enable devtools in production
<原文结束>

# <翻译开始>
// 在生产环境中启用开发工具
# <翻译结束>


<原文开始>
// Create a package for the app after building
<原文结束>

# <翻译开始>
// 在构建后为应用程序创建一个包
# <翻译结束>


<原文开始>
// The platform to build for
<原文结束>

# <翻译开始>
// 需要构建的目标平台
# <翻译结束>


<原文开始>
// The architecture to build for
<原文结束>

# <翻译开始>
// 此处用于构建的架构
# <翻译结束>


<原文开始>
// The compiler command to use
<原文结束>

# <翻译开始>
// 需要使用的编译器命令
# <翻译结束>


<原文开始>
//  Skip mod tidy before compile
<原文结束>

# <翻译开始>
// 在编译前跳过 mod tidy
# <翻译结束>


<原文开始>
// Indicates if the frontend does not need building
<原文结束>

# <翻译开始>
// 表示前端无需构建
# <翻译结束>


<原文开始>
// Indicates if the application does not need building
<原文结束>

# <翻译开始>
// 表示应用程序无需构建
# <翻译结束>


<原文开始>
// Override the output filename
<原文结束>

# <翻译开始>
// 重写输出文件名
# <翻译结束>


<原文开始>
// Directory to use to write the built applications
<原文结束>

# <翻译开始>
// 使用该目录来写入构建的应用程序
# <翻译结束>


<原文开始>
// Indicates if the bin output directory should be cleaned before building
<原文结束>

# <翻译开始>
// 表示在构建之前是否应清理 bin 输出目录
# <翻译结束>


<原文开始>
// Fully qualified path to the compiled binary
<原文结束>

# <翻译开始>
// 完全限定的编译后二进制文件路径
# <翻译结束>


<原文开始>
// Keep the generated assets/files
<原文结束>

# <翻译开始>
// 保留生成的资产/文件
# <翻译结束>


<原文开始>
// Verbosity level (0 - silent, 1 - default, 2 - verbose)
<原文结束>

# <翻译开始>
// 详细程度等级 (0 - 静默模式, 1 - 默认模式, 2 - 详细模式)
# <翻译结束>


<原文开始>
// Compress the final binary
<原文结束>

# <翻译开始>
// 压缩最终的二进制文件
# <翻译结束>


<原文开始>
// WebView2 installer strategy
<原文结束>

# <翻译开始>
// WebView2 安装程序策略
# <翻译结束>


<原文开始>
// Indicates if we should run delve after the build
<原文结束>

# <翻译开始>
// 表示在构建后是否应运行 delve
# <翻译结束>


<原文开始>
// Directory to generate the wailsjs module
<原文结束>

# <翻译开始>
// 用于生成wailsjs模块的目录
# <翻译结束>


<原文开始>
// Indicates that the windows console should be kept
<原文结束>

# <翻译开始>
// 表示应保留Windows控制台
# <翻译结束>


<原文开始>
// Indicates that bound methods should be obfuscated
<原文结束>

# <翻译开始>
// 表示绑定的方法应被混淆
# <翻译结束>


<原文开始>
// The arguments for Garble
<原文结束>

# <翻译开始>
// Garble函数的参数
# <翻译结束>


<原文开始>
// Skip binding generation
<原文结束>

# <翻译开始>
// 跳过绑定生成
# <翻译结束>


<原文开始>
// Set up our clean up method
<原文结束>

# <翻译开始>
// 设置我们的清理方法
# <翻译结束>


<原文开始>
// Create embed directories if they don't exist
<原文结束>

# <翻译开始>
// 如果嵌入式目录不存在，则创建它们
# <翻译结束>


<原文开始>
	// If we are building for windows, we will need to generate the asset bundle before
	// compilation. This will be a .syso file in the project root
<原文结束>

# <翻译开始>
	// 如果我们正在为Windows系统构建，那么在编译之前我们需要生成资源包。
	// 这将在项目根目录下生成一个.syso文件
# <翻译结束>


<原文开始>
// When we finish, we will want to remove the syso file
<原文结束>

# <翻译开始>
// 当我们完成时，我们将需要移除syso文件
# <翻译结束>


<原文开始>
// Compile the application
<原文结束>

# <翻译开始>
// 编译应用程序
# <翻译结束>


<原文开始>
// Do we need to pack the app for non-windows?
<原文结束>

# <翻译开始>
// 非Windows系统下，我们是否需要打包应用？
# <翻译结束>


<原文开始>
// TODO: Allow cross platform build
<原文结束>

# <翻译开始>
// TODO: 允许跨平台构建
# <翻译结束>


<原文开始>
// That's OK, we don't have a specific platform of the hook
<原文结束>

# <翻译开始>
// 这没问题，我们还没有为钩子指定特定的平台
# <翻译结束>


<原文开始>
// The hook is for host platform
<原文结束>

# <翻译开始>
// 此钩子用于宿主平台
# <翻译结束>


<原文开始>
// Skip a hook which is not native
<原文结束>

# <翻译开始>
// 跳过非原生的钩子
# <翻译结束>


<原文开始>
// Remove quarantine attribute
<原文结束>

# <翻译开始>
// 删除隔离属性
# <翻译结束>

