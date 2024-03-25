
<原文开始>
// Application runs the application in dev mode
<原文结束>

# <翻译开始>
// Application 在开发模式下运行应用程序
# <翻译结束>


<原文开始>
// Update go.mod to use current wails version
<原文结束>

# <翻译开始>
// 更新go.mod文件以使用当前wails版本
# <翻译结束>


<原文开始>
// Run go mod tidy to ensure we're up-to-date
<原文结束>

# <翻译开始>
// 运行 go mod tidy 以确保我们的依赖是最新的
# <翻译结束>


<原文开始>
// Build the frontend if requested, but ignore building the application itself.
<原文结束>

# <翻译开始>
// 如果有请求，则构建前端，但忽略构建应用程序本身。
# <翻译结束>


<原文开始>
// frontend:dev:watcher command.
<原文结束>

# <翻译开始>
// 前端：开发：监视器命令。
# <翻译结束>


<原文开始>
// Do initial build but only for the application.
<原文结束>

# <翻译开始>
// 只针对应用程序进行首次构建
# <翻译结束>


<原文开始>
// Show dev server URL in terminal after 3 seconds
<原文结束>

# <翻译开始>
// 在终端中显示开发服务器URL，3秒后
# <翻译结束>


<原文开始>
// Watch for changes and trigger restartApp()
<原文结束>

# <翻译开始>
// 监听变更并触发 restartApp() 函数
# <翻译结束>


<原文开始>
// Kill the current program if running and remove dev binary
<原文结束>

# <翻译开始>
// 如果当前程序正在运行，则终止程序并移除开发版二进制文件
# <翻译结束>


<原文开始>
// Reset the process and the binary so defer knows about it and is a nop.
<原文结束>

# <翻译开始>
// 重置进程和二进制文件，以便defer语句能够识别并将其视为空操作（nop）。
# <翻译结束>


<原文开始>
// runFrontendDevWatcherCommand will run the `frontend:dev:watcher` command if it was given, ex- `npm run dev`
<原文结束>

# <翻译开始>
// runFrontendDevWatcherCommand 将会在接收到相应命令时执行 `frontend:dev:watcher` 命令，例如：`npm run dev`
# <翻译结束>


<原文开始>
// That's fine, then most probably it was not vite that was running
<原文结束>

# <翻译开始>
// 那么，这很可能不是vite在运行
# <翻译结束>


<原文开始>
// restartApp does the actual rebuilding of the application when files change
<原文结束>

# <翻译开始>
// restartApp 当文件发生更改时，执行应用程序的实际重建工作
# <翻译结束>


<原文开始>
// Kill existing binary if need be
<原文结束>

# <翻译开始>
// 如果需要的话，杀死已存在的二进制文件
# <翻译结束>


<原文开始>
// Set environment variables accordingly
<原文结束>

# <翻译开始>
// 根据实际情况设置环境变量
# <翻译结束>


<原文开始>
// Start up new binary with correct args
<原文结束>

# <翻译开始>
// 使用正确的参数启动新的二进制文件
# <翻译结束>


<原文开始>
// doWatcherLoop is the main watch loop that runs while dev is active
<原文结束>

# <翻译开始>
// doWatcherLoop 是主监视循环，在dev处于活动状态时运行
# <翻译结束>


<原文开始>
// create the project files watcher
<原文结束>

# <翻译开始>
// 创建项目文件观察器
# <翻译结束>


<原文开始>
// If we are using an external dev server, the reloading of the frontend part can be skipped or if the user requested it
<原文结束>

# <翻译开始>
// 如果我们正在使用外部开发服务器，前端部分的重新加载可以被跳过，或者如果用户请求这样做
# <翻译结束>


<原文开始>
// Iterate all file patterns
<原文结束>

# <翻译开始>
// 遍历所有文件模式
# <翻译结束>


<原文开始>
// Handle write operations
<原文结束>

# <翻译开始>
// 处理写操作
# <翻译结束>


<原文开始>
// Handle new fs entries that are created
<原文结束>

# <翻译开始>
// 处理新创建的文件系统条目
# <翻译结束>


<原文开始>
// If this is a folder, add it to our watch list
<原文结束>

# <翻译开始>
// 如果这是一个文件夹，将其添加到我们的监视列表中
# <翻译结束>


<原文开始>
// node_modules is BANNED!
<原文结束>

# <翻译开始>
// node_modules 是被禁止的！
# <翻译结束>


<原文开始>
					// Handle creation of new file.
					// Note: On some platforms an update to a file is represented as
					// REMOVE -> CREATE instead of WRITE, so this is not only new files
					// but also updates to existing files
<原文结束>

# <翻译开始>
// 处理新文件的创建。
// 注意：在某些平台上，对文件的更新表现为 REMOVE（删除）-> CREATE（创建），而非 WRITE（写入），因此这不仅包括新建文件，
// 还包括对现有文件的更新操作。
# <翻译结束>


<原文开始>
// If we have a new process, saveConfig it
<原文结束>

# <翻译开始>
// 如果我们有一个新的进程，保存其配置
# <翻译结束>

