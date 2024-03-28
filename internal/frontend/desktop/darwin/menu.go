//go:build darwin
// +build darwin

package darwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa -framework WebKit
#import <Foundation/Foundation.h>
#import "Application.h"
#import "WailsContext.h"

#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/888go/wails/pkg/menu"
	"github.com/888go/wails/pkg/menu/keys"
)

type NSMenu struct {
	context unsafe.Pointer
	nsmenu  unsafe.Pointer
}


// ff:
// name:
// context:
func NewNSMenu(context unsafe.Pointer, name string) *NSMenu {
	c := NewCalloc()
	defer c.Free()
	title := c.String(name)
	nsmenu := C.NewMenu(title)
	return &NSMenu{
		context: context,
		nsmenu:  nsmenu,
	}
}


// ff:
// label:
func (m *NSMenu) AddSubMenu(label string) *NSMenu {
	result := NewNSMenu(m.context, label)
	C.AppendSubmenu(m.nsmenu, result.nsmenu)
	return result
}


// ff:
// role:
func (m *NSMenu) AppendRole(role menu.Role) {
	C.AppendRole(m.context, m.nsmenu, C.int(role))
}

type MenuItem struct {
	id                uint
	nsmenuitem        unsafe.Pointer
	wailsMenuItem     *menu.MenuItem
	radioGroupMembers []*MenuItem
}


// ff:
// menuItem:
func (m *NSMenu) AddMenuItem(menuItem *menu.MenuItem) *MenuItem {
	c := NewCalloc()
	defer c.Free()
	var modifier C.int
	var key *C.char
	if menuItem.Accelerator != nil {
		modifier = C.int(keys.ToMacModifier(menuItem.Accelerator))
		key = c.String(menuItem.Accelerator.Key)
	}

	result := &MenuItem{
		wailsMenuItem: menuItem,
	}

	result.id = createMenuItemID(result)
	result.nsmenuitem = C.AppendMenuItem(m.context, m.nsmenu, c.String(menuItem.Label), key, modifier, bool2Cint(menuItem.Disabled), bool2Cint(menuItem.Checked), C.int(result.id))
	return result
}

// SetApplicationMenu 函数用于设置窗口（Window）的应用程序菜单。
// 参数:
//   w: 指向 Window 结构体的指针，表示当前窗口实例
//   menu: 指向 menu.Menu 结构体的指针，表示要设置的应用程序菜单
//
// 功能:
//   将传入的菜单（menu）设置为窗口（w）的应用程序菜单，并进一步处理该菜单。
//   
// 实现细节:
//   1. 将传入的菜单赋值给窗口的 applicationMenu 成员变量
//   2. 调用 processMenu 函数来处理这个新设置的菜单及其相关操作
// ```go
// (窗口 *Window) 设置应用程序菜单(菜单 *menu.Menu) {
//     窗口.applicationMenu = 菜单
//     处理菜单(窗口, 菜单)
//}

func processMenu(parent *NSMenu, wailsMenu *menu.Menu) {
	var radioGroups []*MenuItem

	for _, menuItem := range wailsMenu.Items {
		if menuItem.SubMenu != nil {
			if len(radioGroups) > 0 {
				processRadioGroups(radioGroups)
				radioGroups = []*MenuItem{}
			}
			submenu := parent.AddSubMenu(menuItem.Label)
			processMenu(submenu, menuItem.SubMenu)
		} else {
			lastMenuItem := processMenuItem(parent, menuItem)
			if menuItem.Type == menu.RadioType {
				radioGroups = append(radioGroups, lastMenuItem)
			} else {
				if len(radioGroups) > 0 {
					processRadioGroups(radioGroups)
					radioGroups = []*MenuItem{}
				}
			}
		}
	}
}

func processRadioGroups(groups []*MenuItem) {
	for _, item := range groups {
		item.radioGroupMembers = groups
	}
}

func processMenuItem(parent *NSMenu, menuItem *menu.MenuItem) *MenuItem {
	if menuItem.Hidden {
		return nil
	}
	if menuItem.Role != 0 {
		parent.AppendRole(menuItem.Role)
		return nil
	}
	if menuItem.Type == menu.SeparatorType {
		C.AppendSeparator(parent.nsmenu)
		return nil
	}

	return parent.AddMenuItem(menuItem)
}


// ff:
// menu:
func (f *Frontend) MenuSetApplicationMenu(menu *menu.Menu) {
	f.mainWindow.SetApplicationMenu(menu)
}


// ff:
func (f *Frontend) MenuUpdateApplicationMenu() {
	f.mainWindow.UpdateApplicationMenu()
}
