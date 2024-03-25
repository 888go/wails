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

// TitleBarHidden results in a hidden title bar and a full size content window,
// yet the title bar still has the standard window controls (“traffic lights”)
// in the top left.

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
