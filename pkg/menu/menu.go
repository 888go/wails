package menu

import "github.com/888go/wails/pkg/menu/keys"

type Menu struct {
	Items []*MenuItem
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) Append(item *MenuItem) {
	m.Items = append(m.Items, item)
}

// Merge将会把给定菜单中的项目追加到此菜单中
func (m *Menu) Merge(menu *Menu) {
	m.Items = append(m.Items, menu.Items...)
}

// AddText 向菜单中添加一个 TextMenu 项
func (m *Menu) AddText(label string, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Text(label, accelerator, click)
	m.Append(item)
	return item
}

// AddCheckbox 向菜单中添加一个 CheckboxMenu 项
func (m *Menu) AddCheckbox(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Checkbox(label, checked, accelerator, click)
	m.Append(item)
	return item
}

// AddRadio 向菜单中添加一个单选按钮项目
func (m *Menu) AddRadio(label string, checked bool, accelerator *keys.Accelerator, click Callback) *MenuItem {
	item := Radio(label, checked, accelerator, click)
	m.Append(item)
	return item
}

// AddSeparator 向菜单添加一个分隔符
func (m *Menu) AddSeparator() {
	item := Separator()
	m.Append(item)
}

func (m *Menu) AddSubmenu(label string) *Menu {
	submenu := NewMenu()
	item := SubMenu(label, submenu)
	m.Append(item)
	return submenu
}

func (m *Menu) Prepend(item *MenuItem) {
	m.Items = append([]*MenuItem{item}, m.Items...)
}

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
