
<原文开始>
// Label is what appears as the menu text
<原文结束>

# <翻译开始>
// Label 是作为菜单文本显示的内容
# <翻译结束>


<原文开始>
	// Role is a predefined menu type
	// Role menu.Role `json:",omitempty"`
	// Accelerator holds a representation of a key binding
<原文结束>

# <翻译开始>
	// Role 是一种预定义的菜单类型
	// Role menu.Role `json:",omitempty"` 	// 这一行表示Role字段使用menu包中的Role类型，并在序列化为JSON时，如果该字段值为空，则忽略（omitempty）
	// Accelerator 用于保存一个键绑定的表示
# <翻译结束>


<原文开始>
// Type of MenuItem, EG: Checkbox, Text, Separator, Radio, Submenu
<原文结束>

# <翻译开始>
// MenuItem的类型，例如：复选框、文本、分隔符、单选按钮、子菜单
# <翻译结束>


<原文开始>
// Disabled makes the item unselectable
<原文结束>

# <翻译开始>
// Disabled 使项目不可选择
# <翻译结束>


<原文开始>
// Hidden ensures that the item is not shown in the menu
<原文结束>

# <翻译开始>
// Hidden 确保该菜单项不会在菜单中显示
# <翻译结束>


<原文开始>
// Checked indicates if the item is selected (used by Checkbox and Radio types only)
<原文结束>

# <翻译开始>
// Checked 表示项目是否被选中（仅用于 Checkbox 和 Radio 类型）
# <翻译结束>


<原文开始>
	// Submenu contains a list of menu items that will be shown as a submenu
	// SubMenu []*MenuItem `json:"SubMenu,omitempty"`
<原文结束>

# <翻译开始>
	// Submenu 包含一个菜单项列表，这些菜单项将作为子菜单显示
	// SubMenu []*MenuItem `json:"SubMenu,omitempty"`
# <翻译结束>


<原文开始>
// Image - base64 image data
<原文结束>

# <翻译开始>
// Image - 基于base64编码的图像数据
# <翻译结束>


<原文开始>
	// Parse ANSI text
	//var styledLabel []*ansi.StyledText
	//tempLabel := menuItem.Label
	//if strings.Contains(tempLabel, "\033[") {
	//	parsedLabel, err := ansi.Parse(menuItem.Label)
	//	if err == nil {
	//		styledLabel = parsedLabel
	//	}
	//}
<原文结束>

# <翻译开始>
	// 解析ANSI文本
	//var styledLabel []*ansi.StyledText 	// 声明一个指向ansi.StyledText结构体的指针切片，用于存储解析后的文本
	//tempLabel := menuItem.Label        	// 将menuItem的Label属性赋值给临时变量tempLabel
	//if strings.Contains(tempLabel, "\033[") {  	// 检查tempLabel是否包含ANSI转义序列（"\033["）
	//	parsedLabel, err := ansi.Parse(menuItem.Label)  	// 使用ansi包中的Parse函数解析原始文本，得到styledLabel和可能的错误信息err
	//	if err == nil {                            	// 如果没有错误发生
	//		styledLabel = parsedLabel              	// 将解析后的文本赋值给styledLabel变量
	//	}
	//}
# <翻译结束>


<原文开始>
// Role:             menuItem.Role,
<原文结束>

# <翻译开始>
// 角色:             menuItem.Role,
# <翻译结束>


<原文开始>
		// BackgroundColour:             menuItem.BackgroundColour,
		// FontSize:         menuItem.FontSize,
		// FontName:         menuItem.FontName,
		// Image:            menuItem.Image,
		// MacTemplateImage: menuItem.MacTemplateImage,
		// MacAlternate:     menuItem.MacAlternate,
		// Tooltip:          menuItem.Tooltip,
		// StyledLabel:      styledLabel,
<原文结束>

# <翻译开始>
		// 背景颜色:             menuItem的背景颜色,
		// 字体大小:         menuItem的字体大小,
		// 字体名称:         menuItem的字体名称,
		// 图像:            menuItem的图像,
		// Mac模板图像:     menuItem的Mac系统模板图像,
		// Mac替代项:       menuItem在Mac系统中的替代项,
		// 工具提示:          menuItem的工具提示文本,
		// 样式化标签:      styledLabel
# <翻译结束>


<原文开始>
// WailsMenu is the original menu with the addition
// of radio groups extracted from the menu data
<原文结束>

# <翻译开始>
// WailsMenu是原始菜单，其中还包含了从菜单数据中提取出来的单选组
# <翻译结束>


<原文开始>
// RadioGroup holds all the members of the same radio group
<原文结束>

# <翻译开始>
// RadioGroup 用于存储同一组单选按钮的所有成员
# <翻译结束>


<原文开始>
// Process the radio groups
<原文结束>

# <翻译开始>
// 处理单选组
# <翻译结束>


<原文开始>
// Loop over top level menus
<原文结束>

# <翻译开始>
// 遍历顶级菜单
# <翻译结束>


<原文开始>
// We need to recurse submenus
<原文结束>

# <翻译开始>
// 我们需要递归子菜单
# <翻译结束>


<原文开始>
// Finalise any current radio groups as they don't trickle down to submenus
<原文结束>

# <翻译开始>
// 结束当前任何无线电组，因为它们不会向下传递到子菜单
# <翻译结束>


<原文开始>
// Process each submenu item
<原文结束>

# <翻译开始>
// 处理每个子菜单项
# <翻译结束>


<原文开始>
// Add the item to the radio group
<原文结束>

# <翻译开始>
// 将项目添加到无线电组
# <翻译结束>


<原文开始>
// If we were processing a radio group, fix up the references
<原文结束>

# <翻译开始>
// 如果我们正在处理一个单选组，修正相关的引用
# <翻译结束>

