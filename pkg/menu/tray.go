package menu

// TrayMenu 是选项
type TrayMenu struct {
	// Label 是我们希望在托盘区显示的文本
	X显示名称 string //hs:显示名称     

// Image 表示我们希望显示的托盘图标的名称。
// 这些图标在构建期间从 <projectdir>/trayicons 目录读取，
// 其文件名将作为 ID 使用，但不包括扩展名
// 例如：<projectdir>/trayicons/main.png 在此处可以通过 "main" 引用
// 如果该图象不是文件名，则它将被视为 base64 图像数据
	X图标名称 string //hs:图标名称     

	// MacTemplateImage 表示在Mac系统中，这个图片是一个模板图片
	Mac模板图标 bool //hs:Mac模板图标     

	// Text Colour
	RGBA string

	// Font
	X字体大小 int //hs:字体大小     
	X字体名称 string //hs:字体名称     

	// Tooltip
	X提示 string //hs:提示     

// 当菜单被点击时的回调函数
// 点击事件回调 `json:"-"`

	// Disabled 使项目不可选择
	X是否禁用 bool //hs:是否禁用     

	// Menu 是我们希望在托盘中使用的初始菜单
	X菜单 *Menu //hs:菜单     

	// OnOpen 当Menu被打开时调用
	X打开回调函数 func() //hs:打开回调函数     

	// OnClose 当Menu被关闭时调用
	X关闭回调函数 func() //hs:关闭回调函数     
}
