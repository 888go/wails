//go:build windows

package windows

import (
	"github.com/wailsapp/go-webview2/pkg/edge"
	"sync"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/win32"
	"github.com/wailsapp/wails/v2/internal/system/operatingsystem"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc"
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	winoptions "github.com/wailsapp/wails/v2/pkg/options/windows"
)

type Window struct {
	winc.Form
	frontendOptions                          *options.App
	applicationMenu                          *menu.Menu
	minWidth, minHeight, maxWidth, maxHeight int
	versionInfo                              *operatingsystem.WindowsVersionInfo
	isDarkMode                               bool
	isActive                                 bool
	hasBeenShown                             bool

	// Theme
	theme        winoptions.Theme
	themeChanged bool

	framelessWithDecorations bool

	OnSuspend func()
	OnResume  func()

	chromium *edge.Chromium
}

func NewWindow(parent winc.Controller, appoptions *options.App, versionInfo *operatingsystem.WindowsVersionInfo, chromium *edge.Chromium) *Window {
	windowsOptions := appoptions.Windows

	result := &Window{
		frontendOptions: appoptions,
		minHeight:       appoptions.MinHeight,
		minWidth:        appoptions.MinWidth,
		maxHeight:       appoptions.MaxHeight,
		maxWidth:        appoptions.MaxWidth,
		versionInfo:     versionInfo,
		isActive:        true,
		themeChanged:    true,
		chromium:        chromium,

		framelessWithDecorations: appoptions.Frameless && (windowsOptions == nil || !windowsOptions.DisableFramelessWindowDecorations),
	}
	result.SetIsForm(true)

	var exStyle int
	if windowsOptions != nil {
		exStyle = w32.WS_EX_CONTROLPARENT | w32.WS_EX_APPWINDOW
		if windowsOptions.WindowIsTranslucent {
			exStyle |= w32.WS_EX_NOREDIRECTIONBITMAP
		}
	}
	if appoptions.AlwaysOnTop {
		exStyle |= w32.WS_EX_TOPMOST
	}

	var dwStyle = w32.WS_OVERLAPPEDWINDOW

	winc.RegClassOnlyOnce("wailsWindow")
	handle := winc.CreateWindow("wailsWindow", parent, uint(exStyle), uint(dwStyle))
	result.SetHandle(handle)
	winc.RegMsgHandler(result)
	result.SetParent(parent)

	loadIcon := true
	if windowsOptions != nil && windowsOptions.DisableWindowIcon == true {
		loadIcon = false
	}
	if loadIcon {
		if ico, err := winc.NewIconFromResource(winc.GetAppInstance(), uint16(winc.AppIconID)); err == nil {
			result.SetIcon(0, ico)
		}
	}

	if appoptions.BackgroundColour != nil {
		win32.SetBackgroundColour(result.Handle(), appoptions.BackgroundColour.R, appoptions.BackgroundColour.G, appoptions.BackgroundColour.B)
	}

	if windowsOptions != nil {
		result.theme = windowsOptions.Theme
	} else {
		result.theme = winoptions.SystemDefault
	}

	result.SetSize(appoptions.Width, appoptions.Height)
	result.SetText(appoptions.Title)
	result.EnableSizable(!appoptions.DisableResize)
	if !appoptions.Fullscreen {
		result.EnableMaxButton(!appoptions.DisableResize)
		result.SetMinSize(appoptions.MinWidth, appoptions.MinHeight)
		result.SetMaxSize(appoptions.MaxWidth, appoptions.MaxHeight)
	}

	result.UpdateTheme()

	if windowsOptions != nil {
		result.OnSuspend = windowsOptions.OnSuspend
		result.OnResume = windowsOptions.OnResume
		if windowsOptions.WindowIsTranslucent {
			if !win32.SupportsBackdropTypes() {
				result.SetTranslucentBackground()
			} else {
				win32.EnableTranslucency(result.Handle(), win32.BackdropType(windowsOptions.BackdropType))
			}
		}

		if windowsOptions.DisableWindowIcon {
			result.DisableIcon()
		}
	}

	// Dlg 在用户开始输入时强制显示焦点矩形框。
	w32.SendMessage(result.Handle(), w32.WM_CHANGEUISTATE, w32.UIS_INITIALIZE, 0)

	result.SetFont(winc.DefaultFont)

	if appoptions.Menu != nil {
		result.SetApplicationMenu(appoptions.Menu)
	}

	return result
}

