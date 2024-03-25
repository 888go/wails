package menu

import (
	"sync"

	"github.com/888go/wails/pkg/menu/keys"
)

// MenuItem 表示菜单中包含的一个菜单项
type MenuItem struct {
	// Label 是作为菜单文本显示的内容
	Label string
	// Role 是一种预定义的菜单类型
	Role Role
	// Accelerator 保存了一个键绑定的表示形式
	Accelerator *keys.Accelerator
	// MenuItem的类型，例如：复选框、文本、分隔符、单选按钮、子菜单
	Type Type
	// Disabled 使项目不可选择
	Disabled bool
	// Hidden 确保该菜单项不会在菜单中显示
	Hidden bool
	// Checked 表示项目是否被选中（仅用于 Checkbox 和 Radio 类型）
	Checked bool
// Submenu 包含一个菜单项列表，这些菜单项将作为子菜单显示
// SubMenu []*MenuItem `json:"SubMenu,omitempty"`
	SubMenu *Menu

	// 菜单被点击时的回调函数
	Click Callback
	/*
		// Text Colour
		RGBA string

		// Font
		FontSize int
		FontName string

		// Image - 基于base64编码的图像数据
		Image string

		// MacTemplateImage 表示在Mac系统中，这个图片是一个模板图片
		MacTemplateImage bool

		// MacAlternate 表示该条目是上一个菜单项的替代项
		MacAlternate bool

		// Tooltip
		Tooltip string
	*/
	// 这个用于保存菜单项的父级。
	parent *MenuItem

	// 用于在删除元素时进行锁定
	removeLock sync.Mutex
}

// Parent 返回菜单项的父级元素。
// 若该菜单项是顶级菜单，则返回 nil。
func (m *MenuItem) Parent() *MenuItem {
	return m.parent
}

// Append 尝试将给定的菜单项添加到此菜单项的子菜单项中。如果此菜单项不是子菜单，则此方法将不会添加该菜单项，而是直接返回 false。
func (m *MenuItem) Append(item *MenuItem) bool {
	if !m.isSubMenu() {
		return false
	}
	item.parent = m
	m.SubMenu.Append(item)
	return true
}

// Prepend 尝试将给定的菜单项添加到
// 该菜单项的子菜单项前面。如果此菜单项不是一个
// 子菜单，那么该方法将不会添加该项，并且
// 简单地返回 false。
func (m *MenuItem) Prepend(item *MenuItem) bool {
	if !m.isSubMenu() {
		return false
	}
	item.parent = m
	m.SubMenu.Prepend(item)
	return true
}

func (m *MenuItem) Remove() {
	// 遍历我的父节点的子节点
	m.Parent().removeChild(m)
}

func (m *MenuItem) removeChild(item *MenuItem) {
	m.removeLock.Lock()
	for index, child := range m.SubMenu.Items {
		if item == child {
			m.SubMenu.Items = append(m.SubMenu.Items[:index], m.SubMenu.Items[index+1:]...)
		}
	}
	m.removeLock.Unlock()
}

// InsertAfter尝试在父菜单中将给定的项添加到本项之后。如果不存在父菜单（即我们是顶级菜单），则返回false
func (m *MenuItem) InsertAfter(item *MenuItem) bool {
	// 我们需要找到我的父级
	if m.parent == nil {
		return false
	}

	// 让我的父级插入这个项目
	return m.parent.insertNewItemAfterGivenItem(m, item)
}

// InsertBefore尝试在父菜单中将给定的项插入到当前项之前。如果不存在父菜单（即我们是一个顶级菜单），则返回false
func (m *MenuItem) InsertBefore(item *MenuItem) bool {
	// 我们需要找到我的父级
	if m.parent == nil {
		return false
	}

	// 让我的父级插入这个项目
	return m.parent.insertNewItemBeforeGivenItem(m, item)
}

// insertNewItemAfterGivenItem 将在本项的子菜单中，将给定的项目插入到给定目标之后。如果本项不是子菜单，则说明出现了问题 :/
func (m *MenuItem) insertNewItemAfterGivenItem(target *MenuItem,
	newItem *MenuItem,
) bool {
	if !m.isSubMenu() {
		return false
	}

	// 查找目标的索引
	targetIndex := m.getItemIndex(target)
	if targetIndex == -1 {
		return false
	}

	// 向切片中插入元素
	return m.insertItemAtIndex(targetIndex+1, newItem)
}

