package webview

import (
	"io"
	"net/http"
)

type Request interface {
	URL() (string, error)
	Method() (string, error) //hs:请求方法     
	Header() (http.Header, error) //hs:请求头     
	Body() (io.ReadCloser, error) //hs:请求体     

	Response() ResponseWriter //hs:请求响应     

	Close() error //hs:关闭     
}
