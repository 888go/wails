package assetserver

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/888go/wails/pkg/assetserver/webview"
)

type assetServerWebView struct {
	// ExpectedWebViewHost 用于检查每个 WebViewRequest 的请求主机，只有与该预期主机相匹配的请求才会被处理，其他主机的请求将不会被处理。
	ExpectedWebViewHost string

	dispatchInit    sync.Once
	dispatchReqC    chan<- webview.Request
	dispatchWorkers int
}

// ServeWebViewRequest 异步处理HTTP请求，通过模拟一个Golang HTTP服务器进行响应。
// 如果没有处理器向响应写入内容，则该请求将以StatusNotImplemented状态码完成。
// AssetServer将接管该请求，调用者不得以任何方式关闭它或访问它。
// 进一步解释：
// - `ServeWebViewRequest`函数用于异步地处理HTTP请求，通过模拟Go语言的HTTP服务端行为来实现。
// - 如果没有对应的处理器向客户端发送响应数据，则该HTTP请求将以501（未实现）的状态码结束。
// - `AssetServer`会获取对该HTTP请求的所有权，因此调用者不应再关闭这个请求或者以其他方式对其进行操作。
func (d *AssetServer) ServeWebViewRequest(req webview.Request) {
	d.dispatchInit.Do(func() {
		workers := d.dispatchWorkers
		if workers <= 0 {
			return
		}

		workerC := make(chan webview.Request, workers*2)
		for i := 0; i < workers; i++ {
			go func() {
				for req := range workerC {
					d.processWebViewRequest(req)
				}
			}()
		}

		dispatchC := make(chan webview.Request)
		go queueingDispatcher(50, dispatchC, workerC)

		d.dispatchReqC = dispatchC
	})

	if d.dispatchReqC == nil {
		go d.processWebViewRequest(req)
	} else {
		d.dispatchReqC <- req
	}
}

func (d *AssetServer) processWebViewRequest(r webview.Request) {
	uri, _ := r.URL()
	d.processWebViewRequestInternal(r)
	if err := r.Close(); err != nil {
		d.logError("Unable to call close for request for uri '%s'", uri)
	}
}

// processHTTPRequest 通过模拟一个 Go 语言 HTTP 服务器处理 HTTP 请求。
// 如果没有处理器向响应写入数据，则该请求将以 StatusNotImplemented 状态码完成。
func (d *AssetServer) processWebViewRequestInternal(r webview.Request) {
	uri := "unknown"
	var err error

	wrw := r.Response()
	defer func() {
		if err := wrw.Finish(); err != nil {
			d.logError("Error finishing request '%s': %s", uri, err)
		}
	}()

	var rw http.ResponseWriter = &contentTypeSniffer{rw: wrw} // 确保我们有一个内容类型探测器
	defer rw.WriteHeader(http.StatusNotImplemented)           // 当处理器已经写入并设置了状态时，这是一个空操作（NOP）

	uri, err = r.URL()
	if err != nil {
		d.logError("Error processing request, unable to get URL: %s (HttpResponse=500)", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	method, err := r.Method()
	if err != nil {
		d.webviewRequestErrorHandler(uri, rw, fmt.Errorf("HTTP-Method: %w", err))
		return
	}

	header, err := r.Header()
	if err != nil {
		d.webviewRequestErrorHandler(uri, rw, fmt.Errorf("HTTP-Header: %w", err))
		return
	}

	body, err := r.Body()
	if err != nil {
		d.webviewRequestErrorHandler(uri, rw, fmt.Errorf("HTTP-Body: %w", err))
		return
	}

	if body == nil {
		body = http.NoBody
	}
	defer body.Close()

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		d.webviewRequestErrorHandler(uri, rw, fmt.Errorf("HTTP-Request: %w", err))
		return
	}

// 对于服务器请求，URL是从Request-Line中存储在RequestURI的URI解析得到的。对于
// 大多数请求，除了Path和RawQuery字段外，其他字段将为空。（参见RFC 7230，第5.3节）
	req.URL.Scheme = ""
	req.URL.Host = ""
	req.URL.Fragment = ""
	req.URL.RawFragment = ""

	if url := req.URL; req.RequestURI == "" && url != nil {
		req.RequestURI = url.String()
	}

	req.Header = header

	if req.RemoteAddr == "" {
		// 192.0.2.0/24 在 RFC 5737 中定义为“TEST-NET”
		req.RemoteAddr = "192.0.2.1:1234"
	}

	if req.ContentLength == 0 {
		req.ContentLength, _ = strconv.ParseInt(req.Header.Get(HeaderContentLength), 10, 64)
	} else {
		size := strconv.FormatInt(req.ContentLength, 10)
		req.Header.Set(HeaderContentLength, size)
	}

	if host := req.Header.Get(HeaderHost); host != "" {
		req.Host = host
	}

	if expectedHost := d.ExpectedWebViewHost; expectedHost != "" && expectedHost != req.Host {
		d.webviewRequestErrorHandler(uri, rw, fmt.Errorf("expected host '%s' in request, but was '%s'", expectedHost, req.Host))
		return
	}

	d.ServeHTTP(rw, req)
}

func (d *AssetServer) webviewRequestErrorHandler(uri string, rw http.ResponseWriter, err error) {
	logInfo := uri
	if uri, err := url.ParseRequestURI(uri); err == nil {
		logInfo = strings.Replace(logInfo, fmt.Sprintf("%s://%s", uri.Scheme, uri.Host), "", 1)
	}

	d.logError("Error processing request '%s': %s (HttpResponse=500)", logInfo, err)
	http.Error(rw, err.Error(), http.StatusInternalServerError)
}

func queueingDispatcher[T any](minQueueSize uint, inC <-chan T, outC chan<- T) {
	q := newRingqueue[T](minQueueSize)
	for {
		in, ok := <-inC
		if !ok {
			return
		}

		q.Add(in)
		for q.Len() != 0 {
			out, _ := q.Peek()
			select {
			case outC <- out:
				q.Remove()
			case in, ok := <-inC:
				if !ok {
					return
				}

				q.Add(in)
			}
		}
	}
}
