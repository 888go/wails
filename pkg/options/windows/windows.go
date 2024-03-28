package windows

type Theme int

type Messages struct {
	InstallationRequired string
	UpdateRequired       string
	MissingRequirements  string
	Webview2NotInstalled string
	Error                string
	FailedToInstall      string
	DownloadPage         string
	PressOKToInstall     string
	ContactAdmin         string
	InvalidFixedWebview2 string
	WebView2ProcessCrash string
}

const (
	// SystemDefault 将使用系统当前的主题。应用程序会跟随系统主题的变化。
	SystemDefault Theme = 0
	// Dark Mode
	Dark Theme = 1
	// Light Mode
	Light Theme = 2
)

type BackdropType int32

const (
	Auto    BackdropType = 0
	None    BackdropType = 1
	Mica    BackdropType = 2
	Acrylic BackdropType = 3
	Tabbed  BackdropType = 4
)

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
	WebviewIsTransparent bool
	WindowIsTranslucent  bool
	DisableWindowIcon    bool

	IsZoomControlEnabled bool
	ZoomFactor           float64

	DisablePinchZoom bool

// 在无边框模式下禁用所有窗口装饰，这意味着不会显示“Aero Shadow”和“圆角”。
// “圆角”仅在Windows 11系统上可用。
	DisableFramelessWindowDecorations bool

// WebView2 存储用户数据的路径。如果为空，则使用 %APPDATA%\[BinaryName.exe]。
// 如果路径无效，将显示一个包含错误信息的消息框，并且应用将以错误代码退出。
	WebviewUserDataPath string

	// WebView2可执行文件目录路径。如果为空，则使用系统已安装的WebView2。
	WebviewBrowserPath string

	// 深色/浅色或系统默认主题
	Theme Theme

	// 自定义暗黑模式和明亮模式的设置
	CustomTheme *ThemeSettings

	// 选择半透明背景类型。需要Windows 11 22621或更高版本。
	BackdropType BackdropType

	// 可自定义的用户消息
	Messages *Messages

// ResizeDebounceMS 是在调整窗口大小时，对 webview2 重绘操作进行防抖动的延时时间（单位：毫秒）
	ResizeDebounceMS uint16

	// OnSuspend 在Windows进入低功耗模式时被调用
	OnSuspend func()

	// OnResume 当Windows从低功耗模式恢复时被调用
	OnResume func()

	// WebviewGpuIsDisabled 用于启用/禁用 webview 的 GPU 加速功能
	WebviewGpuIsDisabled bool

// WebviewDisableRendererCodeIntegrity 禁用 WebView2 的 `RendererCodeIntegrity`。某些安全端点防护软件
// 会使用未签名或签名错误的 dll 注入到 WebView2 中，这是不允许的，并且会导致 WebView2 进程停止运行。
// 这类安全软件需要更新以解决此问题，或者可以通过设置此标志禁用完整性检查来暂时解决。
//
// Windows 事件查看器日志中包含如在 https://github.com/MicrosoftEdge/WebView2Feedback/issues/2051 中提及的 `代码完整性错误`。
//
// !! 请注意，禁用此功能时，也会允许恶意软件注入 WebView2，请谨慎操作 !!
	WebviewDisableRendererCodeIntegrity bool

	// 配置是否启用滑动手势
	EnableSwipeGestures bool
}

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
