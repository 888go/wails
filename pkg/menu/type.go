package menu

// Type of the menu item
type Type string

const (
	// TextType 是文本菜单项类型
	X常量_菜单项类型_文本 Type = "Text"
	// SeparatorType 是菜单项类型中的分隔符类型
	X常量_菜单项类型_分隔符 Type = "Separator"
	// SubmenuType 是子菜单类型（菜单项类型）
	X常量_菜单项类型_子菜单 Type = "Submenu"
	// CheckboxType 是复选框类型菜单项
	X常量_菜单项类型_复选框 Type = "Checkbox"
	// RadioType 是 Radio 菜单项的类型
	X常量_菜单项类型_单选框 Type = "Radio"
)
