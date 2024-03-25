
<原文开始>
// MenuItem represents a menuitem contained in a menu
<原文结束>

# <翻译开始>
// MenuItem 表示菜单中包含的一个菜单项
# <翻译结束>


<原文开始>
// Label is what appears as the menu text
<原文结束>

# <翻译开始>
// Label 是作为菜单文本显示的内容
# <翻译结束>


<原文开始>
// Role is a predefined menu type
<原文结束>

# <翻译开始>
// Role 是一种预定义的菜单类型
# <翻译结束>


<原文开始>
// Accelerator holds a representation of a key binding
<原文结束>

# <翻译开始>
// Accelerator 保存了一个键绑定的表示形式
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
// Callback function when menu clicked
<原文结束>

# <翻译开始>
// 菜单被点击时的回调函数
# <翻译结束>


<原文开始>
// Image - base64 image data
<原文结束>

# <翻译开始>
// Image - 基于base64编码的图像数据
# <翻译结束>


<原文开始>
// MacTemplateImage indicates that on a Mac, this image is a template image
<原文结束>

# <翻译开始>
// MacTemplateImage 表示在Mac系统中，这个图片是一个模板图片
# <翻译结束>


<原文开始>
// MacAlternate indicates that this item is an alternative to the previous menu item
<原文结束>

# <翻译开始>
// MacAlternate 表示该条目是上一个菜单项的替代项
# <翻译结束>


<原文开始>
// This holds the menu item's parent.
<原文结束>

# <翻译开始>
// 这个用于保存菜单项的父级。
# <翻译结束>


<原文开始>
// Used for locking when removing elements
<原文结束>

# <翻译开始>
// 用于在删除元素时进行锁定
# <翻译结束>


<原文开始>
// Parent returns the parent of the menu item.
// If it is a top level menu then it returns nil.
<原文结束>

# <翻译开始>
// Parent 返回菜单项的父级元素。
// 若该菜单项是顶级菜单，则返回 nil。
# <翻译结束>


<原文开始>
// Append will attempt to append the given menu item to
// this item's submenu items. If this menu item is not a
// submenu, then this method will not add the item and
// simply return false.
<原文结束>

# <翻译开始>
// Append 尝试将给定的菜单项添加到此菜单项的子菜单项中。如果此菜单项不是子菜单，则此方法将不会添加该菜单项，而是直接返回 false。
# <翻译结束>


<原文开始>
// Prepend will attempt to prepend the given menu item to
// this item's submenu items. If this menu item is not a
// submenu, then this method will not add the item and
// simply return false.
<原文结束>

# <翻译开始>
// Prepend 尝试将给定的菜单项添加到
// 该菜单项的子菜单项前面。如果此菜单项不是一个
// 子菜单，那么该方法将不会添加该项，并且
// 简单地返回 false。
# <翻译结束>


<原文开始>
// Iterate my parent's children
<原文结束>

# <翻译开始>
// 遍历我的父节点的子节点
# <翻译结束>


<原文开始>
// InsertAfter attempts to add the given item after this item in the parent
// menu. If there is no parent menu (we are a top level menu) then false is
// returned
<原文结束>

# <翻译开始>
// InsertAfter尝试在父菜单中将给定的项添加到本项之后。如果不存在父菜单（即我们是顶级菜单），则返回false
# <翻译结束>


<原文开始>
// We need to find my parent
<原文结束>

# <翻译开始>
// 我们需要找到我的父级
# <翻译结束>


<原文开始>
// Get my parent to insert the item
<原文结束>

# <翻译开始>
// 让我的父级插入这个项目
# <翻译结束>


<原文开始>
// InsertBefore attempts to add the given item before this item in the parent
// menu. If there is no parent menu (we are a top level menu) then false is
// returned
<原文结束>

# <翻译开始>
// InsertBefore尝试在父菜单中将给定的项插入到当前项之前。如果不存在父菜单（即我们是一个顶级菜单），则返回false
# <翻译结束>


<原文开始>
// insertNewItemAfterGivenItem will insert the given item after the given target
// in this item's submenu. If we are not a submenu,
// then something bad has happened :/
<原文结束>

# <翻译开始>
// insertNewItemAfterGivenItem 将在本项的子菜单中，将给定的项目插入到给定目标之后。如果本项不是子菜单，则说明出现了问题 :/
# <翻译结束>


<原文开始>
// Find the index of the target
<原文结束>

# <翻译开始>
// 查找目标的索引
# <翻译结束>


<原文开始>
// Insert element into slice
<原文结束>

# <翻译开始>
// 向切片中插入元素
# <翻译结束>


<原文开始>
// insertNewItemBeforeGivenItem will insert the given item before the given
// target in this item's submenu. If we are not a submenu, then something bad
// has happened :/
<原文结束>

# <翻译开始>
// insertNewItemBeforeGivenItem将在当前子菜单中将给定的项目插入到给定目标之前。如果我们不是子菜单，则说明出现了错误:/
# <翻译结束>


<原文开始>
// getItemIndex returns the index of the given target relative to this menu
<原文结束>

# <翻译开始>
// getItemIndex 返回给定目标相对于此菜单的索引
# <翻译结束>


<原文开始>
// This should only be called on submenus
<原文结束>

# <翻译开始>
// 这个方法应当只在子菜单上调用
# <翻译结束>


<原文开始>
// insertItemAtIndex attempts to insert the given item into the submenu at
// the given index
// Credit: https://stackoverflow.com/a/61822301
<原文结束>

# <翻译开始>
// insertItemAtIndex 尝试在指定索引处将给定的项目插入到子菜单中
// 来源：https://stackoverflow.com/a/61822301
# <翻译结束>


<原文开始>
// If index is OOB, return false
<原文结束>

# <翻译开始>
// 如果索引越界，则返回 false
# <翻译结束>


<原文开始>
// If index is last item, then just regular append
<原文结束>

# <翻译开始>
// 如果索引是最后一个项目，则进行常规追加
# <翻译结束>


<原文开始>
// Text is a helper to create basic Text menu items
<原文结束>

# <翻译开始>
// Text 是一个辅助函数，用于创建基本的文本菜单项
# <翻译结束>


<原文开始>
// Separator provides a menu separator
<原文结束>

# <翻译开始>
// Separator 提供一个菜单分隔符
# <翻译结束>


<原文开始>
// Radio is a helper to create basic Radio menu items with an accelerator
<原文结束>

# <翻译开始>
// Radio 是一个辅助工具，用于创建带快捷键的基本单选菜单项
# <翻译结束>


<原文开始>
// Checkbox is a helper to create basic Checkbox menu items
<原文结束>

# <翻译开始>
// Checkbox 是一个辅助函数，用于创建基本的复选框菜单项
# <翻译结束>


<原文开始>
// SubMenu is a helper to create Submenus
<原文结束>

# <翻译开始>
// SubMenu 是一个用于创建子菜单的辅助函数
# <翻译结束>

