
<原文开始>
// Package devserver provides a web-based frontend so that
// it is possible to run a Wails app in a browsers.
<原文结束>

# <翻译开始>
// Package devserver 提供了一个基于Web的前端界面，以便于
// 在浏览器中运行Wails应用。
# <翻译结束>


<原文开始>
		// WebSockets aren't currently supported in prod mode, so a WebSocket connection is the result of the
		// FrontendDevServer e.g. Vite to support auto reloads.
		// Therefore we direct WebSockets directly to the FrontendDevServer instead of returning a NotImplementedStatus.
<原文结束>

# <翻译开始>
		// 目前在生产模式下不支持WebSockets，因此WebSocket连接是通过
		// 前端开发服务器（如Vite）建立的，目的是为了支持自动重载。
		// 因此，我们直接将WebSocket连接导向前端开发服务器，而不是返回一个未实现状态（NotImplementedStatus）。
# <翻译结束>


<原文开始>
// Setup internal dev server
<原文结束>

# <翻译开始>
// 设置内部开发服务器
# <翻译结束>


<原文开始>
// We do not support drag in browsers
<原文结束>

# <翻译开始>
// 我们不支持在浏览器中拖拽
# <翻译结束>


<原文开始>
// Notify the other browsers of "EventEmit"
<原文结束>

# <翻译开始>
// 通知其他浏览器关于"EventEmit"事件
# <翻译结束>


<原文开始>
// Send the message to dispatch to the frontend
<原文结束>

# <翻译开始>
// 将消息发送至调度程序以分发到前端
# <翻译结束>

