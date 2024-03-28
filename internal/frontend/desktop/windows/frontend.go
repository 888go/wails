//go:build windows
// +build windows

package windows

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/bep/debounce"
	"github.com/wailsapp/go-webview2/pkg/edge"
	"github.com/wailsapp/wails/v2/internal/binding"
	"github.com/wailsapp/wails/v2/internal/frontend"
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/win32"
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc"
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
	wailsruntime "github.com/wailsapp/wails/v2/internal/frontend/runtime"
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/internal/system/operatingsystem"
	"github.com/wailsapp/wails/v2/pkg/assetserver"
	"github.com/wailsapp/wails/v2/pkg/assetserver/webview"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

const startURL = "http://wails.localhost/"

var secondInstanceBuffer = make(chan options.SecondInstanceData, 1)

type Screen = frontend.Screen

type Frontend struct {

	// Context
	ctx context.Context

	frontendOptions *options.App
	logger          *logger.Logger
	chromium        *edge.Chromium
	debug           bool
	devtoolsEnabled bool

	// Assets
	assets   *assetserver.AssetServer
	startURL *url.URL

	// main window handle
	mainWindow *Window
	bindings   *binding.Bindings
	dispatcher frontend.Dispatcher

	hasStarted bool

	// Windows build number
	versionInfo     *operatingsystem.WindowsVersionInfo
	resizeDebouncer func(f func())
}


// ff:
// dispatcher:
// appBindings:
// myLogger:
// appoptions:
// ctx:
func NewFrontend(ctx context.Context, appoptions *options.App, myLogger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) *Frontend {

	// 获取Windows构建号
	versionInfo, _ := operatingsystem.GetWindowsVersionInfo()

	result := &Frontend{
		frontendOptions: appoptions,
		logger:          myLogger,
		bindings:        appBindings,
		dispatcher:      dispatcher,
		ctx:             ctx,
		versionInfo:     versionInfo,
	}

	if appoptions.Windows != nil {
		if appoptions.Windows.ResizeDebounceMS > 0 {
			result.resizeDebouncer = debounce.New(time.Duration(appoptions.Windows.ResizeDebounceMS) * time.Millisecond)
		}
	}

	// 目前我们不能像其他平台那样使用 wails://wails/，因此我们将 assets 服务器映射到以下 URL。
	result.startURL, _ = url.Parse(startURL)

	if _starturl, _ := ctx.Value("starturl").(*url.URL); _starturl != nil {
		result.startURL = _starturl
		return result
	}

	if port, _ := ctx.Value("assetserverport").(string); port != "" {
		result.startURL.Host = net.JoinHostPort(result.startURL.Host, port)
	}

	var bindings string
	var err error
	if _obfuscated, _ := ctx.Value("obfuscated").(bool); !_obfuscated {
		bindings, err = appBindings.ToJSON()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		appBindings.DB().UpdateObfuscatedCallMap()
	}

	assets, err := assetserver.NewAssetServerMainPage(bindings, appoptions, ctx.Value("assetdir") != nil, myLogger, wailsruntime.RuntimeAssetsBundle)
	if err != nil {
		log.Fatal(err)
	}
	result.assets = assets

	go result.startSecondInstanceProcessor()

	return result
}


// ff:
func (f *Frontend) WindowReload() {
	f.ExecJS("runtime.WindowReload();")
}


// ff:
func (f *Frontend) WindowSetSystemDefaultTheme() {
	f.mainWindow.SetTheme(windows.SystemDefault)
}


// ff:
func (f *Frontend) WindowSetLightTheme() {
	f.mainWindow.SetTheme(windows.Light)
}


// ff:
func (f *Frontend) WindowSetDarkTheme() {
	f.mainWindow.SetTheme(windows.Dark)
}


