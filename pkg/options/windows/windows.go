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
	X常量_win主题_默认 Theme = 0
	// Dark Mode
	X常量_win主题_暗黑 Theme = 1
	// Light Mode
	X常量_win主题_浅色 Theme = 2
)

type BackdropType int32

const (
	X常量_半透明类型_自动    BackdropType = 0
	X常量_半透明类型_无    BackdropType = 1
	X常量_半透明类型_Mica    BackdropType = 2
	X常量_半透明类型_亚克力 BackdropType = 3
	X常量_半透明类型_Tabbed  BackdropType = 4
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
	X开启Webview透明 bool
	X开启窗口半透明  bool
	X禁用窗口图标    bool

	X启用缩放控制 bool
	X缩放比例           float64

	X禁用缩放 bool

	// 在无边框模式下禁用所有窗口装饰，这意味着不会显示“Aero Shadow”和“圆角”。
	// “圆角”仅在Windows 11系统上可用。
	X禁用无边框窗口装饰 bool

	// WebView2 存储用户数据的路径。如果为空，则使用 %APPDATA%\[BinaryName.exe]。
	// 如果路径无效，将显示一个包含错误信息的消息框，并且应用将以错误代码退出。
	Webview用户数据路径 string

	// WebView2可执行文件目录路径。如果为空，则使用系统已安装的WebView2。
	Webview浏览器路径 string

	// 深色/浅色或系统默认主题
	X主题 Theme

	// 自定义暗黑模式和明亮模式的设置
	X自定义主题 *ThemeSettings

	// 选择半透明背景类型。需要Windows 11 22621或更高版本。
	X背景半透明类型 BackdropType

	// 可自定义的用户消息
	X用户消息 *Messages

	// ResizeDebounceMS 是在调整窗口大小时，对 webview2 重绘操作进行防抖动的延时时间（单位：毫秒）
	X重置尺寸防抖间隔 uint16

	// OnSuspend 在Windows进入低功耗模式时被调用
	X低功耗模式时回调函数 func()

	// OnResume 当Windows从低功耗模式恢复时被调用
	X低功耗模式恢复时回调函数 func()

	// WebviewGpuIsDisabled 用于启用/禁用 webview 的 GPU 加速功能
	X禁用GPU加速 bool

	// WebviewDisableRendererCodeIntegrity 禁用 WebView2 的 `RendererCodeIntegrity`。某些安全端点防护软件
	// 会使用未签名或签名错误的 dll 注入到 WebView2 中，这是不允许的，并且会导致 WebView2 进程停止运行。
	// 这类安全软件需要更新以解决此问题，或者可以通过设置此标志禁用完整性检查来暂时解决。
	//
	// Windows 事件查看器日志中包含如在 https://github.com/MicrosoftEdge/WebView2Feedback/issues/2051 中提及的 `代码完整性错误`。
	//
	// !! 请注意，禁用此功能时，也会允许恶意软件注入 WebView2，请谨慎操作 !!
	X禁用RendererCodeIntegrity bool

	// 配置是否启用滑动手势
	X启用滑动手势 bool
}

func X运行时默认提示() *Messages {
	return &Messages{
		InstallationRequired: "WebView2运行时是必需的。按Ok下载安装。注意:安装程序会静默下载，请稍等。",
		UpdateRequired:       "WebView2运行时需要更新。按Ok下载安装。注意:安装程序会静默下载，请稍等。",
		MissingRequirements:  "失踪的需求",
		Webview2NotInstalled: "WebView2运行时未安装",
		Error:                "Error",
		FailedToInstall:      "运行时未能正确安装。请再试一次。",
		DownloadPage:         "此应用程序需要WebView2运行时。按OK打开下载页面。最低版本要求:",
		PressOKToInstall:     "按Ok进行安装。",
		ContactAdmin:         "运行此应用程序需要WebView2运行时。请联系您的系统管理员。",
		InvalidFixedWebview2: "WebView2运行时是手动指定的，但无效。检查最低要求版本和webview2路径。",
		WebView2ProcessCrash: "WebView2进程崩溃，应用程序需要重新启动。",
	}
}
