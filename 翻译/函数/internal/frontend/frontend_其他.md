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

[Pattern     string]
hs=扩展名列表

[DefaultDirectory           string]
hs=默认目录
cf=2

[DefaultFilename            string]
hs=默认文件名
cf=2

[Title                      string]
hs=标题
cf=2

[Filters                    #左中括号##右中括号#FileFilter]
hs=过滤器
cf=2

[ShowHiddenFiles            bool]
hs=显示隐藏文件
cf=2

[CanCreateDirectories       bool]
hs=是否可创建目录
cf=2

[ResolvesAliases            bool]
hs=是否解析别名

[TreatPackagesAsDirectories bool]
hs=是否将包视为目录
cf=2

[InfoDialog     DialogType = "info"]
hs=常量_对话框_信息

[WarningDialog  DialogType = "warning"]
hs=常量_对话框_警告

[ErrorDialog    DialogType = "error"]
hs=常量_对话框_错误

[QuestionDialog DialogType = "question"]
hs=常量_对话框_问题

[Width int `json:"width"`]
hs=Width弃用

[Height int `json:"height"`]
hs=Height弃用
cf=2

[Size ScreenSize `json:"size"`]
hs=大小

[Type          DialogType]
hs=对话框类型

[Title         string]
hs=标题

[Message       string]
hs=消息

[Buttons       #左中括号##右中括号#string]
hs=按钮s

[DefaultButton string]
hs=默认按钮

[CancelButton  string]
hs=取消按钮

[Icon          #左中括号##右中括号#byte]
hs=图标

[DisplayName string]
hs=显示名称

[Run(ctx context.Context) error]
hs=运行

[ExecJS(js string)]
hs=窗口执行JS

[Hide()]
hs=隐藏

[Show()]
hs=显示

[Quit()]
hs=退出

[OpenFileDialog(dialogOptions OpenDialogOptions) (string, error)]
hs=对话框选择文件

[OpenMultipleFilesDialog(dialogOptions OpenDialogOptions) (#左中括号##右中括号#string, error)]
hs=对话框多选文件

[OpenDirectoryDialog(dialogOptions OpenDialogOptions) (string, error)]
hs=对话框选择目录

[SaveFileDialog(dialogOptions SaveDialogOptions) (string, error)]
hs=对话框保存文件

[MessageDialog(dialogOptions MessageDialogOptions) (string, error)]
hs=对话框弹出消息

[WindowSetTitle(title string)]
hs=窗口设置标题

[WindowShow()]
hs=窗口显示

[WindowHide()]
hs=窗口隐藏

[WindowCenter()]
hs=窗口居中

[WindowToggleMaximise()]
hs=窗口最大化切换

[WindowMaximise()]
hs=窗口最大化

[WindowUnmaximise()]
hs=窗口取消最大化

[WindowMinimise()]
hs=窗口最小化

[WindowUnminimise()]
hs=窗口取消最小化

[WindowSetAlwaysOnTop(b bool)]
hs=窗口设置置顶

[WindowSetPosition(x int, y int)]
hs=窗口设置位置

[WindowGetPosition() (int, int)]
hs=窗口取位置

[WindowSetSize(width int, height int)]
hs=窗口设置尺寸

[WindowGetSize() (int, int)]
hs=窗口取尺寸

[WindowSetMinSize(width int, height int)]
hs=窗口设置最小尺寸

[WindowSetMaxSize(width int, height int)]
hs=窗口设置最大尺寸

[WindowFullscreen()]
hs=窗口设置全屏

[WindowUnfullscreen()]
hs=窗口取消全屏

[WindowSetBackgroundColour(col *options.RGBA)]
hs=窗口设置背景色

[WindowReload()]
hs=窗口重载

[WindowReloadApp()]
hs=窗口重载应用程序前端

[WindowSetSystemDefaultTheme()]
hs=窗口设置系统默认主题

[WindowSetLightTheme()]
hs=窗口设置浅色主题

[WindowSetDarkTheme()]
hs=窗口设置深色主题

[WindowIsMaximised() bool]
hs=窗口是否最大化

[WindowIsMinimised() bool]
hs=窗口是否最小化

[WindowIsNormal() bool]
hs=窗口是否为正常

[WindowIsFullscreen() bool]
hs=窗口是否全屏

[WindowPrint()]
hs=窗口打开打印对话框

[ScreenGetAll() (#左中括号##右中括号#Screen, error)]
hs=取屏幕信息

[MenuSetApplicationMenu(menu *menu.Menu)]
hs=菜单设置

[MenuUpdateApplicationMenu()]
hs=菜单更新

[BrowserOpenURL(url string)]
hs=默认浏览器打开url

[ClipboardGetText() (string, error)]
hs=剪贴板取文本

[ClipboardSetText(text string) error]
hs=剪贴板置文本
