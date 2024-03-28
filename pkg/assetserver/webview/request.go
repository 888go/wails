package webview

import (
	"io"
	"net/http"
)

type Request interface {
	URL() (string, error)
	X请求方法() (string, error) //hs:请求方法     
	X请求头() (http.Header, error) //hs:请求头     
	X请求体() (io.ReadCloser, error) //hs:请求体     

	X请求响应() ResponseWriter //hs:请求响应     

	X关闭() error //hs:关闭     
}
