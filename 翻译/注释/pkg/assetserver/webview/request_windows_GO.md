
<原文开始>
// NewRequest creates as new WebViewRequest for chromium. This Method must be called from the Main-Thread!
<原文结束>

# <翻译开始>
// NewRequest 创建一个新的 WebViewRequest 用于 Chromium。这个方法必须在主线程中调用！
# <翻译结束>


<原文开始>
// It is safe to access Content from another Thread: https://learn.microsoft.com/en-us/microsoft-edge/webview2/concepts/threading-model#thread-safety
<原文结束>

# <翻译开始>
// 从另一个线程安全地访问 Content：https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/concepts/threading-model#线程安全性
// （该注释表明，在Go语言代码中，可以安全地在不同的线程中访问“Content”，其线程安全特性参照Microsoft Edge WebView2的线程模型文档中的“线程安全性”部分。）
# <翻译结束>


<原文开始>
// finishResponse must be called on the main-thread
<原文结束>

# <翻译开始>
// finishResponse 必须在主线程上调用
# <翻译结束>


<原文开始>
	// WebView2 has problems when a request returns a 304 status code and the WebView2 is going to hang for other
	// requests including IPC calls.
	// So prevent 304 status codes by removing the headers that are used in combinationwith caching.
<原文结束>

# <翻译开始>
// 当WebView2接收到一个304状态码的请求时，可能会出现一些问题，
// 导致WebView2在处理其他请求（包括IPC调用）时会挂起。
// 为避免这种情况发生，通过移除与缓存配合使用的头部信息，防止返回304状态码。
# <翻译结束>


<原文开始>
// TODO use Go1.20 errors.Join
<原文结束>

# <翻译开始>
// TODO：使用Go1.20版本的errors.Join函数
# <翻译结束>

