//go:build windows
// +build windows

package webview

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

var _ http.ResponseWriter = &responseWriter{}

type responseWriter struct {
	req *request

	header      http.Header
	wroteHeader bool
	code        int
	body        *bytes.Buffer

	finished bool
}


// ff:
func (rw *responseWriter) Header() http.Header {
	if rw.header == nil {
		rw.header = http.Header{}
	}
	return rw.header
}


// ff:
// buf:
func (rw *responseWriter) Write(buf []byte) (int, error) {
	if rw.finished {
		return 0, errResponseFinished
	}

	rw.WriteHeader(http.StatusOK)

	return rw.body.Write(buf)
}


// ff:
// code:
func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader || rw.finished {
		return
	}
	rw.wroteHeader = true

	if rw.body == nil {
		rw.body = &bytes.Buffer{}
	}

	rw.code = code
}


// ff:
func (rw *responseWriter) Finish() error {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusNotImplemented)
	}

	if rw.finished {
		return nil
	}
	rw.finished = true

	var errs []error

	code := rw.code
	if code == http.StatusNotModified {
// 当WebView2接收到一个返回304状态码的请求时，可能会出现问题，导致WebView2在处理其他请求（包括IPC调用）时出现卡死现象。
		errs = append(errs, fmt.Errorf("AssetServer returned 304 - StatusNotModified which are going to hang WebView2, changed code to 505 - StatusInternalServerError"))
		code = http.StatusInternalServerError
	}

	rw.req.invokeSync(func() {
		resp := rw.req.response

		hdrs, err := resp.GetHeaders()
		if err != nil {
			errs = append(errs, fmt.Errorf("Resp.GetHeaders failed: %s", err))
		} else {
			for k, v := range rw.header {
				if err := hdrs.AppendHeader(k, strings.Join(v, ",")); err != nil {
					errs = append(errs, fmt.Errorf("Resp.AppendHeader failed: %s", err))
				}
			}
			hdrs.Release()
		}

		if err := resp.PutStatusCode(code); err != nil {
			errs = append(errs, fmt.Errorf("Resp.PutStatusCode failed: %s", err))
		}

		if err := resp.PutByteContent(rw.body.Bytes()); err != nil {
			errs = append(errs, fmt.Errorf("Resp.PutByteContent failed: %s", err))
		}

		if err := rw.req.finishResponse(); err != nil {
			errs = append(errs, fmt.Errorf("Resp.finishResponse failed: %s", err))
		}
	})

	return combineErrs(errs)
}
