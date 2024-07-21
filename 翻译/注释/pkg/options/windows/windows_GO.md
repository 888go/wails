
<原文开始>
// SystemDefault will use whatever the system theme is. The application will follow system theme changes.
<原文结束>

# <翻译开始>
// SystemDefault 将使用系统当前的主题。应用程序会跟随系统主题的变化。
# <翻译结束>


<原文开始>
// ThemeSettings contains optional colours to use.
// They may be set using the hex values: 0x00BBGGRR
<原文结束>

# <翻译开始>
// ThemeSettings 包含可选的颜色设置。
// 可以使用十六进制值 0x00BBGGRR 来设置这些颜色。
# <翻译结束>


<原文开始>
// Options are options specific to Windows
<原文结束>

# <翻译开始>
// Options 是针对 Windows 的特定选项
# <翻译结束>


<原文开始>
	// Disable all window decorations in Frameless mode, which means no "Aero Shadow" and no "Rounded Corner" will be shown.
	// "Rounded Corners" are only available on Windows 11.
<原文结束>

# <翻译开始>
	// 在无边框模式下禁用所有窗口装饰，这意味着不会显示“Aero Shadow”和“圆角”。
	// “圆角”仅在Windows 11系统上可用。
# <翻译结束>


<原文开始>
	// Path where the WebView2 stores the user data. If empty %APPDATA%\[BinaryName.exe] will be used.
	// If the path is not valid, a messagebox will be displayed with the error and the app will exit with error code.
<原文结束>

# <翻译开始>
	// WebView2 存储用户数据的路径。如果为空，则使用 %APPDATA%\[BinaryName.exe]。
	// 如果路径无效，将显示一个包含错误信息的消息框，并且应用将以错误代码退出。
# <翻译结束>


<原文开始>
// Path to the directory with WebView2 executables. If empty WebView2 installed in the system will be used.
<原文结束>

# <翻译开始>
// WebView2可执行文件目录路径。如果为空，则使用系统已安装的WebView2。
# <翻译结束>


<原文开始>
// Dark/Light or System Default Theme
<原文结束>

# <翻译开始>
// 深色/浅色或系统默认主题
# <翻译结束>


<原文开始>
// Custom settings for dark/light mode
<原文结束>

# <翻译开始>
// 自定义暗黑模式和明亮模式的设置
# <翻译结束>


<原文开始>
// Select the type of translucent backdrop. Requires Windows 11 22621 or later.
<原文结束>

# <翻译开始>
// 选择半透明背景类型。需要Windows 11 22621或更高版本。
# <翻译结束>


<原文开始>
// User messages that can be customised
<原文结束>

# <翻译开始>
// 可自定义的用户消息
# <翻译结束>


<原文开始>
	// ResizeDebounceMS is the amount of time to debounce redraws of webview2
	// when resizing the window
<原文结束>

# <翻译开始>
	// ResizeDebounceMS 是在调整窗口大小时，对 webview2 重绘操作进行防抖动的延时时间（单位：毫秒）
# <翻译结束>


<原文开始>
// OnSuspend is called when Windows enters low power mode
<原文结束>

# <翻译开始>
// OnSuspend 在Windows进入低功耗模式时被调用
# <翻译结束>


<原文开始>
// OnResume is called when Windows resumes from low power mode
<原文结束>

# <翻译开始>
// OnResume 当Windows从低功耗模式恢复时被调用
# <翻译结束>


<原文开始>
// WebviewGpuIsDisabled is used to enable / disable GPU acceleration for the webview
<原文结束>

# <翻译开始>
// WebviewGpuIsDisabled 用于启用/禁用 webview 的 GPU 加速功能
# <翻译结束>


<原文开始>
	// WebviewDisableRendererCodeIntegrity disables the `RendererCodeIntegrity` of WebView2. Some Security Endpoint
	// Protection Software inject themself into the WebView2 with unsigned or wrongly signed dlls, which is not allowed
	// and will stop the WebView2 processes. Those security software need an update to fix this issue or one can disable
	// the integrity check with this flag.
	//
	// The event viewer log contains `Code Integrity Errors` like mentioned here: https://github.com/MicrosoftEdge/WebView2Feedback/issues/2051
	//
	// !! Please keep in mind when disabling this feature, this also allows malicious software to inject into the WebView2 !!
<原文结束>

# <翻译开始>
	// WebviewDisableRendererCodeIntegrity 禁用 WebView2 的 `RendererCodeIntegrity`。某些安全端点防护软件
	// 会使用未签名或签名错误的 dll 注入到 WebView2 中，这是不允许的，并且会导致 WebView2 进程停止运行。
	// 这类安全软件需要更新以解决此问题，或者可以通过设置此标志禁用完整性检查来暂时解决。
	//
	// Windows 事件查看器日志中包含如在 https:	//github.com/MicrosoftEdge/WebView2Feedback/issues/2051 中提及的 `代码完整性错误`。
	//
	// !! 请注意，禁用此功能时，也会允许恶意软件注入 WebView2，请谨慎操作 !!
# <翻译结束>


<原文开始>
// Configure whether swipe gestures should be enabled
<原文结束>

# <翻译开始>
// 配置是否启用滑动手势
# <翻译结束>

