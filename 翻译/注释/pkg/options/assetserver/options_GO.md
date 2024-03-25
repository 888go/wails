
<原文开始>
// Options defines the configuration of the AssetServer.
<原文结束>

# <翻译开始>
// Options 定义了 AssetServer 的配置。
# <翻译结束>


<原文开始>
	// Assets defines the static assets to be used. A GET request is first tried to be served from this Assets. If the Assets returns
	// `os.ErrNotExist` for that file, the request handling will fallback to the Handler and tries to serve the GET
	// request from it.
	//
	// If set to nil, all GET requests will be forwarded to Handler.
<原文结束>

# <翻译开始>
// Assets 定义了要使用的静态资源。对于 GET 请求，首先尝试从这个 Assets 提供服务。
// 如果Assets针对该文件返回`os.ErrNotExist`错误，则请求处理将回退到Handler，并尝试从中提供GET请求的服务。
//
// 如果设置为nil，所有GET请求都将转发到Handler。
# <翻译结束>


<原文开始>
	// Handler will be called for every GET request that can't be served from Assets, due to `os.ErrNotExist`. Furthermore all
	// non GET requests will always be served from this Handler.
	//
	// If not defined, the result is the following in cases where the Handler would have been called:
	//   GET request:   `http.StatusNotFound`
	//   Other request: `http.StatusMethodNotAllowed`
<原文结束>

# <翻译开始>
// Handler 函数会在以下两种情况下被调用：
// 1. 当接收到 GET 请求但无法通过 Assets 服务提供资源时（由于返回错误 `os.ErrNotExist`）；
// 2. 对于所有非 GET 类型的请求，该 Handler 总会被调用以处理请求。
// 若未定义此 Handler，则在原本应当调用 Handler 的场景下，其默认行为如下：
//   - GET 请求：返回状态码 `http.StatusNotFound`（404，未找到）
//   - 非 GET 请求：返回状态码 `http.StatusMethodNotAllowed`（405，不允许的方法）
# <翻译结束>


<原文开始>
	// Middleware is a HTTP Middleware which allows to hook into the AssetServer request chain. It allows to skip the default
	// request handler dynamically, e.g. implement specialized Routing etc.
	// The Middleware is called to build a new `http.Handler` used by the AssetSever and it also receives the default
	// handler used by the AssetServer as an argument.
	//
	// If not defined, the default AssetServer request chain is executed.
	//
	// Multiple Middlewares can be chained together with:
	//   ChainMiddleware(middleware ...Middleware) Middleware
<原文结束>

# <翻译开始>
// Middleware 是一个HTTP中间件，它允许接入AssetServer请求处理链。它支持动态跳过默认请求处理器，例如实现特殊路由等功能。
// 当构建AssetServer所使用的新的`http.Handler`时会调用此Middleware，并且它还会接收到AssetServer所使用的默认处理器作为参数。
//
// 若未定义，则执行AssetServer的默认请求处理链。
//
// 可以通过以下方式将多个Middleware串联起来：
//   ChainMiddleware(middleware ...Middleware) Middleware
# <翻译结束>

