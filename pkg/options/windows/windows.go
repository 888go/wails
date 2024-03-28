package windows

type Theme int

type Messages struct {
	InstallationRequired string //hs:WebView2需安装     
	UpdateRequired       string //hs:WebView2需更新     
	MissingRequirements  string //hs:缺少必要组件     
	Webview2NotInstalled string //hs:WebView2未安装     
	Error                string //hs:出错     
	FailedToInstall      string //hs:安装失败     
	DownloadPage         string //hs:跳转WebView2下载页面     
	PressOKToInstall     string //hs:按OK安装     
	ContactAdmin         string //hs:联系管理员     
	InvalidFixedWebview2 string //hs:WebView2指定目录无效     
	WebView2ProcessCrash string //hs:WebView2进程崩溃     
}

const (
	// SystemDefault 将使用系统当前的主题。应用程序会跟随系统主题的变化。
	SystemDefault Theme = 0 //hs:常量_win主题_默认     
	// Dark Mode
	Dark Theme = 1 //hs:常量_win主题_暗黑     
	// Light Mode
	Light Theme = 2 //hs:常量_win主题_浅色     
)

type BackdropType int32

const (
	Auto    BackdropType = 0 //hs:常量_半透明类型_自动     
	None    BackdropType = 1 //hs:常量_半透明类型_无     
	Mica    BackdropType = 2 //hs:常量_半透明类型_Mica     
	Acrylic BackdropType = 3 //hs:常量_半透明类型_亚克力     
	Tabbed  BackdropType = 4 //hs:常量_半透明类型_Tabbed     
)


// ff:
// b:
// g:
// r:
func RGB(r, g, b uint8) int32 {
	col := int32(b)
	col = col<<8 | int32(g)
	col = col<<8 | int32(r)
	return col
}

// ThemeSettings 包含可选的颜色设置。
// 可以使用十六进制值 0x00BBGGRR 来设置这些颜色。
type ThemeSettings struct {
	DarkModeTitleBar           int32
	DarkModeTitleBarInactive   int32
	DarkModeTitleText          int32
	DarkModeTitleTextInactive  int32
	DarkModeBorder             int32
	DarkModeBorderInactive     int32
	LightModeTitleBar          int32
	LightModeTitleBarInactive  int32
	LightModeTitleText         int32
	LightModeTitleTextInactive int32
	LightModeBorder            int32
	LightModeBorderInactive    int32
}

// Options 是针对 Windows 的特定选项
type Options struct {
	WebviewIsTransparent bool //hs:开启Webview透明     
	WindowIsTranslucent  bool //hs:开启窗口半透明     
	DisableWindowIcon    bool //hs:禁用窗口图标     

	IsZoomControlEnabled bool //hs:启用缩放控制     
	ZoomFactor           float64 //hs:缩放比例     

	DisablePinchZoom bool //hs:禁用缩放     

// 在无边框模式下禁用所有窗口装饰，这意味着不会显示“Aero Shadow”和“圆角”。
// “圆角”仅在Windows 11系统上可用。
	DisableFramelessWindowDecorations bool //hs:禁用无边框窗口装饰     

// WebView2 存储用户数据的路径。如果为空，则使用 %APPDATA%\[BinaryName.exe]。
// 如果路径无效，将显示一个包含错误信息的消息框，并且应用将以错误代码退出。
	WebviewUserDataPath string //hs:webview用户数据路径     

	// WebView2可执行文件目录路径。如果为空，则使用系统已安装的WebView2。
	WebviewBrowserPath string //hs:webview浏览器路径     

	// 深色/浅色或系统默认主题
	Theme Theme //hs:主题     

	// 自定义暗黑模式和明亮模式的设置
	CustomTheme *ThemeSettings //hs:自定义主题     

	// 选择半透明背景类型。需要Windows 11 22621或更高版本。
	BackdropType BackdropType //hs:背景半透明类型     

	// 可自定义的用户消息
	Messages *Messages //hs:用户消息     

// ResizeDebounceMS 是在调整窗口大小时，对 webview2 重绘操作进行防抖动的延时时间（单位：毫秒）
	ResizeDebounceMS uint16 //hs:重置尺寸防抖间隔     

	// OnSuspend 在Windows进入低功耗模式时被调用
	OnSuspend func() //hs:低功耗模式时回调函数     

	// OnResume 当Windows从低功耗模式恢复时被调用
	OnResume func() //hs:低功耗模式恢复时回调函数     

	// WebviewGpuIsDisabled 用于启用/禁用 webview 的 GPU 加速功能
	WebviewGpuIsDisabled bool //hs:禁用GPU加速     

// WebviewDisableRendererCodeIntegrity 禁用 WebView2 的 `RendererCodeIntegrity`。某些安全端点防护软件
// 会使用未签名或签名错误的 dll 注入到 WebView2 中，这是不允许的，并且会导致 WebView2 进程停止运行。
// 这类安全软件需要更新以解决此问题，或者可以通过设置此标志禁用完整性检查来暂时解决。
//
// Windows 事件查看器日志中包含如在 https://github.com/MicrosoftEdge/WebView2Feedback/issues/2051 中提及的 `代码完整性错误`。
//
// !! 请注意，禁用此功能时，也会允许恶意软件注入 WebView2，请谨慎操作 !!
	WebviewDisableRendererCodeIntegrity bool //hs:禁用RendererCodeIntegrity     

	// 配置是否启用滑动手势
	EnableSwipeGestures bool //hs:启用滑动手势     
}


// ff:运行时默认提示
func DefaultMessages() *Messages {
	return &Messages{
		InstallationRequired: "The WebView2 runtime is required. Press Ok to download and install. Note: The installer will download silently so please wait.",
		UpdateRequired:       "The WebView2 runtime needs updating. Press Ok to download and install. Note: The installer will download silently so please wait.",
		MissingRequirements:  "Missing Requirements",
		Webview2NotInstalled: "WebView2 runtime not installed",
		Error:                "Error",
		FailedToInstall:      "The runtime failed to install correctly. Please try again.",
		DownloadPage:         "This application requires the WebView2 runtime. Press OK to open the download page. Minimum version required: ",
		PressOKToInstall:     "Press Ok to install.",
		ContactAdmin:         "The WebView2 runtime is required to run this application. Please contact your system administrator.",
		InvalidFixedWebview2: "The WebView2 runtime is manually specified, but It is not valid. Check minimum required version and webview2 path.",
		WebView2ProcessCrash: "The WebView2 process crashed and the application needs to be restarted.",
	}
}
