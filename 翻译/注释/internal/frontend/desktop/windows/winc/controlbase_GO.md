
<原文开始>
// use MouseControl to capture onMouseHover and onMouseLeave events.
<原文结束>

# <翻译开始>
// 使用MouseControl来捕获onMouseHover和onMouseLeave事件。
# <翻译结束>


<原文开始>
// initControl is called by controls: edit, button, treeview, listview, and so on.
<原文结束>

# <翻译开始>
// initControl 由控件调用，如：编辑框、按钮、树视图、列表视图等。
# <翻译结束>


<原文开始>
// InitWindow is called by custom window based controls such as split, panel, etc.
<原文结束>

# <翻译开始>
// InitWindow 由自定义窗口组件（如 split、panel 等）调用。
# <翻译结束>


<原文开始>
// SetTheme for TreeView and ListView controls.
<原文结束>

# <翻译开始>
// SetTheme 为 TreeView 和 ListView 控件设置主题。
# <翻译结束>


<原文开始>
		// GetDpiForWindow is supported beginning with Windows 10, 1607 and is the most accureate
		// one, especially it is consistent with the WM_DPICHANGED event.
<原文结束>

# <翻译开始>
		// GetDpiForWindow 函数从 Windows 10，版本1607开始支持，并且是最精确的一个，
		// 特别是它与 WM_DPICHANGED 事件保持一致。
# <翻译结束>


<原文开始>
// GetDpiForWindow is supported beginning with Windows 8.1
<原文结束>

# <翻译开始>
// GetDpiForWindow 从 Windows 8.1 开始支持
# <翻译结束>


<原文开始>
// If none of the above is supported fallback to the System DPI.
<原文结束>

# <翻译开始>
// 如果以上都不支持，则回退到系统DPI。
# <翻译结束>


<原文开始>
// Ensure we set max if min > max
<原文结束>

# <翻译开始>
// 确保当min大于max时，我们设置max
# <翻译结束>


<原文开始>
// Ensure we set min if max > min
<原文结束>

# <翻译开始>
// 确保当max大于min时，我们设置min的值
# <翻译结束>


<原文开始>
	// WindowPos is used with HWND_TOPMOST to guarantee bring our app on top
	// force set our main window on top
<原文结束>

# <翻译开始>
	// WindowPos 与 HWND_TOPMOST 一起使用，用于确保将我们的应用置于顶部
	// 强制设置我们的主窗口置于顶层
# <翻译结束>


<原文开始>
// remove topmost to allow normal windows manipulations
<原文结束>

# <翻译开始>
// 移除最顶层以允许进行常规窗口操作
# <翻译结束>


<原文开始>
// put main window on tops foreground
<原文结束>

# <翻译开始>
// 将主窗口置于顶部并激活（使其成为前景窗口）
# <翻译结束>


<原文开始>
	// pRect := w32.GetClientRect(cba.hwnd)
	// if cba.isForm {
	// 	w32.InvalidateRect(cba.hwnd, pRect, erase)
	// } else {
	// 	rc := ScreenToClientRect(cba.hwnd, pRect)
	// 	w32.InvalidateRect(cba.hwnd, rc.GetW32Rect(), erase)
	// }
<原文结束>

# <翻译开始>
	// 获取cba.hwnd的客户区矩形，并将其赋值给pRect
	// 如果cba.isForm为真（即表示当前是表单）
	// 则调用w32.InvalidateRect函数，传入cba.hwnd和pRect，以及erase参数，以更新该窗口区域
	// 否则（即表示当前不是表单）
	// 将pRect从屏幕坐标转换为客户区坐标，得到rc
	// 然后调用w32.InvalidateRect函数，传入cba.hwnd、rc转换后的W32矩形结构体，以及erase参数，以更新该窗口区域
	// 在上述代码中：
	// - `GetClientRect`：获取窗口的客户区矩形。
	// - `InvalidateRect`：使指定窗口区域无效，导致系统重绘该区域。
	// - `ScreenToClientRect`：将屏幕坐标点或矩形转换为客户区坐标点或矩形。
	// - `erase` 参数是一个布尔值，决定是否擦除窗口区域并重绘背景。
# <翻译结束>

