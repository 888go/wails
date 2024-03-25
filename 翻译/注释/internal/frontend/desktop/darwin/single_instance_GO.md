
<原文开始>
//export HandleSecondInstanceData
<原文结束>

# <翻译开始>
//export HandleSecondInstanceData
// 导出HandleSecondInstanceData函数，以便在C语言或其他外部语言中调用
// （由于上下文不完整，无法提供更详尽的翻译，但大体上，Go语言中的`//export`注释是用来标记一个函数，表示该函数需要被cgo暴露给C代码或者其他使用CGO的环境，使得它们可以调用Go编写的这个函数。）
# <翻译结束>


<原文开始>
// CreateLockFile tries to create a file with given name and acquire an
// exclusive lock on it. If the file already exists AND is still locked, it will
// fail.
<原文结束>

# <翻译开始>
// CreateLockFile尝试使用给定的名称创建一个文件并获取其独占锁。
// 如果文件已存在且仍然被锁定，则操作将会失败。
# <翻译结束>


<原文开始>
// Flock failed for some other reason than other instance already lock it. Print it in logs for possible debugging.
<原文结束>

# <翻译开始>
// 如果 flock（文件锁定）由于除其他实例已锁定之外的其他原因失败，则将其打印到日志中，以便于可能的调试。
# <翻译结束>


<原文开始>
// If app is sandboxed, golang os.TempDir() will return path that will not be accessible. So use native macOS temp dir function.
<原文结束>

# <翻译开始>
// 如果应用处于沙箱环境，golang 的 os.TempDir() 函数会返回一个无法访问的路径。因此，这里使用 macOS 原生的临时目录函数。
# <翻译结束>

