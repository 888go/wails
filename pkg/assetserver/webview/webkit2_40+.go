//go:build linux && webkit2_40

package webview

/*
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0 gio-unix-2.0

#include "gtk/gtk.h"
#include "webkit2/webkit2.h"
#include "gio/gunixinputstream.h"
*/
import "C"

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

func webkit_uri_scheme_request_get_http_body(req *C.WebKitURISchemeRequest) io.ReadCloser {
	stream := C.webkit_uri_scheme_request_get_http_body(req)
	if stream == nil {
		return http.NoBody
	}
	return &webkitRequestBody{stream: stream}
}

type webkitRequestBody struct {
	stream *C.GInputStream
	closed bool
}

// Read 实现了 io.Reader 接口

// ff:
// p:
func (r *webkitRequestBody) Read(p []byte) (int, error) {
	if r.closed {
		return 0, io.ErrClosedPipe
	}

	var content unsafe.Pointer
	var contentLen int
	if p != nil {
		content = unsafe.Pointer(&p[0])
		contentLen = len(p)
	}

	var n C.gsize
	var gErr *C.GError
	res := C.g_input_stream_read_all(r.stream, content, C.gsize(contentLen), &n, nil, &gErr)
	if res == 0 {
		return 0, formatGError("stream read failed", gErr)
	} else if n == 0 {
		return 0, io.EOF
	}
	return int(n), nil
}


// ff:关闭
func (r *webkitRequestBody) X关闭() error {
	if r.closed {
		return nil
	}
	r.closed = true

// https://docs.gtk.org/gio/method.InputStream.close.html
// 当最后一个引用被释放时，流会自动关闭，但您可能希望调用此函数，
// 以确保资源尽早释放。
	var err error
	var gErr *C.GError
	if C.g_input_stream_close(r.stream, nil, &gErr) == 0 {
		err = formatGError("stream close failed", gErr)
	}
	C.g_object_unref(C.gpointer(r.stream))
	r.stream = nil
	return err
}

func formatGError(msg string, gErr *C.GError, args ...any) error {
	if gErr != nil && gErr.message != nil {
		msg += ": " + C.GoString(gErr.message)
		C.g_error_free(gErr)
	}
	return fmt.Errorf(msg, args...)
}
