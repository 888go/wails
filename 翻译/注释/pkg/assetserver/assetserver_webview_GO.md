
<原文开始>
// ExpectedWebViewHost is checked against the Request Host of every WebViewRequest, other hosts won't be processed.
<原文结束>

# <翻译开始>
// ExpectedWebViewHost 用于检查每个 WebViewRequest 的请求主机，只有与该预期主机相匹配的请求才会被处理，其他主机的请求将不会被处理。
# <翻译结束>


<原文开始>
// ServeWebViewRequest processes the HTTP Request asynchronously by faking a golang HTTP Server.
// The request will be finished with a StatusNotImplemented code if no handler has written to the response.
// The AssetServer takes ownership of the request and the caller mustn't close it or access it in any other way.
<原文结束>

# <翻译开始>
// ServeWebViewRequest 异步处理HTTP请求，通过模拟一个Golang HTTP服务器进行响应。
// 如果没有处理器向响应写入内容，则该请求将以StatusNotImplemented状态码完成。
// AssetServer将接管该请求，调用者不得以任何方式关闭它或访问它。
// 进一步解释：
// - `ServeWebViewRequest`函数用于异步地处理HTTP请求，通过模拟Go语言的HTTP服务端行为来实现。
// - 如果没有对应的处理器向客户端发送响应数据，则该HTTP请求将以501（未实现）的状态码结束。
// - `AssetServer`会获取对该HTTP请求的所有权，因此调用者不应再关闭这个请求或者以其他方式对其进行操作。
# <翻译结束>


<原文开始>
// processHTTPRequest processes the HTTP Request by faking a golang HTTP Server.
// The request will be finished with a StatusNotImplemented code if no handler has written to the response.
<原文结束>

# <翻译开始>
// processHTTPRequest 通过模拟一个 Go 语言 HTTP 服务器处理 HTTP 请求。
// 如果没有处理器向响应写入数据，则该请求将以 StatusNotImplemented 状态码完成。
# <翻译结束>


<原文开始>
// Make sure we have a Content-Type sniffer
<原文结束>

# <翻译开始>
// 确保我们有一个内容类型探测器
# <翻译结束>


<原文开始>
// This is a NOP when a handler has already written and set the status
<原文结束>

# <翻译开始>
// 当处理器已经写入并设置了状态时，这是一个空操作（NOP）
# <翻译结束>


<原文开始>
	// For server requests, the URL is parsed from the URI supplied on the Request-Line as stored in RequestURI. For
	// most requests, fields other than Path and RawQuery will be empty. (See RFC 7230, Section 5.3)
<原文结束>

# <翻译开始>
	// 对于服务器请求，URL是从Request-Line中存储在RequestURI的URI解析得到的。对于
	// 大多数请求，除了Path和RawQuery字段外，其他字段将为空。（参见RFC 7230，第5.3节）
# <翻译结束>


<原文开始>
// 192.0.2.0/24 is "TEST-NET" in RFC 5737
<原文结束>

# <翻译开始>
// 192.0.2.0/24 在 RFC 5737 中定义为“TEST-NET”
# <翻译结束>

