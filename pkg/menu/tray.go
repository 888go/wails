package menu

// TrayMenu 是选项
type TrayMenu struct {
	// Label 是我们希望在托盘区显示的文本
	Label string

// Image 表示我们希望显示的托盘图标的名称。
// 这些图标在构建期间从 <projectdir>/trayicons 目录读取，
// 其文件名将作为 ID 使用，但不包括扩展名
// 例如：<projectdir>/trayicons/main.png 在此处可以通过 "main" 引用
// 如果该图象不是文件名，则它将被视为 base64 图像数据
	Image string

	// MacTemplateImage 表示在Mac系统中，这个图片是一个模板图片
	MacTemplateImage bool

	// Text Colour
	RGBA string

	// Font
	FontSize int
	FontName string

	// Tooltip
	Tooltip string

// 当菜单被点击时的回调函数
// 点击事件回调 `json:"-"`

	// Disabled 使项目不可选择
	Disabled bool

	// Menu 是我们希望在托盘中使用的初始菜单
	Menu *Menu

	// OnOpen 当Menu被打开时调用
	OnOpen func()

	// OnClose 当Menu被关闭时调用
	OnClose func()
}
