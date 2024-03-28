package windows

type Theme int

type Messages struct {
	WebView2需安装 string //hs:WebView2需安装     
	WebView2需更新       string //hs:WebView2需更新     
	X缺少必要组件  string //hs:缺少必要组件     
	WebView2未安装 string //hs:WebView2未安装     
	X出错                string //hs:出错     
	X安装失败      string //hs:安装失败     
	X跳转WebView2下载页面         string //hs:跳转WebView2下载页面     
	X按OK安装     string //hs:按OK安装     
	X联系管理员         string //hs:联系管理员     
	WebView2指定目录无效 string //hs:WebView2指定目录无效     
	WebView2进程崩溃 string //hs:WebView2进程崩溃     
}

const (
	// SystemDefault 将使用系统当前的主题。应用程序会跟随系统主题的变化。
	X常量_win主题_默认 Theme = 0 //hs:常量_win主题_默认     
	// Dark Mode
	X常量_win主题_暗黑 Theme = 1 //hs:常量_win主题_暗黑     
	// Light Mode
	X常量_win主题_浅色 Theme = 2 //hs:常量_win主题_浅色     
)

type BackdropType int32

const (
	X常量_半透明类型_自动    BackdropType = 0 //hs:常量_半透明类型_自动     
	X常量_半透明类型_无    BackdropType = 1 //hs:常量_半透明类型_无     
	X常量_半透明类型_Mica    BackdropType = 2 //hs:常量_半透明类型_Mica     
	X常量_半透明类型_亚克力 BackdropType = 3 //hs:常量_半透明类型_亚克力     
	X常量_半透明类型_Tabbed  BackdropType = 4 //hs:常量_半透明类型_Tabbed     
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
	X开启Webview透明 bool //hs:开启Webview透明     
	X开启窗口半透明  bool //hs:开启窗口半透明     
	X禁用窗口图标    bool //hs:禁用窗口图标     

	X启用缩放控制 bool //hs:启用缩放控制     
	X缩放比例           float64 //hs:缩放比例     

	X禁用缩放 bool //hs:禁用缩放     

// 在无边框模式下禁用所有窗口装饰，这意味着不会显示“Aero Shadow”和“圆角”。
// “圆角”仅在Windows 11系统上可用。
	X禁用无边框窗口装饰 bool //hs:禁用无边框窗口装饰     

// WebView2 存储用户数据的路径。如果为空，则使用 %APPDATA%\[BinaryName.exe]。
// 如果路径无效，将显示一个包含错误信息的消息框，并且应用将以错误代码退出。
	Webview用户数据路径 string //hs:webview用户数据路径     

	// WebView2可执行文件目录路径。如果为空，则使用系统已安装的WebView2。
	Webview浏览器路径 string //hs:webview浏览器路径     

	// 深色/浅色或系统默认主题
	X主题 Theme //hs:主题     

	// 自定义暗黑模式和明亮模式的设置
	X自定义主题 *ThemeSettings //hs:自定义主题     

	// 选择半透明背景类型。需要Windows 11 22621或更高版本。
	X背景半透明类型 BackdropType //hs:背景半透明类型     

	// 可自定义的用户消息
	X用户消息 *Messages //hs:用户消息     

// ResizeDebounceMS 是在调整窗口大小时，对 webview2 重绘操作进行防抖动的延时时间（单位：毫秒）
	X重置尺寸防抖间隔 uint16 //hs:重置尺寸防抖间隔     

	// OnSuspend 在Windows进入低功耗模式时被调用
	X低功耗模式时回调函数 func() //hs:低功耗模式时回调函数     

	// OnResume 当Windows从低功耗模式恢复时被调用
	X低功耗模式恢复时回调函数 func() //hs:低功耗模式恢复时回调函数     

	// WebviewGpuIsDisabled 用于启用/禁用 webview 的 GPU 加速功能
	X禁用GPU加速 bool //hs:禁用GPU加速     

// WebviewDisableRendererCodeIntegrity 禁用 WebView2 的 `RendererCodeIntegrity`。某些安全端点防护软件
// 会使用未签名或签名错误的 dll 注入到 WebView2 中，这是不允许的，并且会导致 WebView2 进程停止运行。
// 这类安全软件需要更新以解决此问题，或者可以通过设置此标志禁用完整性检查来暂时解决。
//
// Windows 事件查看器日志中包含如在 https://github.com/MicrosoftEdge/WebView2Feedback/issues/2051 中提及的 `代码完整性错误`。
//
// !! 请注意，禁用此功能时，也会允许恶意软件注入 WebView2，请谨慎操作 !!
	X禁用RendererCodeIntegrity bool //hs:禁用RendererCodeIntegrity     

	// 配置是否启用滑动手势
	X启用滑动手势 bool //hs:启用滑动手势     
}


// ff:运行时默认提示
func X运行时默认提示() *Messages {
	return &Messages{
		WebView2需安装: "The WebView2 runtime is required. Press Ok to download and install. Note: The installer will download silently so please wait.",
		WebView2需更新:       "The WebView2 runtime needs updating. Press Ok to download and install. Note: The installer will download silently so please wait.",
		X缺少必要组件:  "Missing Requirements",
		WebView2未安装: "WebView2 runtime not installed",
		X出错:                "Error",
		X安装失败:      "The runtime failed to install correctly. Please try again.",
		X跳转WebView2下载页面:         "This application requires the WebView2 runtime. Press OK to open the download page. Minimum version required: ",
		X按OK安装:     "Press Ok to install.",
		X联系管理员:         "The WebView2 runtime is required to run this application. Please contact your system administrator.",
		WebView2指定目录无效: "The WebView2 runtime is manually specified, but It is not valid. Check minimum required version and webview2 path.",
		WebView2进程崩溃: "The WebView2 process crashed and the application needs to be restarted.",
	}
}
