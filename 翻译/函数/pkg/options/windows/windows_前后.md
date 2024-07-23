# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[InstallationRequired string]
qm=WebView2需安装
cz=InstallationRequired string

[UpdateRequired string]
qm=WebView2需更新
cz=UpdateRequired string

[MissingRequirements string]
qm=缺少必要组件
cz=MissingRequirements string

[Webview2NotInstalled string]
qm=WebView2未安装
cz=Webview2NotInstalled string

[Error string]
qm=出错
cz=Error string

[FailedToInstall string]
qm=安装失败
cz=FailedToInstall string

[DownloadPage string]
qm=跳转WebView2下载页面
cz=DownloadPage string

[PressOKToInstall string]
qm=按OK安装
cz=PressOKToInstall string

[ContactAdmin string]
qm=联系管理员
cz=ContactAdmin string

[InvalidFixedWebview2 string]
qm=WebView2指定目录无效
cz=InvalidFixedWebview2 string

[WebView2ProcessCrash string]
qm=WebView2进程崩溃
cz=WebView2ProcessCrash string

[SystemDefault Theme = 0]
qm=常量_win主题_默认
cz=SystemDefault Theme #等号# 0

[Dark Theme = 1]
qm=常量_win主题_暗黑
cz=Dark Theme #等号# 1

[Light Theme = 2]
qm=常量_win主题_浅色
cz=Light Theme #等号# 2

[Auto BackdropType = 0]
qm=常量_半透明类型_自动
cz=Auto BackdropType #等号# 0

[None BackdropType = 1]
qm=常量_半透明类型_无
cz=None BackdropType #等号# 1

[Mica BackdropType = 2]
qm=常量_半透明类型_Mica
cz=Mica BackdropType #等号# 2

[Acrylic BackdropType = 3]
qm=常量_半透明类型_亚克力
cz=Acrylic BackdropType #等号# 3

[Tabbed BackdropType = 4]
qm=常量_半透明类型_Tabbed
cz=Tabbed BackdropType #等号# 4

[WebviewIsTransparent bool]
qm=开启Webview透明
cz=WebviewIsTransparent bool

[WindowIsTranslucent bool]
qm=开启窗口半透明
cz=WindowIsTranslucent bool

[DisableWindowIcon bool]
qm=禁用窗口图标
cz=DisableWindowIcon bool

[IsZoomControlEnabled bool]
qm=启用缩放控制
cz=IsZoomControlEnabled bool

[ZoomFactor float64]
qm=缩放比例
cz=ZoomFactor float64

[DisablePinchZoom bool]
qm=禁用缩放
cz=DisablePinchZoom bool

[DisableFramelessWindowDecorations bool]
qm=禁用无边框窗口装饰
cz=DisableFramelessWindowDecorations bool

[WebviewUserDataPath string]
qm=webview用户数据路径
cz=WebviewUserDataPath string

[WebviewBrowserPath string]
qm=webview浏览器路径
cz=WebviewBrowserPath string

[Theme Theme]
qm=主题
cz=Theme Theme

[CustomTheme *ThemeSettings]
qm=自定义主题
cz=CustomTheme *ThemeSettings

[BackdropType BackdropType]
qm=背景半透明类型
cz=BackdropType BackdropType

[Messages *Messages]
qm=用户消息
cz=Messages *Messages

[ResizeDebounceMS uint16]
qm=重置尺寸防抖间隔
cz=ResizeDebounceMS uint16

[OnSuspend func()]
qm=低功耗模式时回调函数
cz=OnSuspend func()

[OnResume func()]
qm=低功耗模式恢复时回调函数
cz=OnResume func()

[WebviewGpuIsDisabled bool]
qm=禁用GPU加速
cz=WebviewGpuIsDisabled bool

[WebviewDisableRendererCodeIntegrity bool]
qm=禁用RendererCodeIntegrity
cz=WebviewDisableRendererCodeIntegrity bool

[EnableSwipeGestures bool]
qm=启用滑动手势
cz=EnableSwipeGestures bool
