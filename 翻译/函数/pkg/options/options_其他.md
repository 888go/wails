# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 
# **_package.md 文件备注:
# bm= 包名,更换新的包名称, 如: package gin //bm:gin类
#
# **_其他.md 文件备注:
# hs= 行首,跳转到行首进行重命名.文档内如果有多个相同的,会一起重命名.
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# cf= 重复,用于重命名多次,如: 一个文档内有2个"One(result interface{}) error"需要重命名.
#     但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"
# zz= 正则表达式,用于结构名称替换或者复杂替换
#     如待替换: type authPair struct { //zz:^type *authPair

[Normal     WindowStartState = 0]
hs=常量_正常

[Maximised  WindowStartState = 1]
hs=常量_最大化

[Minimised  WindowStartState = 2]
hs=常量_最小化

[Fullscreen WindowStartState = 3]
hs=常量_全屏

[Title             string]
hs=标题

[Width             int]
hs=宽度

[Height            int]
hs=高度

[DisableResize     bool]
hs=禁用调整大小

[Fullscreen        bool]
hs=全屏

[Frameless         bool]
hs=无边框

[MinWidth          int]
hs=最小宽度

[MinHeight         int]
hs=最小高度

[MaxWidth          int]
hs=最大宽度

[MaxHeight         int]
hs=最大高度

[StartHidden       bool]
hs=启动时隐藏窗口

[HideWindowOnClose bool]
hs=关闭时隐藏窗口

[AlwaysOnTop       bool]
hs=始终置顶

[BackgroundColour *RGBA]
hs=背景颜色

[Assets fs.FS]
hs=Assets弃用

[AssetsHandler http.Handler]
hs=AssetsHandler弃用

[Menu               *menu.Menu]
hs=菜单

[Logger             logger.Logger   `json:"-"`]
hs=日志记录器

[LogLevel           logger.LogLevel]
hs=日志级别

[OnStartup          func(ctx context.Context)                `json:"-"`]
hs=绑定启动前函数

[OnDomReady         func(ctx context.Context)                `json:"-"`]
hs=绑定DOM就绪函数

[OnShutdown         func(ctx context.Context)                `json:"-"`]
hs=绑定应用退出函数

[WindowStartState   WindowStartState]
hs=窗口启动状态

[CSSDragProperty string]
hs=CSS拖动属性

[EnableDefaultContextMenu bool]
hs=右键菜单

[EnableFraudulentWebsiteDetection bool]
hs=启用欺诈网站检测

[SingleInstanceLock *SingleInstanceLock]
hs=单实例锁

[Windows *windows.Options]
hs=Windows选项

[Mac     *mac.Options]
hs=Mac选项

[Linux   *linux.Options]
hs=Linux选项

[Debug Debug]
hs=调试选项

[OnBeforeClose      func(ctx context.Context) (prevent bool) `json:"-"`]
hs=绑定应用关闭前函数

[CSSDragValue string]
hs=CSS拖动值

[LogLevelProduction logger.LogLevel]
hs=生产日志级别

[AssetServer        *assetserver.Options]
hs=绑定http请求

[Logger             logger.Logger                            `json:"-"`]
hs=日志

[Bind               #左中括号##右中括号#interface{}]
hs=绑定调用方法

[EnumBind           #左中括号##右中括号#interface{}]
hs=绑定常量枚举

[ErrorFormatter ErrorFormatter]
hs=错误格式化

[Experimental *Experimental]
hs=Experimental实验性
