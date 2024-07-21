
<原文开始>
// MenuManager manages the menus for the application
<原文结束>

# <翻译开始>
// MenuManager 管理应用程序的菜单
# <翻译结束>


<原文开始>
// Click updates the radio group state based on the item clicked
<原文结束>

# <翻译开始>
// 点击根据所点击的项目更新单选组状态
# <翻译结束>


<原文开始>
// updateMenuItemCallback is called when the menu item needs to be updated in the UI
<原文结束>

# <翻译开始>
// updateMenuItemCallback 当菜单项需要在用户界面中更新时被调用
# <翻译结束>


<原文开始>
// items is a map of all menu items in this menu
<原文结束>

# <翻译开始>
// items 是此菜单中所有菜单项的映射
# <翻译结束>


<原文开始>
// radioGroups tracks which radiogroup a menu item belongs to
<原文结束>

# <翻译开始>
// radioGroups 用于跟踪菜单项所属的单选组
# <翻译结束>


<原文开始>
// Save the reference to the top level menu for this item
<原文结束>

# <翻译开始>
// 保存对该项顶级菜单的引用
# <翻译结束>


<原文开始>
// If this is a radio item, add it to the radio group
<原文结束>

# <翻译开始>
// 如果这是一个单选按钮项，则将其添加到单选组中
# <翻译结束>


<原文开始>
		// If this is not a radio item, or we are processing the last item in the menu,
		// then we need to add the current radio group to the map if it has items
<原文结束>

# <翻译开始>
		// 如果当前项目不是单选按钮项，或者我们正在处理菜单中的最后一个项目，
		// 那么如果有项目的话，我们需要将当前单选组添加到映射中
# <翻译结束>


<原文开始>
// If this item is not in our menu, then we can't process it
<原文结束>

# <翻译开始>
// 如果这个项目不在我们的菜单中，那么我们无法处理它
# <翻译结束>


<原文开始>
// If this is a radio item, then we need to update the radio group
<原文结束>

# <翻译开始>
// 如果这是一个单选按钮项，那么我们需要更新单选组
# <翻译结束>


<原文开始>
// Get the radio groups for this item
<原文结束>

# <翻译开始>
// 获取此项目的单选组
# <翻译结束>


<原文开始>
		// Iterate each radio group this item belongs to and set the checked state
		// of all items apart from the one that was clicked to false
<原文结束>

# <翻译开始>
		// 遍历该选项所属的每个单选组，并将除被点击项之外的所有其他项的选中状态设置为 false
# <翻译结束>


<原文开始>
// if menuitem is a checkbox, then we need to toggle the state
<原文结束>

# <翻译开始>
// 如果menuitem是复选框，那么我们需要切换其状态
# <翻译结束>


<原文开始>
// Set the radio item to checked
<原文结束>

# <翻译开始>
// 设置单选按钮项为选中状态
# <翻译结束>

