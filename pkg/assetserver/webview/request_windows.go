//go:build windows
// +build windows

package webview

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/wailsapp/go-webview2/pkg/edge"
)

// NewRequest 创建一个新的 WebViewRequest 用于 Chromium。这个方法必须在主线程中调用！

// ff:创建请求对象
// Request:
// fn:
func X创建请求对象(env *edge.ICoreWebView2Environment, args *edge.ICoreWebView2WebResourceRequestedEventArgs, invokeSync func(fn func())) (Request, error) {
	req, err := args.GetRequest()
	if err != nil {
		return nil, fmt.Errorf("GetRequest failed: %s", err)
	}
	defer req.Release()

	r := &request{
		invokeSync: invokeSync,
	}

	code := http.StatusInternalServerError
	r.response, err = env.CreateWebResourceResponse(nil, code, http.StatusText(code), "")
	if err != nil {
		return nil, fmt.Errorf("CreateWebResourceResponse failed: %s", err)
	}

	if err := args.PutResponse(r.response); err != nil {
		r.finishResponse()
		return nil, fmt.Errorf("PutResponse failed: %s", err)
	}

	r.deferral, err = args.GetDeferral()
	if err != nil {
		r.finishResponse()
		return nil, fmt.Errorf("GetDeferral failed: %s", err)
	}

	r.url, r.urlErr = req.GetUri()
	r.method, r.methodErr = req.GetMethod()
	r.header, r.headerErr = getHeaders(req)

	if content, err := req.GetContent(); err != nil {
		r.bodyErr = err
	} else if content != nil {
		// 从另一个线程安全地访问 Content：https://learn.microsoft.com/zh-cn/microsoft-edge/webview2/concepts/threading-model#线程安全性
// （该注释表明，在Go语言代码中，可以安全地在不同的线程中访问“Content”，其线程安全特性参照Microsoft Edge WebView2的线程模型文档中的“线程安全性”部分。）
		r.body = &iStreamReleaseCloser{stream: content}
	}

	return r, nil
}

var _ Request = &request{}

type request struct {
	response *edge.ICoreWebView2WebResourceResponse
	deferral *edge.ICoreWebView2Deferral

	url    string
	urlErr error

	method    string
	methodErr error

	header    http.Header
	headerErr error

	body    io.ReadCloser
	bodyErr error
	rw      *responseWriter

	invokeSync func(fn func())
}


// ff:
func (r *request) URL() (string, error) {
	return r.url, r.urlErr
}


// ff:请求方法
func (r *request) X请求方法() (string, error) {
	return r.method, r.methodErr
}


// ff:请求头
func (r *request) X请求头() (http.Header, error) {
	return r.header, r.headerErr
}


// ff:请求体
func (r *request) X请求体() (io.ReadCloser, error) {
	return r.body, r.bodyErr
}


// ff:请求响应
func (r *request) X请求响应() ResponseWriter {
	if r.rw != nil {
		return r.rw
	}

	r.rw = &responseWriter{req: r}
	return r.rw
}


// ff:关闭
func (r *request) X关闭() error {
	var errs []error
	if r.body != nil {
		if err := r.body.Close(); err != nil {
			errs = append(errs, err)
		}
		r.body = nil
	}

	if err := r.X请求响应().Finish(); err != nil {
		errs = append(errs, err)
	}

	return combineErrs(errs)
}

// finishResponse 必须在主线程上调用
func (r *request) finishResponse() error {
	var errs []error
	if r.response != nil {
		if err := r.response.Release(); err != nil {
			errs = append(errs, err)
		}
		r.response = nil
	}
	if r.deferral != nil {
		if err := r.deferral.Complete(); err != nil {
			errs = append(errs, err)
		}

		if err := r.deferral.Release(); err != nil {
			errs = append(errs, err)
		}
		r.deferral = nil
	}
	return combineErrs(errs)
}

type iStreamReleaseCloser struct {
	stream *edge.IStream
	closed bool
}


// ff:
// p:
func (i *iStreamReleaseCloser) Read(p []byte) (int, error) {
	if i.closed {
		return 0, io.ErrClosedPipe
	}
	return i.stream.Read(p)
}


// ff:
func (i *iStreamReleaseCloser) Close() error {
	if i.closed {
		return nil
	}
	i.closed = true
	return i.stream.Release()
}

func getHeaders(req *edge.ICoreWebView2WebResourceRequest) (http.Header, error) {
	header := http.Header{}
	headers, err := req.GetHeaders()
	if err != nil {
		return nil, fmt.Errorf("GetHeaders Error: %s", err)
	}
	defer headers.Release()

	headersIt, err := headers.GetIterator()
	if err != nil {
		return nil, fmt.Errorf("GetIterator Error: %s", err)
	}
	defer headersIt.Release()

	for {
		has, err := headersIt.HasCurrentHeader()
		if err != nil {
			return nil, fmt.Errorf("HasCurrentHeader Error: %s", err)
		}
		if !has {
			break
		}

		name, value, err := headersIt.GetCurrentHeader()
		if err != nil {
			return nil, fmt.Errorf("GetCurrentHeader Error: %s", err)
		}

		header.Set(name, value)
		if _, err := headersIt.MoveNext(); err != nil {
			return nil, fmt.Errorf("MoveNext Error: %s", err)
		}
	}

// 当WebView2接收到一个304状态码的请求时，可能会出现一些问题，
// 导致WebView2在处理其他请求（包括IPC调用）时会挂起。
// 为避免这种情况发生，通过移除与缓存配合使用的头部信息，防止返回304状态码。
	header.Del("If-Modified-Since")
	header.Del("If-None-Match")
	return header, nil
}

func combineErrs(errs []error) error {
	// TODO：使用Go1.20版本的errors.Join函数
	if len(errs) == 0 {
		return nil
	}

	errStrings := make([]string, len(errs))
	for i, err := range errs {
		errStrings[i] = err.Error()
	}

	return fmt.Errorf(strings.Join(errStrings, "\n"))
}