func (w *Window) Fullscreen() {
	if w.Form.IsFullScreen() {
		return
	}
	if w.framelessWithDecorations {
		win32.ExtendFrameIntoClientArea(w.Handle(), false)
	}
	w.Form.SetMaxSize(0, 0)
	w.Form.SetMinSize(0, 0)
	w.Form.Fullscreen()
}

func (w *Window) UnFullscreen() {
	if !w.Form.IsFullScreen() {
		return
	}
	if w.framelessWithDecorations {
		win32.ExtendFrameIntoClientArea(w.Handle(), true)
	}
	w.Form.UnFullscreen()
	w.SetMinSize(w.minWidth, w.minHeight)
	w.SetMaxSize(w.maxWidth, w.maxHeight)
}

func (w *Window) Restore() {
	if w.Form.IsFullScreen() {
		w.UnFullscreen()
	} else {
		w.Form.Restore()
	}
}

func (w *Window) SetMinSize(minWidth int, minHeight int) {
	w.minWidth = minWidth
	w.minHeight = minHeight
	w.Form.SetMinSize(minWidth, minHeight)
}

func (w *Window) SetMaxSize(maxWidth int, maxHeight int) {
	w.maxWidth = maxWidth
	w.maxHeight = maxHeight
	w.Form.SetMaxSize(maxWidth, maxHeight)
}

func (w *Window) IsVisible() bool {
	return win32.IsVisible(w.Handle())
}

