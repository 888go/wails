//go:build linux
// +build linux

package webview

/*
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0 gio-unix-2.0

#include "gtk/gtk.h"
#include "webkit2/webkit2.h"
*/
import "C"

import (
	"io"
	"net/http"
	"unsafe"
)

// NewRequest 创建一个新的 WebViewRequest，基于 `WebKitURISchemeRequest` 指针
// 这个函数根据给定的 WebKitURISchemeRequest 指针创建一个新的 WebViewRequest 对象

// ff:创建请求对象
// webKitURISchemeRequest:
func NewRequest(webKitURISchemeRequest unsafe.Pointer) Request {
	webkitReq := (*C.WebKitURISchemeRequest)(webKitURISchemeRequest)
	C.g_object_ref(C.gpointer(webkitReq))

	req := &request{req: webkitReq}
	return newRequestFinalizer(req)
}

var _ Request = &request{}

type request struct {
	req *C.WebKitURISchemeRequest

	header http.Header
	body   io.ReadCloser
	rw     *responseWriter
}


// ff:
func (r *request) URL() (string, error) {
	return C.GoString(C.webkit_uri_scheme_request_get_uri(r.req)), nil
}


// ff:请求方法
func (r *request) Method() (string, error) {
	return webkit_uri_scheme_request_get_http_method(r.req), nil
}


// ff:请求头
func (r *request) Header() (http.Header, error) {
	if r.header != nil {
		return r.header, nil
	}

	r.header = webkit_uri_scheme_request_get_http_headers(r.req)
	return r.header, nil
}


// ff:请求体
func (r *request) Body() (io.ReadCloser, error) {
	if r.body != nil {
		return r.body, nil
	}

	r.body = webkit_uri_scheme_request_get_http_body(r.req)

	return r.body, nil
}


// ff:请求响应
func (r *request) Response() ResponseWriter {
	if r.rw != nil {
		return r.rw
	}

	r.rw = &responseWriter{req: r.req}
	return r.rw
}


// ff:关闭
func (r *request) Close() error {
	var err error
	if r.body != nil {
		err = r.body.Close()
	}
	r.Response().Finish()
	C.g_object_unref(C.gpointer(r.req))
	return err
}