// insertNewItemBeforeGivenItem将在当前子菜单中将给定的项目插入到给定目标之前。如果我们不是子菜单，则说明出现了错误:/
func (m *MenuItem) insertNewItemBeforeGivenItem(target *MenuItem,
	newItem *MenuItem,
) bool {
	if !m.isSubMenu() {
		return false
	}

	// 查找目标的索引
	targetIndex := m.getItemIndex(target)
	if targetIndex == -1 {
		return false
	}

	// 向切片中插入元素
	return m.insertItemAtIndex(targetIndex, newItem)
}

func (m *MenuItem) isSubMenu() bool {
	return m.Type == SubmenuType
}

// getItemIndex 返回给定目标相对于此菜单的索引
func (m *MenuItem) getItemIndex(target *MenuItem) int {
	// 这个方法应当只在子菜单上调用
	if !m.isSubMenu() {
		return -1
	}

	// hunt down that bad boy
	for index, item := range m.SubMenu.Items {
		if item == target {
			return index
		}
	}

	return -1
}

// insertItemAtIndex 尝试在指定索引处将给定的项目插入到子菜单中
// 来源：https://stackoverflow.com/a/61822301
func (m *MenuItem) insertItemAtIndex(index int, target *MenuItem) bool {
	// 如果索引越界，则返回 false
	if index > len(m.SubMenu.Items) {
		return false
	}

	// Save parent reference
	target.parent = m

	// 如果索引是最后一个项目，则进行常规追加
	if index == len(m.SubMenu.Items) {
		m.SubMenu.Items = append(m.SubMenu.Items, target)
		return true
	}

	m.SubMenu.Items = append(m.SubMenu.Items[:index+1], m.SubMenu.Items[index:]...)
	m.SubMenu.Items[index] = target
	return true
}

func (m *MenuItem) SetLabel(name string) {
	if m.Label == name {
		return
	}
	m.Label = name
}

func (m *MenuItem) IsSeparator() bool {
	return m.Type == SeparatorType
}

func (m *MenuItem) IsCheckbox() bool {
	return m.Type == CheckboxType
}

func (m *MenuItem) Disable() *MenuItem {
	m.Disabled = true
	return m
}

func (m *MenuItem) Enable() *MenuItem {
	m.Disabled = false
	return m
}

func (m *MenuItem) OnClick(click Callback) *MenuItem {
	m.Click = click
	return m
}

func (m *MenuItem) SetAccelerator(acc *keys.Accelerator) *MenuItem {
	m.Accelerator = acc
	return m
}

func (m *MenuItem) SetChecked(value bool) *MenuItem {
	m.Checked = value
	if m.Type != RadioType {
		m.Type = CheckboxType
	}
	return m
}

func (m *MenuItem) Hide() *MenuItem {
	m.Hidden = true
	return m
}

func (m *MenuItem) Show() *MenuItem {
	m.Hidden = false
	return m
}

func (m *MenuItem) IsRadio() bool {
	return m.Type == RadioType
}

func Label(label string) *MenuItem {
	return &MenuItem{
		Type:  TextType,
		Label: label,
	}
}

// Text 是一个辅助函数，用于创建基本的文本菜单项
func Text(label string, accelerator *keys.Accelerator, click Callback) *MenuItem {
	return &MenuItem{
		Label:       label,
		Type:        TextType,
		Accelerator: accelerator,
		Click:       click,
	}
}

// Separator 提供一个菜单分隔符
func Separator() *MenuItem {
	return &MenuItem{
		Type: SeparatorType,
	}
}

// Radio 是一个辅助工具，用于创建带快捷键的基本单选菜单项
func Radio(label string, selected bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	return &MenuItem{
		Label:       label,
		Type:        RadioType,
		Checked:     selected,
		Accelerator: accelerator,
		Click:       click,
	}
}

// Checkbox 是一个辅助函数，用于创建基本的复选框菜单项
func Checkbox(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	return &MenuItem{
		Label:       label,
		Type:        CheckboxType,
		Checked:     checked,
		Accelerator: accelerator,
		Click:       click,
	}
}

// SubMenu 是一个用于创建子菜单的辅助函数
func SubMenu(label string, menu *Menu) *MenuItem {
	result := &MenuItem{
		Label:   label,
		SubMenu: menu,
		Type:    SubmenuType,
	}

	menu.setParent(result)

	return result
}
