//go:build windows
// +build windows

package windows

import (
	"github.com/888go/wails/internal/frontend/desktop/windows/winc"
	"github.com/888go/wails/pkg/menu"
)

var checkboxMap = map[*menu.MenuItem][]*winc.MenuItem{}
var radioGroupMap = map[*menu.MenuItem][]*winc.MenuItem{}

func toggleCheckBox(menuItem *menu.MenuItem) {
	menuItem.X是否选中 = !menuItem.X是否选中
	for _, wincMenu := range checkboxMap[menuItem] {
		wincMenu.SetChecked(menuItem.X是否选中)
	}
}

func addCheckBoxToMap(menuItem *menu.MenuItem, wincMenuItem *winc.MenuItem) {
	if checkboxMap[menuItem] == nil {
		checkboxMap[menuItem] = []*winc.MenuItem{}
	}
	checkboxMap[menuItem] = append(checkboxMap[menuItem], wincMenuItem)
}

func toggleRadioItem(menuItem *menu.MenuItem) {
	menuItem.X是否选中 = !menuItem.X是否选中
	for _, wincMenu := range radioGroupMap[menuItem] {
		wincMenu.SetChecked(menuItem.X是否选中)
	}
}

func addRadioItemToMap(menuItem *menu.MenuItem, wincMenuItem *winc.MenuItem) {
	if radioGroupMap[menuItem] == nil {
		radioGroupMap[menuItem] = []*winc.MenuItem{}
	}
	radioGroupMap[menuItem] = append(radioGroupMap[menuItem], wincMenuItem)
}

// ff:
// menu:
func (w *Window) SetApplicationMenu(menu *menu.Menu) {
	w.applicationMenu = menu
	processMenu(w, menu)
}

func processMenu(window *Window, menu *menu.Menu) {
	mainMenu := window.NewMenu()
	for _, menuItem := range menu.Items {
		submenu := mainMenu.AddSubMenu(menuItem.X显示名称)
		if menuItem.X子菜单 != nil {
			for _, menuItem := range menuItem.X子菜单.Items {
				processMenuItem(submenu, menuItem)
			}
		}
	}
	mainMenu.Show()
}

func processMenuItem(parent *winc.MenuItem, menuItem *menu.MenuItem) {
	if menuItem.X是否隐藏 {
		return
	}
	switch menuItem.X常量_菜单项类型 {
	case menu.X常量_菜单项类型_分隔符:
		parent.AddSeparator()
	case menu.X常量_菜单项类型_文本:
		shortcut := acceleratorToWincShortcut(menuItem.X快捷键)
		newItem := parent.AddItem(menuItem.X显示名称, shortcut)
		// 如果menuItem的Tooltip属性不为空字符串 {
		//	为newItem设置Tooltip属性，值为menuItem的Tooltip属性值
		//}
		if menuItem.X单击回调函数 != nil {
			newItem.OnClick().Bind(func(e *winc.Event) {
				menuItem.X单击回调函数(&menu.CallbackData{
					MenuItem: menuItem,
				})
			})
		}
		newItem.SetEnabled(!menuItem.X是否禁用)

	case menu.X常量_菜单项类型_复选框:
		shortcut := acceleratorToWincShortcut(menuItem.X快捷键)
		newItem := parent.AddItem(menuItem.X显示名称, shortcut)
		newItem.SetCheckable(true)
		newItem.SetChecked(menuItem.X是否选中)
		// 如果menuItem的Tooltip属性不为空字符串 {
		//	为newItem设置Tooltip属性，值为menuItem的Tooltip属性值
		//}
		if menuItem.X单击回调函数 != nil {
			newItem.OnClick().Bind(func(e *winc.Event) {
				toggleCheckBox(menuItem)
				menuItem.X单击回调函数(&menu.CallbackData{
					MenuItem: menuItem,
				})
			})
		}
		newItem.SetEnabled(!menuItem.X是否禁用)
		addCheckBoxToMap(menuItem, newItem)
	case menu.X常量_菜单项类型_单选框:
		shortcut := acceleratorToWincShortcut(menuItem.X快捷键)
		newItem := parent.AddItemRadio(menuItem.X显示名称, shortcut)
		newItem.SetCheckable(true)
		newItem.SetChecked(menuItem.X是否选中)
		// 如果menuItem的Tooltip属性不为空字符串 {
		//	为newItem设置Tooltip属性，值为menuItem的Tooltip属性值
		//}
		if menuItem.X单击回调函数 != nil {
			newItem.OnClick().Bind(func(e *winc.Event) {
				toggleRadioItem(menuItem)
				menuItem.X单击回调函数(&menu.CallbackData{
					MenuItem: menuItem,
				})
			})
		}
		newItem.SetEnabled(!menuItem.X是否禁用)
		addRadioItemToMap(menuItem, newItem)
	case menu.X常量_菜单项类型_子菜单:
		submenu := parent.AddSubMenu(menuItem.X显示名称)
		for _, menuItem := range menuItem.X子菜单.Items {
			processMenuItem(submenu, menuItem)
		}
	}
}

// ff:菜单设置
// menu:菜单
func (f *Frontend) MenuSetApplicationMenu(menu *menu.Menu) {
	f.mainWindow.SetApplicationMenu(menu)
}

// ff:菜单更新
func (f *Frontend) MenuUpdateApplicationMenu() {
	processMenu(f.mainWindow, f.mainWindow.applicationMenu)
}
