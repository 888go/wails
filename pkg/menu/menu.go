package menu

import "github.com/888go/wails/pkg/menu/keys"

type Menu struct {
	Items []*MenuItem
}

func X创建() *Menu {
	return &Menu{}
}

func (m *Menu) X加入(菜单项 *MenuItem) {
	m.Items = append(m.Items, 菜单项)
}

// Merge将会把给定菜单中的项目追加到此菜单中
func (m *Menu) X合并(菜单 *Menu) {
	m.Items = append(m.Items, 菜单.Items...)
}

// AddText 向菜单中添加一个 TextMenu 项
func (m *Menu) X加入文本菜单项(显示名称 string, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	item := X创建文本菜单项2(显示名称, 快捷键, 单击回调函数)
	m.X加入(item)
	return item
}

// AddCheckbox 向菜单中添加一个 CheckboxMenu 项
func (m *Menu) X加入复选框(显示名称 string, 选中 bool, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	item := X创建复选框菜单项(显示名称, 选中, 快捷键, 单击回调函数)
	m.X加入(item)
	return item
}

// AddRadio 向菜单中添加一个单选按钮项目
func (m *Menu) X加入单选框(显示名称 string, 选中 bool, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	item := X创建单选框菜单项(显示名称, 选中, 快捷键, 单击回调函数)
	m.X加入(item)
	return item
}

// AddSeparator 向菜单添加一个分隔符
func (m *Menu) X加入分隔符() {
	item := X创建分隔符菜单项()
	m.X加入(item)
}

func (m *Menu) X加入子菜单(显示名称 string) *Menu {
	submenu := X创建()
	item := X创建子菜单(显示名称, submenu)
	m.X加入(item)
	return submenu
}

func (m *Menu) X加入子菜单最前(菜单项 *MenuItem) {
	m.Items = append([]*MenuItem{菜单项}, m.Items...)
}

func X创建菜单并按菜单项(first *MenuItem, 第一个 ...*MenuItem) *Menu {
	result := X创建()
	result.X加入(first)
	for _, item := range 第一个 {
		result.X加入(item)
	}

	return result
}

func (m *Menu) setParent(menuItem *MenuItem) {
	for _, item := range m.Items {
		item.parent = menuItem
	}
}
