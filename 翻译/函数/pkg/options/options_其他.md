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

[Normal     WindowStartState = 0]
qm=常量_正常

[Maximised  WindowStartState = 1]
qm=常量_最大化

[Minimised  WindowStartState = 2]
qm=常量_最小化

[Fullscreen WindowStartState = 3]
qm=常量_全屏

[Title             string]
qm=标题

[Width             int]
qm=宽度

[Height            int]
qm=高度

[DisableResize     bool]
qm=禁用调整大小

[Fullscreen        bool]
qm=全屏

[Frameless         bool]
qm=无边框

[MinWidth          int]
qm=最小宽度

[MinHeight         int]
qm=最小高度

[MaxWidth          int]
qm=最大宽度

[MaxHeight         int]
qm=最大高度

[StartHidden       bool]
qm=启动时隐藏窗口

[HideWindowOnClose bool]
qm=关闭时隐藏窗口

[AlwaysOnTop       bool]
qm=始终置顶

[BackgroundColour *RGBA]
qm=背景颜色

[Assets fs.FS]
qm=Assets弃用

[AssetsHandler http.Handler]
qm=AssetsHandler弃用

[Menu               *menu.Menu]
qm=菜单

[Logger             logger.Logger   `json:"-"`]
qm=日志记录器

[LogLevel           logger.LogLevel]
qm=日志级别

[OnStartup          func(ctx context.Context)                `json:"-"`]
qm=绑定启动前函数

[OnDomReady         func(ctx context.Context)                `json:"-"`]
qm=绑定DOM就绪函数

[OnShutdown         func(ctx context.Context)                `json:"-"`]
qm=绑定应用退出函数

[WindowStartState   WindowStartState]
qm=窗口启动状态

[CSSDragProperty string]
qm=CSS拖动属性

[EnableDefaultContextMenu bool]
qm=右键菜单

[EnableFraudulentWebsiteDetection bool]
qm=启用欺诈网站检测

[SingleInstanceLock *SingleInstanceLock]
qm=单实例锁

[Windows *windows.Options]
qm=Windows选项

[Mac     *mac.Options]
qm=Mac选项

[Linux   *linux.Options]
qm=Linux选项

[Debug Debug]
qm=调试选项

[OnBeforeClose      func(ctx context.Context) (prevent bool) `json:"-"`]
qm=绑定应用关闭前函数

[CSSDragValue string]
qm=CSS拖动值

[LogLevelProduction logger.LogLevel]
qm=生产日志级别

[AssetServer        *assetserver.Options]
qm=绑定http请求

 
[Bind               #左中括号##右中括号#interface{}]
qm=绑定调用方法

[EnumBind           #左中括号##右中括号#interface{}]
qm=绑定常量枚举

[ErrorFormatter ErrorFormatter]
qm=错误格式化

[Experimental *Experimental]
qm=Experimental实验性
