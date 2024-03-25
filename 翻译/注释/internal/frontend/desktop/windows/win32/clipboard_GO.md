
<原文开始>
// waitOpenClipboard opens the clipboard, waiting for up to a second to do so.
<原文结束>

# <翻译开始>
// waitOpenClipboard 尝试打开剪贴板，最多等待1秒钟以完成此操作。
# <翻译结束>


<原文开始>
	// LockOSThread ensure that the whole method will keep executing on the same thread from begin to end (it actually locks the goroutine thread attribution).
	// Otherwise if the goroutine switch thread during execution (which is a common practice), the OpenClipboard and CloseClipboard will happen on two different threads, and it will result in a clipboard deadlock.
<原文结束>

# <翻译开始>
// LockOSThread 确保整个方法从开始到结束都在同一个线程上执行（实际上是锁定了 goroutine 的线程归属）。
// 否则，如果 goroutine 在执行过程中切换了线程（这是常见情况），OpenClipboard 和 CloseClipboard 将在两个不同的线程上发生，这将导致剪贴板死锁。
# <翻译结束>


<原文开始>
	// "If the hMem parameter identifies a memory object, the object must have
	// been allocated using the function with the GMEM_MOVEABLE flag."
<原文结束>

# <翻译开始>
// "如果hMem参数标识了一个内存对象，则该对象必须使用带有GMEM_MOVEABLE标志的函数分配。"
# <翻译结束>


<原文开始>
// suppress deferred cleanup
<原文结束>

# <翻译开始>
// 抑制延迟清理
# <翻译结束>

