# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# qm= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[WindowIsTranslucent  bool]
qm=开启窗口半透明

[SystemDefault Theme = 0]
qm=常量_win主题_默认

[Dark Theme = 1]
qm=常量_win主题_暗黑

[Light Theme = 2]
qm=常量_win主题_浅色

[Auto    BackdropType = 0]
qm=常量_半透明类型_自动

[None    BackdropType = 1]
qm=常量_半透明类型_无

[Mica    BackdropType = 2]
qm=常量_半透明类型_Mica

[Acrylic BackdropType = 3]
qm=常量_半透明类型_亚克力

[Tabbed  BackdropType = 4]
qm=常量_半透明类型_Tabbed

[WebviewIsTransparent bool]
qm=开启Webview透明

[DisableWindowIcon    bool]
qm=禁用窗口图标

[IsZoomControlEnabled bool]
qm=启用缩放控制

[ZoomFactor           float64]
qm=缩放比例

[DisablePinchZoom bool]
qm=禁用缩放

[DisableFramelessWindowDecorations bool]
qm=禁用无边框窗口装饰

[WebviewUserDataPath string]
qm=webview用户数据路径

[WebviewBrowserPath string]
qm=webview浏览器路径

[Theme Theme]
qm=主题

[CustomTheme *ThemeSettings]
qm=自定义主题

[BackdropType BackdropType]
qm=背景半透明类型

[Messages *Messages]
qm=用户消息

[ResizeDebounceMS uint16]
qm=重置尺寸防抖间隔

[OnSuspend func()]
qm=低功耗模式时回调函数

[OnResume func()]
qm=低功耗模式恢复时回调函数

[WebviewGpuIsDisabled bool]
qm=禁用GPU加速

[WebviewDisableRendererCodeIntegrity bool]
qm=禁用RendererCodeIntegrity

[EnableSwipeGestures bool]
qm=启用滑动手势

[InstallationRequired string]
qm=WebView2需安装

[UpdateRequired       string]
qm=WebView2需更新

[MissingRequirements  string]
qm=缺少必要组件

[Webview2NotInstalled string]
qm=WebView2未安装

[Error                string]
qm=出错

[FailedToInstall      string]
qm=安装失败

[DownloadPage         string]
qm=跳转WebView2下载页面

[PressOKToInstall     string]
qm=按OK安装

[ContactAdmin         string]
qm=联系管理员

[InvalidFixedWebview2 string]
qm=WebView2指定目录无效

[WebView2ProcessCrash string]
qm=WebView2进程崩溃
