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

[Normal WindowStartState = 0]
qm=常量_正常
cz=Normal WindowStartState #等号# 0

[Maximised WindowStartState = 1]
qm=常量_最大化
cz=Maximised WindowStartState #等号# 1

[Minimised WindowStartState = 2]
qm=常量_最小化
cz=Minimised WindowStartState #等号# 2

[Fullscreen WindowStartState = 3]
qm=常量_全屏
cz=Fullscreen WindowStartState #等号# 3

[Title string]
qm=标题
cz=Title string

[Width int]
qm=宽度
cz=Width int

[Height int]
qm=高度
cz=Height int

[DisableResize bool]
qm=禁用调整大小
cz=DisableResize bool

[Fullscreen bool]
qm=全屏
cz=Fullscreen bool

[Frameless bool]
qm=无边框
cz=Frameless bool

[MinWidth int]
qm=最小宽度
cz=MinWidth int

[MinHeight int]
qm=最小高度
cz=MinHeight int

[MaxWidth int]
qm=最大宽度
cz=MaxWidth int

[MaxHeight int]
qm=最大高度
cz=MaxHeight int

[StartHidden bool]
qm=启动时隐藏窗口
cz=StartHidden bool

[HideWindowOnClose bool]
qm=关闭时隐藏窗口
cz=HideWindowOnClose bool

[AlwaysOnTop bool]
qm=始终置顶
cz=AlwaysOnTop bool

[BackgroundColour *RGBA]
qm=背景颜色
cz=BackgroundColour *RGBA

[Assets fs.FS]
qm=Assets弃用
cz=Assets fs.FS

[AssetsHandler http.Handler]
qm=AssetsHandler弃用
cz=AssetsHandler http.Handler

[AssetServer *assetserver.Options]
qm=绑定http请求
cz=AssetServer *assetserver.Options

[Menu *menu.Menu]
qm=菜单
cz=Menu *menu.Menu

[Logger logger.Logger `json:"-"`]
qm=日志记录器
cz=Logger logger.Logger `json:"-"`

[LogLevel logger.LogLevel]
qm=日志级别
cz=LogLevel logger.LogLevel

[LogLevelProduction logger.LogLevel]
qm=生产日志级别
cz=LogLevelProduction logger.LogLevel

[OnStartup func(ctx context.Context) `json:"-"`]
qm=绑定启动前函数
cz=OnStartup func(ctx context.Context) `json:"-"`

[OnDomReady func(ctx context.Context) `json:"-"`]
qm=绑定DOM就绪函数
cz=OnDomReady func(ctx context.Context) `json:"-"`

[OnShutdown func(ctx context.Context) `json:"-"`]
qm=绑定应用退出函数
cz=OnShutdown func(ctx context.Context) `json:"-"`

[OnBeforeClose func(ctx context.Context) (prevent bool) `json:"-"`]
qm=绑定应用关闭前函数
cz=OnBeforeClose func(ctx context.Context) (prevent bool) `json:"-"`

[Bind #左中括号##右中括号#interface{}]
qm=绑定调用方法
cz=Bind []interface{}

[EnumBind #左中括号##右中括号#interface{}]
qm=绑定常量枚举
cz=EnumBind []interface{}

[WindowStartState WindowStartState]
qm=窗口启动状态
cz=WindowStartState WindowStartState

[ErrorFormatter ErrorFormatter]
qm=错误格式化
cz=ErrorFormatter ErrorFormatter

[CSSDragProperty string]
qm=CSS拖动属性
cz=CSSDragProperty string

[CSSDragValue string]
qm=CSS拖动值
cz=CSSDragValue string

[EnableDefaultContextMenu bool]
qm=右键菜单
cz=EnableDefaultContextMenu bool

[EnableFraudulentWebsiteDetection bool]
qm=启用欺诈网站检测
cz=EnableFraudulentWebsiteDetection bool

[SingleInstanceLock *SingleInstanceLock]
qm=单实例锁
cz=SingleInstanceLock *SingleInstanceLock

[Windows *windows.Options]
qm=Windows选项
cz=Windows *windows.Options

[Mac *mac.Options]
qm=Mac选项
cz=Mac *mac.Options

[Linux *linux.Options]
qm=Linux选项
cz=Linux *linux.Options

[Experimental *Experimental]
qm=Experimental实验性
cz=Experimental *Experimental

[Debug Debug]
qm=调试选项
cz=Debug Debug
