
<原文开始>
// WebviewGpuPolicy values used for determining the webview's hardware acceleration policy.
<原文结束>

# <翻译开始>
// WebviewGpuPolicy 用于确定 webview 硬件加速策略的值。
# <翻译结束>


<原文开始>
// WebviewGpuPolicyAlways Hardware acceleration is always enabled.
<原文结束>

# <翻译开始>
// WebviewGpuPolicyAlways 硬件加速始终启用。
# <翻译结束>


<原文开始>
// WebviewGpuPolicyOnDemand Hardware acceleration is enabled/disabled as request by web contents.
<原文结束>

# <翻译开始>
// WebviewGpuPolicyOnDemand 当网页内容请求时，硬件加速启用/禁用。
# <翻译结束>


<原文开始>
// WebviewGpuPolicyNever Hardware acceleration is always disabled.
<原文结束>

# <翻译开始>
// WebviewGpuPolicyNever 硬件加速始终禁用。
# <翻译结束>


<原文开始>
// Options specific to Linux builds
<原文结束>

# <翻译开始>
// 以下选项针对Linux系统构建时特定
# <翻译结束>


<原文开始>
	// Icon Sets up the icon representing the window. This icon is used when the window is minimized
	// (also known as iconified).
<原文结束>

# <翻译开始>
	// Icon 设置代表窗口的图标。当窗口被最小化（也称为图标化）时，会使用这个图标。
# <翻译结束>


<原文开始>
// WindowIsTranslucent sets the window's background to transparent when enabled.
<原文结束>

# <翻译开始>
// WindowIsTranslucent 在启用时将窗口背景设置为透明
# <翻译结束>


<原文开始>
// Messages are messages that can be customised
<原文结束>

# <翻译开始>
// Messages 是可以自定义的消息
# <翻译结束>


<原文开始>
	// WebviewGpuPolicy used for determining the hardware acceleration policy for the webview.
	//   - WebviewGpuPolicyAlways
	//   - WebviewGpuPolicyOnDemand
	//   - WebviewGpuPolicyNever
	//
	// Due to https://github.com/wailsapp/wails/issues/2977, if options.Linux is nil
	// in the call to wails.Run(), WebviewGpuPolicy is set by default to WebviewGpuPolicyNever.
	// Client code may override this behavior by passing a non-nil Options and set
	// WebviewGpuPolicy as needed.
<原文结束>

# <翻译开始>
	// WebviewGpuPolicy 用于确定 webview 的硬件加速策略。
	//   - WebviewGpuPolicyAlways   	// 始终启用硬件加速
	//   - WebviewGpuPolicyOnDemand 	// 按需启用硬件加速
	//   - WebviewGpuPolicyNever    	// 从不启用硬件加速
	// 由于 https:	//github.com/wailsapp/wails/issues/2977 中的问题，如果在调用 wails.Run() 时 options.Linux 为 nil，
	// 则默认将 WebviewGpuPolicy 设置为 WebviewGpuPolicyNever。
	// 客户端代码可以通过传递非空的 Options 并按需设置 WebviewGpuPolicy 来覆盖此默认行为。
# <翻译结束>


<原文开始>
	// ProgramName is used to set the program's name for the window manager via GTK's g_set_prgname().
	//This name should not be localized. [see the docs]
	//
	//When a .desktop file is created this value helps with window grouping and desktop icons when the .desktop file's Name
	//property differs form the executable's filename.
	//
	//[see the docs]: https://docs.gtk.org/glib/func.set_prgname.html
<原文结束>

# <翻译开始>
	// ProgramName 用于通过 GTK 的 g_set_prgname() 函数设置程序名称，以便在窗口管理器中显示。
	// 此名称不应进行本地化处理。[参见文档]
	//
	// 当创建 .desktop 文件时，如果 .desktop 文件的 Name 属性与可执行文件名不同，
	// 这个值有助于窗口分组和桌面图标功能。
	//
	// [参见文档]: https:	//docs.gtk.org/glib/func.set_prgname.html
# <翻译结束>

