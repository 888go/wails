
<原文开始>
// Get Windows build number
<原文结束>

# <翻译开始>
// 获取Windows构建号
# <翻译结束>


<原文开始>
// We currently can't use wails://wails/ as other platforms do, therefore we map the assets sever onto the following url.
<原文结束>

# <翻译开始>
// 目前我们不能像其他平台那样使用 wails://wails/，因此我们将 assets 服务器映射到以下 URL。
# <翻译结束>


<原文开始>
			// If the window is frameless and we are minimizing, then we need to suppress the Resize on the
			// WebView2. If we don't do this, restoring does not work as expected and first restores with some wrong
			// size during the restore animation and only fully renders when the animation is done. This highly
			// depends on the content in the WebView, see https://github.com/wailsapp/wails/issues/1319
<原文结束>

# <翻译开始>
			// 如果窗口是无边框的并且我们正在进行最小化操作，那么我们需要抑制WebView2上的Resize事件。如果不这样做，在恢复窗口大小时无法按预期工作，并且在恢复动画期间首先会以错误的尺寸还原，直到动画完成后才会完全渲染。这高度依赖于WebView中的内容，详细信息参见 https:			//github.com/wailsapp/wails/issues/1319
# <翻译结束>


<原文开始>
// WebView2 only has 0 and 255 as valid values.
<原文结束>

# <翻译开始>
// WebView2仅将0和255视为有效值。
# <翻译结束>


<原文开始>
	// Exit must be called on the Main-Thread. It calls PostQuitMessage which sends the WM_QUIT message to the thread's
	// message queue and our message queue runs on the Main-Thread.
<原文结束>

# <翻译开始>
	// Exit 必须在主线程上调用。它会调用 PostQuitMessage，该函数向线程的消息队列发送 WM_QUIT 消息，
	// 而我们的消息队列是在主线程上运行的。
# <翻译结束>


<原文开始>
// Check if CTRL is pressed
<原文结束>

# <翻译开始>
// 检查是否按下了CTRL键
# <翻译结束>


<原文开始>
// => The app has to recreate a new WebView to recover from this failure.
<原文结束>

# <翻译开始>
// => 为了从这个故障中恢复，应用必须重新创建一个新的WebView。
# <翻译结束>


<原文开始>
			// => A new render process is created automatically and navigated to an error page.
			// => Make sure that the error page is shown.
<原文结束>

# <翻译开始>
			// => 自动创建一个新的渲染进程，并导航到错误页面。
			// => 确保错误页面被展示出来。
# <翻译结束>


<原文开始>
// NavgiationCompleted didn't come in, make sure the chromium is shown
<原文结束>

# <翻译开始>
// NavgiationCompleted 事件未触发，确保 Chromium 正在显示
# <翻译结束>


<原文开始>
// The window has never been shown, make sure to show it
<原文结束>

# <翻译开始>
// 窗口从未被显示过，确保将其显示出来
# <翻译结束>


<原文开始>
// Setup focus event handler
<原文结束>

# <翻译开始>
// 设置焦点事件处理器
# <翻译结束>


<原文开始>
	// Setting the UserAgent on the CoreWebView2Settings clears the whole default UserAgent of the Edge browser, but
	// we want to just append our ApplicationIdentifier. So we adjust the UserAgent for every request.
<原文结束>

# <翻译开始>
	// 在CoreWebView2Settings上设置UserAgent会清空Edge浏览器的整个默认UserAgent，
	// 但我们只想追加我们的ApplicationIdentifier。因此，我们对每个请求调整UserAgent。
# <翻译结束>


<原文开始>
// We are using the devServer let the WebView2 handle the request with its default handler
<原文结束>

# <翻译开始>
// 我们使用devServer让WebView2通过其默认处理器来处理请求
# <翻译结束>


<原文开始>
// Let the WebView2 handle the request with its default handler
<原文结束>

# <翻译开始>
// 让WebView2使用其默认处理器处理请求
# <翻译结束>


<原文开始>
// Callback from a method call
<原文结束>

# <翻译开始>
// 从方法调用返回的回调函数
# <翻译结束>


<原文开始>
// Use PostMessage because we don't want to block the caller until dragging has been finished.
<原文结束>

# <翻译开始>
// 使用PostMessage是因为我们不希望在拖拽操作完成之前阻塞调用者。
# <翻译结束>


<原文开始>
// Use PostMessage because we don't want to block the caller until resizing has been finished.
<原文结束>

# <翻译开始>
// 使用PostMessage是因为我们不希望在调整大小完成之前阻塞调用者。
# <翻译结束>


<原文开始>
// Hack to make it visible: https://github.com/MicrosoftEdge/WebView2Feedback/issues/1077#issuecomment-825375026
<原文结束>

# <翻译开始>
// 临时解决方案以使其可见：https://github.com/MicrosoftEdge/WebView2Feedback/issues/1077#issuecomment-825375026
// （该段英文注释描述了一个临时性的解决方案，用于解决某个特定问题以达到使其可见的目的。具体问题和方案请参考链接中的GitHub讨论，该讨论位于MicrosoftEdge/WebView2Feedback仓库的第1077号Issue中的一条编号为825375026的评论。）
# <翻译结束>

