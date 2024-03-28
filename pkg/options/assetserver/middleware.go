package assetserver

import (
	"net/http"
)

// Middleware 定义了一个可以应用于 AssetServer 的 HTTP 中间件。
// 作为参数传递的 next 处理器是链中的下一个处理器。开发者可以根据需要决定是否调用下一个处理器，
// 或者实现一种特殊处理方式。
type Middleware func(next http.Handler) http.Handler

// ChainMiddleware 允许将多个中间件链接到一个中间件。

// ff:
// middleware:
func ChainMiddleware(middleware ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(middleware) - 1; i >= 0; i-- {
			h = middleware[i](h)
		}
		return h
	}
}
