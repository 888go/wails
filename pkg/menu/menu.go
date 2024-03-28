package menu

import "github.com/wailsapp/wails/v2/pkg/menu/keys"

type Menu struct {
	Items []*MenuItem
}


// ff:创建
func NewMenu() *Menu {
	return &Menu{}
}


// ff:加入
// item:菜单项
func (m *Menu) Append(item *MenuItem) {
	m.Items = append(m.Items, item)
}

// Merge将会把给定菜单中的项目追加到此菜单中

// ff:合并
// menu:菜单
func (m *Menu) Merge(menu *Menu) {
	m.Items = append(m.Items, menu.Items...)
}

// AddText 向菜单中添加一个 TextMenu 项

// ff:加入文本菜单项
// click:单击回调函数
// accelerator:快捷键
// label:显示名称
func (m *Menu) AddText(label string, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Text(label, accelerator, click)
	m.Append(item)
	return item
}

// AddCheckbox 向菜单中添加一个 CheckboxMenu 项

// ff:加入复选框
// click:单击回调函数
// accelerator:快捷键
// checked:选中
// label:显示名称
func (m *Menu) AddCheckbox(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Checkbox(label, checked, accelerator, click)
	m.Append(item)
	return item
}

// AddRadio 向菜单中添加一个单选按钮项目

// ff:加入单选框
// click:单击回调函数
// accelerator:快捷键
// checked:选中
// label:显示名称
func (m *Menu) AddRadio(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Radio(label, checked, accelerator, click)
	m.Append(item)
	return item
}

// AddSeparator 向菜单添加一个分隔符

// ff:加入分隔符
func (m *Menu) AddSeparator() {
	item := Separator()
	m.Append(item)
}


// ff:加入子菜单
// label:显示名称
func (m *Menu) AddSubmenu(label string) *Menu {
	submenu := NewMenu()
	item := SubMenu(label, submenu)
	m.Append(item)
	return submenu
}


// ff:加入子菜单最前
// item:菜单项
func (m *Menu) Prepend(item *MenuItem) {
	m.Items = append([]*MenuItem{item}, m.Items...)
}


// ff:创建菜单并按菜单项
// rest:第一个
// first:
func NewMenuFromItems(first *MenuItem, rest ...*MenuItem) *Menu {
	result := NewMenu()
	result.Append(first)
	for _, item := range rest {
		result.Append(item)
	}

	return result
}

func (m *Menu) setParent(menuItem *MenuItem) {
	for _, item := range m.Items {
		item.parent = menuItem
	}
}
