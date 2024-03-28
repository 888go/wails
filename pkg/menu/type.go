package menu

// Type of the menu item
type Type string

const (
	// TextType 是文本菜单项类型
	TextType Type = "Text"
	// SeparatorType 是菜单项类型中的分隔符类型
	SeparatorType Type = "Separator"
	// SubmenuType 是子菜单类型（菜单项类型）
	SubmenuType Type = "Submenu"
	// CheckboxType 是复选框类型菜单项
	CheckboxType Type = "Checkbox"
	// RadioType 是 Radio 菜单项的类型
	RadioType Type = "Radio"
)