func (w *Window) WndProc(msg uint32, wparam, lparam uintptr) uintptr {

	switch msg {
	case win32.WM_POWERBROADCAST:
		switch wparam {
		case win32.PBT_APMSUSPEND:
			if w.OnSuspend != nil {
				w.OnSuspend()
			}
		case win32.PBT_APMRESUMEAUTOMATIC:
			if w.OnResume != nil {
				w.OnResume()
			}
		}
	case w32.WM_SETTINGCHANGE:
		settingChanged := w32.UTF16PtrToString((*uint16)(unsafe.Pointer(lparam)))
		if settingChanged == "ImmersiveColorSet" {
			w.themeChanged = true
			w.UpdateTheme()
		}
		return 0
	case w32.WM_NCLBUTTONDOWN:
		w32.SetFocus(w.Handle())
	case w32.WM_MOVE, w32.WM_MOVING:
		w.chromium.NotifyParentWindowPositionChanged()
	case w32.WM_ACTIVATE:
		// 如果!w.frontendOptions.Frameless 也就是说，如果w.frontendOptions.Frameless为false（非框架模式），则执行以下代码
		w.themeChanged = true
		if int(wparam) == w32.WA_INACTIVE {
			w.isActive = false
			w.UpdateTheme()
		} else {
			w.isActive = true
			w.UpdateTheme()
			//}
		}

	case 0x02E0: //w32.WM_DPICHANGED
		newWindowSize := (*w32.RECT)(unsafe.Pointer(lparam))
		w32.SetWindowPos(w.Handle(),
			uintptr(0),
			int(newWindowSize.Left),
			int(newWindowSize.Top),
			int(newWindowSize.Right-newWindowSize.Left),
			int(newWindowSize.Bottom-newWindowSize.Top),
			w32.SWP_NOZORDER|w32.SWP_NOACTIVATE)
	}

	if w.frontendOptions.Frameless {
		switch msg {
		case w32.WM_ACTIVATE:
// 如果我们想要一个无边框的窗口，但保留默认的框架装饰样式，则扩展DWM客户端区域。
// 此选项不受在WM_NCCALCSIZE消息中返回0的影响。
// 结果是隐藏了标题栏，但仍保留了默认的窗口框架样式。
// 参考：https://docs.microsoft.com/zh-cn/windows/win32/api/dwmapi/nf-dwmapi-dwmextendframeintoclientarea#remarks
			if w.framelessWithDecorations {
				win32.ExtendFrameIntoClientArea(w.Handle(), true)
			}
		case w32.WM_NCCALCSIZE:
// 禁用标准边框，允许客户区占据整个窗口大小。
// 参考：https://docs.microsoft.com/zh-cn/windows/win32/winmsg/wm-nccalcsize#remarks
// 这将隐藏标题栏，并由于未显示标准边框，也将禁用用户交互的窗口调整大小功能。但我们仍需要WS_THICKFRAME样式来支持前端进行窗口调整大小的操作。
			if wparam != 0 {
				rgrc := (*w32.RECT)(unsafe.Pointer(lparam))
				if w.Form.IsFullScreen() {
					// 在全屏模式下，我们无需进行任何调整
					w.chromium.SetPadding(edge.Rect{})
				} else if w.IsMaximised() {
// 如果窗口最大化，我们必须调整客户区以适应显示器的工作区。否则，
// 有些内容会超出显示器的可见部分。
// 确保使用提供的RECT来获取显示器，因为在多屏幕模式下最大化时，
// 使用MonitorFromWindow可能会返回错误的显示器。
// 参考：https://github.com/MicrosoftEdge/WebView2Feedback/issues/2549
					monitor := w32.MonitorFromRect(rgrc, w32.MONITOR_DEFAULTTONULL)

					var monitorInfo w32.MONITORINFO
					monitorInfo.CbSize = uint32(unsafe.Sizeof(monitorInfo))
					if monitor != 0 && w32.GetMonitorInfo(monitor, &monitorInfo) {
						*rgrc = monitorInfo.RcWork

						maxWidth := w.frontendOptions.MaxWidth
						maxHeight := w.frontendOptions.MaxHeight
						if maxWidth > 0 || maxHeight > 0 {
							var dpiX, dpiY uint
							w32.GetDPIForMonitor(monitor, w32.MDT_EFFECTIVE_DPI, &dpiX, &dpiY)

							maxWidth := int32(winc.ScaleWithDPI(maxWidth, dpiX))
							if maxWidth > 0 && rgrc.Right-rgrc.Left > maxWidth {
								rgrc.Right = rgrc.Left + maxWidth
							}

							maxHeight := int32(winc.ScaleWithDPI(maxHeight, dpiY))
							if maxHeight > 0 && rgrc.Bottom-rgrc.Top > maxHeight {
								rgrc.Bottom = rgrc.Top + maxHeight
							}
						}
					}
					w.chromium.SetPadding(edge.Rect{})
				} else {
// 这是为了解决无边框模式（frameless mode）下使用WindowDecorations时出现的窗口大小调整闪烁问题
// 参考：https://stackoverflow.com/a/6558508
// 原始解决方案建议减小底部1px，但这在某些Windows版本上似乎会导致由于DrawBackground也使用了这一减少而导致底部出现一条细白线。
// 增加底部尺寸同样可以规避闪烁问题，但我们会损失WebView内容的1px高度，因此我们选择在内容底部填充1px作为补偿。
					rgrc.Bottom += 1
					w.chromium.SetPadding(edge.Rect{Bottom: 1})
				}
				return 0
			}
		}
	}
	return w.Form.WndProc(msg, wparam, lparam)
}

func (w *Window) IsMaximised() bool {
	return win32.IsWindowMaximised(w.Handle())
}

func (w *Window) IsMinimised() bool {
	return win32.IsWindowMinimised(w.Handle())
}

func (w *Window) IsNormal() bool {
	return win32.IsWindowNormal(w.Handle())
}

func (w *Window) IsFullScreen() bool {
	return win32.IsWindowFullScreen(w.Handle())
}

func (w *Window) SetTheme(theme winoptions.Theme) {
	w.theme = theme
	w.themeChanged = true
	w.Invoke(func() {
		w.UpdateTheme()
	})
}

func invokeSync[T any](cba *Window, fn func() (T, error)) (res T, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	cba.Invoke(func() {
		res, err = fn()
		wg.Done()
	})
	wg.Wait()
	return res, err
}
