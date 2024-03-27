package mac

// TitleBar 包含 Mac 标题栏的选项
type TitleBar struct {
	TitlebarAppearsTransparent bool
	HideTitle                  bool
	HideTitleBar               bool
	FullSizeContent            bool
	UseToolbar                 bool
	HideToolbarSeparator       bool
}

// TitleBarDefault 会导致默认的Mac标题栏

// ff:
func TitleBarDefault() *TitleBar {
	return &TitleBar{
		TitlebarAppearsTransparent: false,
		HideTitle:                  false,
		HideTitleBar:               false,
		FullSizeContent:            false,
		UseToolbar:                 false,
		HideToolbarSeparator:       false,
	}
}

// 信用：注释来自Electron网站

// TitleBarHidden会显示一个隐藏的标题栏和一个全尺寸的内容窗口，
//然而标题栏仍然有标准的窗口控件(" traffic lights ")
//在左上角

// ff:
func TitleBarHidden() *TitleBar {
	return &TitleBar{
		TitlebarAppearsTransparent: true,
		HideTitle:                  true,
		HideTitleBar:               false,
		FullSizeContent:            true,
		UseToolbar:                 false,
		HideToolbarSeparator:       false,
	}
}

// TitleBarHiddenInset 会导致标题栏隐藏，并采用一种替代样式，其中交通灯按钮（关闭、最小化、最大化按钮）与窗口边缘的内嵌程度稍大一些。

// ff:
func TitleBarHiddenInset() *TitleBar {
	return &TitleBar{
		TitlebarAppearsTransparent: true,
		HideTitle:                  true,
		HideTitleBar:               false,
		FullSizeContent:            true,
		UseToolbar:                 true,
		HideToolbarSeparator:       true,
	}
}