// ff:
// ctx:
func (f *Frontend) Run(ctx context.Context) error {
	f.ctx = ctx

	f.chromium = edge.NewChromium()

	if f.frontendOptions.SingleInstanceLock != nil {
		SetupSingleInstance(f.frontendOptions.SingleInstanceLock.UniqueId)
	}

	mainWindow := NewWindow(nil, f.frontendOptions, f.versionInfo, f.chromium)
	f.mainWindow = mainWindow

	var _debug = ctx.Value("debug")
	var _devtoolsEnabled = ctx.Value("devtoolsEnabled")

	if _debug != nil {
		f.debug = _debug.(bool)
	}
	if _devtoolsEnabled != nil {
		f.devtoolsEnabled = _devtoolsEnabled.(bool)
	}

	f.WindowCenter()
	f.setupChromium()

	mainWindow.OnSize().Bind(func(arg *winc.Event) {
		if f.frontendOptions.Frameless {
// 如果窗口是无边框的并且我们正在进行最小化操作，那么我们需要抑制WebView2上的Resize事件。如果不这样做，在恢复窗口大小时无法按预期工作，并且在恢复动画期间首先会以错误的尺寸还原，直到动画完成后才会完全渲染。这高度依赖于WebView中的内容，详细信息参见 https://github.com/wailsapp/wails/issues/1319
			event, _ := arg.Data.(*winc.SizeEventData)
			if event != nil && event.Type == w32.SIZE_MINIMIZED {
				return
			}
		}

		if f.resizeDebouncer != nil {
			f.resizeDebouncer(func() {
				f.mainWindow.Invoke(func() {
					f.chromium.Resize()
				})
			})
		} else {
			f.chromium.Resize()
		}
	})

	mainWindow.OnClose().Bind(func(arg *winc.Event) {
		if f.frontendOptions.HideWindowOnClose {
			f.WindowHide()
		} else {
			f.Quit()
		}
	})

	go func() {
		if f.frontendOptions.OnStartup != nil {
			f.frontendOptions.OnStartup(f.ctx)
		}
	}()
	mainWindow.UpdateTheme()
	return nil
}


// ff:
func (f *Frontend) WindowClose() {
	if f.mainWindow != nil {
		f.mainWindow.Close()
	}
}


// ff:
func (f *Frontend) RunMainLoop() {
	_ = winc.RunMainLoop()
}


// ff:
func (f *Frontend) WindowCenter() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.Center()
}


// ff:
// b:
func (f *Frontend) WindowSetAlwaysOnTop(b bool) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetAlwaysOnTop(b)
}


// ff:
// y:
// x:
func (f *Frontend) WindowSetPosition(x, y int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetPos(x, y)
}

// ff:
func (f *Frontend) WindowGetPosition() (int, int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	return f.mainWindow.Pos()
}


// ff:
// height:
// width:
func (f *Frontend) WindowSetSize(width, height int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetSize(width, height)
}


// ff:
func (f *Frontend) WindowGetSize() (int, int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	return f.mainWindow.Size()
}


// ff:
// title:
func (f *Frontend) WindowSetTitle(title string) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetText(title)
}


// ff:
func (f *Frontend) WindowFullscreen() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = false;")
	}
	f.mainWindow.Fullscreen()
}


// ff:
func (f *Frontend) WindowReloadApp() {
	f.ExecJS(fmt.Sprintf("window.location.href = '%s';", f.startURL))
}


// ff:
func (f *Frontend) WindowUnfullscreen() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = true;")
	}
	f.mainWindow.UnFullscreen()
}


// ff:
func (f *Frontend) WindowShow() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.ShowWindow()
}


// ff:
func (f *Frontend) WindowHide() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.Hide()
}


// ff:
func (f *Frontend) WindowMaximise() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.hasStarted {
		if !f.frontendOptions.DisableResize {
			f.mainWindow.Maximise()
		}
	} else {
		f.frontendOptions.WindowStartState = options.Maximised
	}
}


// ff:
func (f *Frontend) WindowToggleMaximise() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if !f.hasStarted {
		return
	}
	if f.mainWindow.IsMaximised() {
		f.WindowUnmaximise()
	} else {
		f.WindowMaximise()
	}
}


// ff:
func (f *Frontend) WindowUnmaximise() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.mainWindow.Form.IsFullScreen() {
		return
	}
	f.mainWindow.Restore()
}


// ff:
func (f *Frontend) WindowMinimise() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.hasStarted {
		f.mainWindow.Minimise()
	} else {
		f.frontendOptions.WindowStartState = options.Minimised
	}
}


// ff:
func (f *Frontend) WindowUnminimise() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if f.mainWindow.Form.IsFullScreen() {
		return
	}
	f.mainWindow.Restore()
}


