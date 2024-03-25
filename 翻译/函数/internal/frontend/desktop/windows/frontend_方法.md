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

[func (f *Frontend) WindowSetSystemDefaultTheme() {]
ff=窗口设置系统默认主题

[func (f *Frontend) WindowSetLightTheme() {]
ff=窗口设置浅色主题

[func (f *Frontend) WindowSetDarkTheme() {]
ff=窗口设置深色主题

[func (f *Frontend) Run(ctx context.Context) error {]
ff=运行
ctx=上下文

[func (f *Frontend) WindowSetSize(width, height int) {]
ff=窗口设置尺寸
height=高
width=宽

[func (f *Frontend) WindowGetSize() (int, int) {]
ff=窗口取尺寸

[func (f *Frontend) WindowSetTitle(title string) {]
ff=窗口设置标题
title=标题

[func (f *Frontend) WindowFullscreen() {]
ff=窗口设置全屏

[func (f *Frontend) WindowReloadApp() {]
ff=窗口重载应用程序前端

[func (f *Frontend) WindowUnfullscreen() {]
ff=窗口取消全屏

[func (f *Frontend) WindowShow() {]
ff=窗口显示

[func (f *Frontend) WindowHide() {]
ff=窗口隐藏

[func (f *Frontend) WindowMaximise() {]
ff=窗口最大化

[func (f *Frontend) WindowToggleMaximise() {]
ff=窗口最大化切换

[func (f *Frontend) WindowUnmaximise() {]
ff=窗口取消最大化

[func (f *Frontend) WindowMinimise() {]
ff=窗口最小化

[func (f *Frontend) WindowUnminimise() {]
ff=窗口取消最小化

[func (f *Frontend) WindowSetMinSize(width int, height int) {]
ff=窗口设置最小尺寸
height=高
width=宽

[func (f *Frontend) WindowSetMaxSize(width int, height int) {]
ff=窗口设置最大尺寸
height=高
width=宽

[func (f *Frontend) WindowSetBackgroundColour(col *options.RGBA) {]
ff=窗口设置背景色
col=颜色

[func (f *Frontend) ScreenGetAll() (#左中括号##右中括号#Screen, error) {]
ff=取屏幕信息

[func (f *Frontend) Show() {]
ff=显示

[func (f *Frontend) Hide() {]
ff=隐藏

[func (f *Frontend) WindowIsMaximised() bool {]
ff=窗口是否最大化

[func (f *Frontend) WindowIsMinimised() bool {]
ff=窗口是否最小化

[func (f *Frontend) WindowIsNormal() bool {]
ff=窗口是否为正常

[func (f *Frontend) WindowIsFullscreen() bool {]
ff=窗口是否全屏

[func (f *Frontend) Quit() {]
ff=退出

[func (f *Frontend) WindowPrint() {]
ff=窗口打开打印对话框

[func (f *Frontend) ExecJS(js string) {]
ff=窗口执行JS
js=js代码

[func (f *Frontend) WindowReload() {]
ff=窗口重载

[func (f *Frontend) WindowCenter() {]
ff=窗口居中

[func (f *Frontend) WindowSetAlwaysOnTop(b bool) {]
ff=窗口设置置顶
b=置顶

[func (f *Frontend) WindowSetPosition(x, y int) {]
ff=窗口设置位置

[func (f *Frontend) WindowGetPosition() (int, int) {]
ff=窗口取位置
