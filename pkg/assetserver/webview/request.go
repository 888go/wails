package webview

import (
	"io"
	"net/http"
)

type Request interface {
	URL() (string, error)
	X请求方法() (string, error)
	X请求头() (http.Header, error)
	X请求体() (io.ReadCloser, error)

	X请求响应() ResponseWriter

	X关闭() error
}
