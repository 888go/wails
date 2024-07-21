
<原文开始>
// Dlg forces display of focus rectangles, as soon as the user starts to type.
<原文结束>

# <翻译开始>
// Dlg 在用户开始输入时强制显示焦点矩形框。
# <翻译结束>


<原文开始>
//if !w.frontendOptions.Frameless {
<原文结束>

# <翻译开始>
// 如果!w.frontendOptions.Frameless 也就是说，如果w.frontendOptions.Frameless为false（非框架模式），则执行以下代码
# <翻译结束>


<原文开始>
			// If we want to have a frameless window but with the default frame decorations, extend the DWM client area.
			// This Option is not affected by returning 0 in WM_NCCALCSIZE.
			// As a result we have hidden the titlebar but still have the default window frame styling.
			// See: https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
<原文结束>

# <翻译开始>
			// 如果我们想要一个无边框的窗口，但保留默认的框架装饰样式，则扩展DWM客户端区域。
			// 此选项不受在WM_NCCALCSIZE消息中返回0的影响。
			// 结果是隐藏了标题栏，但仍保留了默认的窗口框架样式。
			// 参考：https:			//docs.microsoft.com/zh-cn/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
# <翻译结束>


<原文开始>
			// Disable the standard frame by allowing the client area to take the full
			// window size.
			// See: https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize#remarks
			// This hides the titlebar and also disables the resizing from user interaction because the standard frame is not
			// shown. We still need the WS_THICKFRAME style to enable resizing from the frontend.
<原文结束>

# <翻译开始>
			// 禁用标准边框，允许客户区占据整个窗口大小。
			// 参考：https:			//docs.microsoft.com/zh-cn/windows/win32/winmsg/wm-nccalcsize#remarks
			// 这将隐藏标题栏，并由于未显示标准边框，也将禁用用户交互的窗口调整大小功能。但我们仍需要WS_THICKFRAME样式来支持前端进行窗口调整大小的操作。
# <翻译结束>


<原文开始>
// In Full-Screen mode we don't need to adjust anything
<原文结束>

# <翻译开始>
// 在全屏模式下，我们无需进行任何调整
# <翻译结束>


<原文开始>
					// If the window is maximized we must adjust the client area to the work area of the monitor. Otherwise
					// some content goes beyond the visible part of the monitor.
					// Make sure to use the provided RECT to get the monitor, because during maximizig there might be
					// a wrong monitor returned in multi screen mode when using MonitorFromWindow.
					// See: https://github.com/MicrosoftEdge/WebView2Feedback/issues/2549
<原文结束>

# <翻译开始>
					// 如果窗口最大化，我们必须调整客户区以适应显示器的工作区。否则，
					// 有些内容会超出显示器的可见部分。
					// 确保使用提供的RECT来获取显示器，因为在多屏幕模式下最大化时，
					// 使用MonitorFromWindow可能会返回错误的显示器。
					// 参考：https:					//github.com/MicrosoftEdge/WebView2Feedback/issues/2549
# <翻译结束>


<原文开始>
					// This is needed to workaround the resize flickering in frameless mode with WindowDecorations
					// See: https://stackoverflow.com/a/6558508
					// The workaround originally suggests to decrese the bottom 1px, but that seems to bring up a thin
					// white line on some Windows-Versions, due to DrawBackground using also this reduces ClientSize.
					// Increasing the bottom also worksaround the flickering but we would loose 1px of the WebView content
					// therefore let's pad the content with 1px at the bottom.
<原文结束>

# <翻译开始>
					// 这是为了解决无边框模式（frameless mode）下使用WindowDecorations时出现的窗口大小调整闪烁问题
					// 参考：https:					//stackoverflow.com/a/6558508
					// 原始解决方案建议减小底部1px，但这在某些Windows版本上似乎会导致由于DrawBackground也使用了这一减少而导致底部出现一条细白线。
					// 增加底部尺寸同样可以规避闪烁问题，但我们会损失WebView内容的1px高度，因此我们选择在内容底部填充1px作为补偿。
# <翻译结束>

