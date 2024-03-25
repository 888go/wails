
<原文开始>
// newRequestFinalizer returns a request with a runtime finalizer to make sure it will be closed from the finalizer
// if it has not been already closed.
// It also makes sure Close() of the wrapping request is only called once.
<原文结束>

# <翻译开始>
// newRequestFinalizer 返回一个带有运行时终结器的请求，确保即使尚未被显式关闭，也能在终结器中进行关闭。
// 同时，它还确保包装请求的 Close() 方法仅被调用一次。
# <翻译结束>


<原文开始>
// Make sure to async release since it might block the finalizer goroutine for a longer period
<原文结束>

# <翻译开始>
// 确保异步释放，因为它可能会阻塞终结器 goroutine 较长时间
# <翻译结束>