// ff:
// height:
// width:
func (f *Frontend) WindowSetMinSize(width int, height int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetMinSize(width, height)
}

// ff:
// height:
// width:
func (f *Frontend) WindowSetMaxSize(width int, height int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	f.mainWindow.SetMaxSize(width, height)
}


// ff:
// col:
func (f *Frontend) WindowSetBackgroundColour(col *options.RGBA) {
	if col == nil {
		return
	}

	f.mainWindow.Invoke(func() {
		win32.SetBackgroundColour(f.mainWindow.Handle(), col.R, col.G, col.B)

		controller := f.chromium.GetController()
		controller2 := controller.GetICoreWebView2Controller2()

		backgroundCol := edge.COREWEBVIEW2_COLOR{
			A: col.A,
			R: col.R,
			G: col.G,
			B: col.B,
		}

		// WebView2仅将0和255视为有效值。
		if backgroundCol.A > 0 && backgroundCol.A < 255 {
			backgroundCol.A = 255
		}

		if f.frontendOptions.Windows != nil && f.frontendOptions.Windows.WebviewIsTransparent {
			backgroundCol.A = 0
		}

		err := controller2.PutDefaultBackgroundColor(backgroundCol)
		if err != nil {
			log.Fatal(err)
		}
	})

}


// ff:
func (f *Frontend) ScreenGetAll() ([]Screen, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	screens := []Screen{}
	err := error(nil)
	f.mainWindow.Invoke(func() {
		screens, err = GetAllScreens(f.mainWindow.Handle())
		wg.Done()

	})
	wg.Wait()
	return screens, err
}


// ff:
func (f *Frontend) Show() {
	f.mainWindow.Show()
}


// ff:
func (f *Frontend) Hide() {
	f.mainWindow.Hide()
}


// ff:
func (f *Frontend) WindowIsMaximised() bool {
	return f.mainWindow.IsMaximised()
}


// ff:
func (f *Frontend) WindowIsMinimised() bool {
	return f.mainWindow.IsMinimised()
}


// ff:
func (f *Frontend) WindowIsNormal() bool {
	return f.mainWindow.IsNormal()
}


// ff:
func (f *Frontend) WindowIsFullscreen() bool {
	return f.mainWindow.IsFullScreen()
}


// ff:
func (f *Frontend) Quit() {
	if f.frontendOptions.OnBeforeClose != nil && f.frontendOptions.OnBeforeClose(f.ctx) {
		return
	}
// Exit 必须在主线程上调用。它会调用 PostQuitMessage，该函数向线程的消息队列发送 WM_QUIT 消息，
// 而我们的消息队列是在主线程上运行的。
	f.mainWindow.Invoke(winc.Exit)
}


// ff:
func (f *Frontend) WindowPrint() {
	f.ExecJS("window.print();")
}

