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

[func WindowSetTitle(ctx context.Context, title string) {]
ff=窗口设置标题
title=标题
ctx=上下文

[func WindowFullscreen(ctx context.Context) {]
ff=窗口设置全屏
ctx=上下文

[func WindowUnfullscreen(ctx context.Context) {]
ff=窗口取消全屏
ctx=上下文

[func WindowCenter(ctx context.Context) {]
ff=窗口居中
ctx=上下文

[func WindowReload(ctx context.Context) {]
ff=窗口重载
ctx=上下文

[func WindowReloadApp(ctx context.Context) {]
ff=窗口重载应用程序前端
ctx=上下文

[func WindowSetSystemDefaultTheme(ctx context.Context) {]
ff=窗口设置系统默认主题
ctx=上下文

[func WindowSetLightTheme(ctx context.Context) {]
ff=窗口设置浅色主题
ctx=上下文

[func WindowSetDarkTheme(ctx context.Context) {]
ff=窗口设置深色主题
ctx=上下文

[func WindowShow(ctx context.Context) {]
ff=窗口显示
ctx=上下文

[func WindowHide(ctx context.Context) {]
ff=窗口隐藏
ctx=上下文

[func WindowSetSize(ctx context.Context, width int, height int) {]
ff=窗口设置尺寸
height=高
width=宽
ctx=上下文

[func WindowGetSize(ctx context.Context) (int, int) {]
ff=窗口取尺寸
ctx=上下文

[func WindowSetMinSize(ctx context.Context, width int, height int) {]
ff=窗口设置最小尺寸
height=高
width=宽
ctx=上下文

[func WindowSetMaxSize(ctx context.Context, width int, height int) {]
ff=窗口设置最大尺寸
height=高
width=宽
ctx=上下文

[func WindowSetAlwaysOnTop(ctx context.Context, b bool) {]
ff=窗口设置置顶
b=置顶
ctx=上下文

[func WindowSetPosition(ctx context.Context, x int, y int) {]
ff=窗口设置位置
ctx=上下文

[func WindowGetPosition(ctx context.Context) (int, int) {]
ff=窗口取位置
ctx=上下文

[func WindowMaximise(ctx context.Context) {]
ff=窗口最大化
ctx=上下文

[func WindowToggleMaximise(ctx context.Context) {]
ff=窗口最大化切换
ctx=上下文

[func WindowUnmaximise(ctx context.Context) {]
ff=窗口取消最大化
ctx=上下文

[func WindowMinimise(ctx context.Context) {]
ff=窗口最小化
ctx=上下文

[func WindowUnminimise(ctx context.Context) {]
ff=窗口取消最小化
ctx=上下文

[func WindowIsFullscreen(ctx context.Context) bool {]
ff=窗口是否全屏
ctx=上下文

[func WindowIsMaximised(ctx context.Context) bool {]
ff=窗口是否最大化
ctx=上下文

[func WindowIsMinimised(ctx context.Context) bool {]
ff=窗口是否最小化
ctx=上下文

[func WindowIsNormal(ctx context.Context) bool {]
ff=窗口是否为正常
ctx=上下文

[func WindowExecJS(ctx context.Context, js string) {]
ff=窗口执行JS
js=js代码
ctx=上下文

[func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) {]
ff=窗口设置背景色
ctx=上下文

[func WindowPrint(ctx context.Context) {]
ff=窗口打开打印对话框
ctx=上下文
