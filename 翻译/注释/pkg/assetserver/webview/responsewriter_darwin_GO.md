
<原文开始>
			// This is very bad to detect a stopped schemeTask this should be implemented in a better way
			// But it seems to be very tricky to not deadlock when keeping a lock curing executing fn()
			// It seems like those call switch the thread back to the main thread and then deadlocks when they reentrant want
			// to get the lock again to start another request or stop it.
<原文结束>

# <翻译开始>
// 这是一种非常不好的方式来检测一个已停止的schemeTask，应该采用更好的实现方法
// 但在执行fn()函数时保持锁的状态下，似乎很难避免死锁
// 看起来这些调用会将线程切回主线程，然后在它们重新进入并再次尝试获取锁以开始另一个请求或停止任务时产生死锁
# <翻译结束>