func (f *Frontend) setupChromium() {
	chromium := f.chromium

	disableFeatues := []string{}
	if !f.frontendOptions.EnableFraudulentWebsiteDetection {
		disableFeatues = append(disableFeatues, "msSmartScreenProtection")
	}

	if opts := f.frontendOptions.Windows; opts != nil {
		chromium.DataPath = opts.WebviewUserDataPath
		chromium.BrowserPath = opts.WebviewBrowserPath

		if opts.WebviewGpuIsDisabled {
			chromium.AdditionalBrowserArgs = append(chromium.AdditionalBrowserArgs, "--disable-gpu")
		}
		if opts.WebviewDisableRendererCodeIntegrity {
			disableFeatues = append(disableFeatues, "RendererCodeIntegrity")
		}
	}

	if len(disableFeatues) > 0 {
		arg := fmt.Sprintf("--disable-features=%s", strings.Join(disableFeatues, ","))
		chromium.AdditionalBrowserArgs = append(chromium.AdditionalBrowserArgs, arg)
	}

	chromium.MessageCallback = f.processMessage
	chromium.WebResourceRequestedCallback = f.processRequest
	chromium.NavigationCompletedCallback = f.navigationCompleted
	chromium.AcceleratorKeyCallback = func(vkey uint) bool {
		if vkey == w32.VK_F12 && f.devtoolsEnabled {
			var keyState [256]byte
			if w32.GetKeyboardState(keyState[:]) {
				// 检查是否按下了CTRL键
				if keyState[w32.VK_CONTROL]&0x80 != 0 && keyState[w32.VK_SHIFT]&0x80 != 0 {
					chromium.OpenDevToolsWindow()
					return true
				}
			} else {
				f.logger.Error("Call to GetKeyboardState failed")
			}
		}
		w32.PostMessage(f.mainWindow.Handle(), w32.WM_KEYDOWN, uintptr(vkey), 0)
		return false
	}
	chromium.ProcessFailedCallback = func(sender *edge.ICoreWebView2, args *edge.ICoreWebView2ProcessFailedEventArgs) {
		kind, err := args.GetProcessFailedKind()
		if err != nil {
			f.logger.Error("GetProcessFailedKind: %s", err)
			return
		}

		f.logger.Error("WebVie2wProcess failed with kind %d", kind)
		switch kind {
		case edge.COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED:
			// => 为了从这个故障中恢复，应用必须重新创建一个新的WebView。
			messages := windows.DefaultMessages()
			if f.frontendOptions.Windows != nil && f.frontendOptions.Windows.Messages != nil {
				messages = f.frontendOptions.Windows.Messages
			}
			winc.Errorf(f.mainWindow, messages.WebView2ProcessCrash)
			os.Exit(-1)
		case edge.COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED,
			edge.COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED:
// => 自动创建一个新的渲染进程，并导航到错误页面。
// => 确保错误页面被展示出来。
			if !f.hasStarted {
				// NavgiationCompleted 事件未触发，确保 Chromium 正在显示
				chromium.Show()
			}
			if !f.mainWindow.hasBeenShown {
				// 窗口从未被显示过，确保将其显示出来
				f.ShowWindow()
			}
		}
	}

	chromium.Embed(f.mainWindow.Handle())

	if chromium.HasCapability(edge.SwipeNavigation) {
		swipeGesturesEnabled := f.frontendOptions.Windows != nil && f.frontendOptions.Windows.EnableSwipeGestures
		err := chromium.PutIsSwipeNavigationEnabled(swipeGesturesEnabled)
		if err != nil {
			log.Fatal(err)
		}
	}
	chromium.Resize()
	settings, err := chromium.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	err = settings.PutAreDefaultContextMenusEnabled(f.debug || f.frontendOptions.EnableDefaultContextMenu)
	if err != nil {
		log.Fatal(err)
	}
	err = settings.PutAreDevToolsEnabled(f.devtoolsEnabled)
	if err != nil {
		log.Fatal(err)
	}

	if opts := f.frontendOptions.Windows; opts != nil {
		if opts.ZoomFactor > 0.0 {
			chromium.PutZoomFactor(opts.ZoomFactor)
		}
		err = settings.PutIsZoomControlEnabled(opts.IsZoomControlEnabled)
		if err != nil {
			log.Fatal(err)
		}
		err = settings.PutIsPinchZoomEnabled(!opts.DisablePinchZoom)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = settings.PutIsStatusBarEnabled(false)
	if err != nil {
		log.Fatal(err)
	}
	err = settings.PutAreBrowserAcceleratorKeysEnabled(false)
	if err != nil {
		log.Fatal(err)
	}

	if f.debug && f.frontendOptions.Debug.OpenInspectorOnStartup {
		chromium.OpenDevToolsWindow()
	}

	// 设置焦点事件处理器
	onFocus := f.mainWindow.OnSetFocus()
	onFocus.Bind(f.onFocus)

	// Set background colour
	f.WindowSetBackgroundColour(f.frontendOptions.BackgroundColour)

	chromium.SetGlobalPermission(edge.CoreWebView2PermissionStateAllow)
	chromium.AddWebResourceRequestedFilter("*", edge.COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL)
	chromium.Navigate(f.startURL.String())
}

type EventNotify struct {
	Name string        `json:"name"`
	Data []interface{} `json:"data"`
}


// ff:
// data:
// name:
func (f *Frontend) Notify(name string, data ...interface{}) {
	notification := EventNotify{
		Name: name,
		Data: data,
	}
	payload, err := json.Marshal(notification)
	if err != nil {
		f.logger.Error(err.Error())
		return
	}
	f.ExecJS(`window.wails.EventsNotify('` + template.JSEscapeString(string(payload)) + `');`)
}

func (f *Frontend) processRequest(req *edge.ICoreWebView2WebResourceRequest, args *edge.ICoreWebView2WebResourceRequestedEventArgs) {
// 在CoreWebView2Settings上设置UserAgent会清空Edge浏览器的整个默认UserAgent，
// 但我们只想追加我们的ApplicationIdentifier。因此，我们对每个请求调整UserAgent。
	if reqHeaders, err := req.GetHeaders(); err == nil {
		useragent, _ := reqHeaders.GetHeader(assetserver.HeaderUserAgent)
		useragent = strings.Join([]string{useragent, assetserver.WailsUserAgentValue}, " ")
		reqHeaders.SetHeader(assetserver.HeaderUserAgent, useragent)
		reqHeaders.Release()
	}

	if f.assets == nil {
		// 我们使用devServer让WebView2通过其默认处理器来处理请求
		return
	}

	//Get the request
	uri, _ := req.GetUri()
	reqUri, err := url.ParseRequestURI(uri)
	if err != nil {
		f.logger.Error("Unable to parse equest uri %s: %s", uri, err)
		return
	}

	if reqUri.Scheme != f.startURL.Scheme {
		// 让WebView2使用其默认处理器处理请求
		return
	} else if reqUri.Host != f.startURL.Host {
		// 让WebView2使用其默认处理器处理请求
		return
	}

	webviewRequest, err := webview.NewRequest(
		f.chromium.Environment(),
		args,
		func(fn func()) {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			if f.mainWindow.InvokeRequired() {
				var wg sync.WaitGroup
				wg.Add(1)
				f.mainWindow.Invoke(func() {
					fn()
					wg.Done()
				})
				wg.Wait()
			} else {
				fn()
			}
		})

	if err != nil {
		f.logger.Error("%s: NewRequest failed: %s", uri, err)
		return
	}

	f.assets.ServeWebViewRequest(webviewRequest)
}

var edgeMap = map[string]uintptr{
	"n-resize":  w32.HTTOP,
	"ne-resize": w32.HTTOPRIGHT,
	"e-resize":  w32.HTRIGHT,
	"se-resize": w32.HTBOTTOMRIGHT,
	"s-resize":  w32.HTBOTTOM,
	"sw-resize": w32.HTBOTTOMLEFT,
	"w-resize":  w32.HTLEFT,
	"nw-resize": w32.HTTOPLEFT,
}

func (f *Frontend) processMessage(message string) {
	if message == "drag" {
		if !f.mainWindow.IsFullScreen() {
			err := f.startDrag()
			if err != nil {
				f.logger.Error(err.Error())
			}
		}
		return
	}

	if message == "runtime:ready" {
		cmd := fmt.Sprintf("window.wails.setCSSDragProperties('%s', '%s');", f.frontendOptions.CSSDragProperty, f.frontendOptions.CSSDragValue)
		f.ExecJS(cmd)
		return
	}

	if strings.HasPrefix(message, "resize:") {
		if !f.mainWindow.IsFullScreen() {
			sl := strings.Split(message, ":")
			if len(sl) != 2 {
				f.logger.Info("Unknown message returned from dispatcher: %+v", message)
				return
			}
			edge := edgeMap[sl[1]]
			err := f.startResize(edge)
			if err != nil {
				f.logger.Error(err.Error())
			}
		}
		return
	}

	go func() {
		result, err := f.dispatcher.ProcessMessage(message, f)
		if err != nil {
			f.logger.Error(err.Error())
			f.Callback(result)
			return
		}
		if result == "" {
			return
		}

		switch result[0] {
		case 'c':
			// 从方法调用返回的回调函数
			f.Callback(result[1:])
		default:
			f.logger.Info("Unknown message returned from dispatcher: %+v", result)
		}
	}()
}


// ff:
// message:
func (f *Frontend) Callback(message string) {
	escaped, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	f.mainWindow.Invoke(func() {
		f.chromium.Eval(`window.wails.Callback(` + string(escaped) + `);`)
	})
}

func (f *Frontend) startDrag() error {
	if !w32.ReleaseCapture() {
		return fmt.Errorf("unable to release mouse capture")
	}
	// 使用PostMessage是因为我们不希望在拖拽操作完成之前阻塞调用者。
	w32.PostMessage(f.mainWindow.Handle(), w32.WM_NCLBUTTONDOWN, w32.HTCAPTION, 0)
	return nil
}

func (f *Frontend) startResize(border uintptr) error {
	if !w32.ReleaseCapture() {
		return fmt.Errorf("unable to release mouse capture")
	}
	// 使用PostMessage是因为我们不希望在调整大小完成之前阻塞调用者。
	w32.PostMessage(f.mainWindow.Handle(), w32.WM_NCLBUTTONDOWN, border, 0)
	return nil
}


// ff:
// js:
func (f *Frontend) ExecJS(js string) {
	f.mainWindow.Invoke(func() {
		f.chromium.Eval(js)
	})
}

func (f *Frontend) navigationCompleted(sender *edge.ICoreWebView2, args *edge.ICoreWebView2NavigationCompletedEventArgs) {
	if f.frontendOptions.OnDomReady != nil {
		go f.frontendOptions.OnDomReady(f.ctx)
	}

	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = true;")
	}

	if f.hasStarted {
		return
	}
	f.hasStarted = true

	// 临时解决方案以使其可见：https://github.com/MicrosoftEdge/WebView2Feedback/issues/1077#issuecomment-825375026
// （该段英文注释描述了一个临时性的解决方案，用于解决某个特定问题以达到使其可见的目的。具体问题和方案请参考链接中的GitHub讨论，该讨论位于MicrosoftEdge/WebView2Feedback仓库的第1077号Issue中的一条编号为825375026的评论。）
	err := f.chromium.Hide()
	if err != nil {
		log.Fatal(err)
	}
	err = f.chromium.Show()
	if err != nil {
		log.Fatal(err)
	}

	if f.frontendOptions.StartHidden {
		return
	}

	switch f.frontendOptions.WindowStartState {
	case options.Maximised:
		if !f.frontendOptions.DisableResize {
			win32.ShowWindowMaximised(f.mainWindow.Handle())
		} else {
			win32.ShowWindow(f.mainWindow.Handle())
		}
	case options.Minimised:
		win32.ShowWindowMinimised(f.mainWindow.Handle())
	case options.Fullscreen:
		f.mainWindow.Fullscreen()
		win32.ShowWindow(f.mainWindow.Handle())
	default:
		if f.frontendOptions.Fullscreen {
			f.mainWindow.Fullscreen()
		}
		win32.ShowWindow(f.mainWindow.Handle())
	}

	f.mainWindow.hasBeenShown = true

}


