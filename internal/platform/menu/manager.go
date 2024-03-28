//go:build windows

package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
)

// MenuManager 管理应用程序的菜单
var MenuManager = NewManager()

type radioGroup []*menu.MenuItem

// 点击根据所点击的项目更新单选组状态

// ff:
// item:
func (g *radioGroup) Click(item *menu.MenuItem) {
	for _, radioGroupItem := range *g {
		if radioGroupItem != item {
			radioGroupItem.Checked = false
		}
	}
}

type processedMenu struct {

	// the menu we processed
	menu *menu.Menu

	// updateMenuItemCallback 当菜单项需要在用户界面中更新时被调用
	updateMenuItemCallback func(*menu.MenuItem)

	// items 是此菜单中所有菜单项的映射
	items map[*menu.MenuItem]struct{}

	// radioGroups 用于跟踪菜单项所属的单选组
	radioGroups map[*menu.MenuItem][]*radioGroup
}

func newProcessedMenu(topLevelMenu *menu.Menu, updateMenuItemCallback func(*menu.MenuItem)) *processedMenu {
	result := &processedMenu{
		updateMenuItemCallback: updateMenuItemCallback,
		menu:                   topLevelMenu,
		items:                  make(map[*menu.MenuItem]struct{}),
		radioGroups:            make(map[*menu.MenuItem][]*radioGroup),
	}
	result.process(topLevelMenu.Items)
	return result
}

func (p *processedMenu) process(items []*menu.MenuItem) {
	var currentRadioGroup radioGroup
	for index, item := range items {
		// 保存对该项顶级菜单的引用
		p.items[item] = struct{}{}

		// 如果这是一个单选按钮项，则将其添加到单选组中
		if item.Type == menu.RadioType {
			currentRadioGroup = append(currentRadioGroup, item)
		}

// 如果当前项目不是单选按钮项，或者我们正在处理菜单中的最后一个项目，
// 那么如果有项目的话，我们需要将当前单选组添加到映射中
		if item.Type != menu.RadioType || index == len(items)-1 {
			if len(currentRadioGroup) > 0 {
				p.addRadioGroup(currentRadioGroup)
				currentRadioGroup = nil
			}
		}

		// Process the submenu
		if item.SubMenu != nil {
			p.process(item.SubMenu.Items)
		}
	}
}

func (p *processedMenu) processClick(item *menu.MenuItem) {
	// 如果这个项目不在我们的菜单中，那么我们无法处理它
	if _, ok := p.items[item]; !ok {
		return
	}

	// 如果这是一个单选按钮项，那么我们需要更新单选组
	if item.Type == menu.RadioType {
		// 获取此项目的单选组
		radioGroups := p.radioGroups[item]
// 遍历该选项所属的每个单选组，并将除被点击项之外的所有其他项的选中状态设置为 false
		for _, thisRadioGroup := range radioGroups {
			thisRadioGroup.Click(item)
			for _, thisRadioGroupItem := range *thisRadioGroup {
				p.updateMenuItemCallback(thisRadioGroupItem)
			}
		}
	}

	if item.Type == menu.CheckboxType {
		p.updateMenuItemCallback(item)
	}

}

func (p *processedMenu) addRadioGroup(r radioGroup) {
	for _, item := range r {
		p.radioGroups[item] = append(p.radioGroups[item], &r)
	}
}

type Manager struct {
	menus map[*menu.Menu]*processedMenu
}


// ff:
func NewManager() *Manager {
	return &Manager{
		menus: make(map[*menu.Menu]*processedMenu),
	}
}


// ff:
// updateMenuItemCallback:
// menu:
func (m *Manager) AddMenu(menu *menu.Menu, updateMenuItemCallback func(*menu.MenuItem)) {
	m.menus[menu] = newProcessedMenu(menu, updateMenuItemCallback)
}


// ff:
// item:
func (m *Manager) ProcessClick(item *menu.MenuItem) {

	// 如果menuitem是复选框，那么我们需要切换其状态
	if item.Type == menu.CheckboxType {
		item.Checked = !item.Checked
	}

	// 设置单选按钮项为选中状态
	if item.Type == menu.RadioType {
		item.Checked = true
	}

	for _, thisMenu := range m.menus {
		thisMenu.processClick(item)
	}

	if item.Click != nil {
		item.Click(&menu.CallbackData{
			MenuItem: item,
		})
	}
}


// ff:
// data:
func (m *Manager) RemoveMenu(data *menu.Menu) {
	delete(m.menus, data)
}
