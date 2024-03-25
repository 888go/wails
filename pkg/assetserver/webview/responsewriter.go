package webview

import (
	"errors"
	"net/http"
)

const (
	HeaderContentLength = "Content-Length"
	HeaderContentType   = "Content-Type"
)

var (
	errRequestStopped   = errors.New("request has been stopped")
	errResponseFinished = errors.New("response has been finished")
)

// ResponseWriter 接口由 HTTP 处理程序使用，用于为 WebView 构造 HTTP 响应。
type ResponseWriter interface {
	http.ResponseWriter

	// 完成响应并刷新所有数据。如果请求已经完成，再次调用Finish将不会产生任何效果。
	Finish() error
}