// ff:
func (f *Frontend) ShowWindow() {
	f.mainWindow.Invoke(func() {
		if !f.mainWindow.hasBeenShown {
			f.mainWindow.hasBeenShown = true
			switch f.frontendOptions.WindowStartState {
			case options.Maximised:
				if !f.frontendOptions.DisableResize {
					win32.ShowWindowMaximised(f.mainWindow.Handle())
				} else {
					win32.ShowWindow(f.mainWindow.Handle())
				}
			case options.Minimised:
				win32.RestoreWindow(f.mainWindow.Handle())
			case options.Fullscreen:
				f.mainWindow.Fullscreen()
				win32.ShowWindow(f.mainWindow.Handle())
			default:
				if f.frontendOptions.Fullscreen {
					f.mainWindow.Fullscreen()
				}
				win32.ShowWindow(f.mainWindow.Handle())
			}
		} else {
			if win32.IsWindowMinimised(f.mainWindow.Handle()) {
				win32.RestoreWindow(f.mainWindow.Handle())
			} else {
				win32.ShowWindow(f.mainWindow.Handle())
			}
		}
		w32.SetForegroundWindow(f.mainWindow.Handle())
		w32.SetFocus(f.mainWindow.Handle())
	})

}

func (f *Frontend) onFocus(arg *winc.Event) {
	f.chromium.Focus()
}

func (f *Frontend) startSecondInstanceProcessor() {
	for secondInstanceData := range secondInstanceBuffer {
		if f.frontendOptions.SingleInstanceLock != nil &&
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch != nil {
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch(secondInstanceData)
		}
	}
}
