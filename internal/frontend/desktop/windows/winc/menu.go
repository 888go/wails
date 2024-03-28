//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 */

package winc

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

var (
	nextMenuItemID  uint16 = 3
	actionsByID            = make(map[uint16]*MenuItem)
	shortcut2Action        = make(map[Shortcut]*MenuItem)
	menuItems              = make(map[w32.HMENU][]*MenuItem)
	radioGroups            = make(map[*MenuItem]*RadioGroup)
	initialised     bool
)

var NoShortcut = Shortcut{}

// 主窗口菜单及控件上下文菜单的菜单结构。
// 大多数方法既适用于主窗口菜单，也适用于上下文菜单。
type Menu struct {
	hMenu w32.HMENU
	hwnd  w32.HWND // 如果是上下文菜单，hwnd 可能为 nil。
}

type MenuItem struct {
	hMenu    w32.HMENU
	hSubMenu w32.HMENU // 如果此项自身就是一个子菜单，则该值不为零。

	text     string
	toolTip  string
	image    *Bitmap
	shortcut Shortcut
	enabled  bool

	checkable bool
	checked   bool
	isRadio   bool

	id uint16

	onClick EventManager
}

type RadioGroup struct {
	members []*MenuItem
	hwnd    w32.HWND
}

func NewContextMenu() *MenuItem {
	hMenu := w32.CreatePopupMenu()
	if hMenu == 0 {
		panic("failed CreateMenu")
	}

	item := &MenuItem{
		hMenu:    hMenu,
		hSubMenu: hMenu,
	}
	return item
}

func (m *Menu) Dispose() {
	if m.hMenu != 0 {
		w32.DestroyMenu(m.hMenu)
		m.hMenu = 0
	}
}

func (m *Menu) IsDisposed() bool {
	return m.hMenu == 0
}

func initMenuItemInfoFromAction(mii *w32.MENUITEMINFO, a *MenuItem) {
	mii.CbSize = uint32(unsafe.Sizeof(*mii))
	mii.FMask = w32.MIIM_FTYPE | w32.MIIM_ID | w32.MIIM_STATE | w32.MIIM_STRING
	if a.image != nil {
		mii.FMask |= w32.MIIM_BITMAP
		mii.HbmpItem = a.image.handle
	}
	if a.IsSeparator() {
		mii.FType = w32.MFT_SEPARATOR
	} else {
		mii.FType = w32.MFT_STRING
		var text string
		if s := a.shortcut; s.Key != 0 {
			text = fmt.Sprintf("%s\t%s", a.text, s.String())
			shortcut2Action[a.shortcut] = a
		} else {
			text = a.text
		}
		mii.DwTypeData = syscall.StringToUTF16Ptr(text)
		mii.Cch = uint32(len([]rune(a.text)))
	}
	mii.WID = uint32(a.id)

	if a.Enabled() {
		mii.FState &^= w32.MFS_DISABLED
	} else {
		mii.FState |= w32.MFS_DISABLED
	}

	if a.Checkable() {
		mii.FMask |= w32.MIIM_CHECKMARKS
	}
	if a.Checked() {
		mii.FState |= w32.MFS_CHECKED
	}

	if a.hSubMenu != 0 {
		mii.FMask |= w32.MIIM_SUBMENU
		mii.HSubMenu = a.hSubMenu
	}
}

// 在主窗口上显示菜单。
func (m *Menu) Show() {
	initialised = true
	updateRadioGroups()
	if !w32.DrawMenuBar(m.hwnd) {
		panic("DrawMenuBar failed")
	}
}

// AddSubMenu 返回一个用于作为子菜单以执行 AddItem(s) 操作的项目。
func (m *Menu) AddSubMenu(text string) *MenuItem {
	hSubMenu := w32.CreateMenu()
	if hSubMenu == 0 {
		panic("failed CreateMenu")
	}
	return addMenuItem(m.hMenu, hSubMenu, text, Shortcut{}, nil, false)
}

// 此方法将遍历菜单项，将单选项目分组在一起，构建一个快速访问映射，并设置初始项目
func updateRadioGroups() {

	if !initialised {
		return
	}

	radioItemsChecked := []*MenuItem{}
	radioGroups = make(map[*MenuItem]*RadioGroup)
	var currentRadioGroupMembers []*MenuItem
	// Iterate the menus
	for _, menu := range menuItems {
		menuLength := len(menu)
		for index, menuItem := range menu {
			if menuItem.isRadio {
				currentRadioGroupMembers = append(currentRadioGroupMembers, menuItem)
				if menuItem.checked {
					radioItemsChecked = append(radioItemsChecked, menuItem)
				}

				// If end of menu
				if index == menuLength-1 {
					radioGroup := &RadioGroup{
						members: currentRadioGroupMembers,
						hwnd:    menuItem.hMenu,
					}
					// 将群组保存到radiomap中的每个成员
					for _, member := range currentRadioGroupMembers {
						radioGroups[member] = radioGroup
					}
					currentRadioGroupMembers = []*MenuItem{}
				}
				continue
			}

			// Not a radio item
			if len(currentRadioGroupMembers) > 0 {
				radioGroup := &RadioGroup{
					members: currentRadioGroupMembers,
					hwnd:    menuItem.hMenu,
				}
				// 将群组保存到radiomap中的每个成员
				for _, member := range currentRadioGroupMembers {
					radioGroups[member] = radioGroup
				}
				currentRadioGroupMembers = []*MenuItem{}
			}
		}
	}

	// 启用已勾选的项
	for _, item := range radioItemsChecked {
		radioGroup := radioGroups[item]
		startID := radioGroup.members[0].id
		endID := radioGroup.members[len(radioGroup.members)-1].id
		w32.SelectRadioMenuItem(item.id, startID, endID, radioGroup.hwnd)
	}

}

func (mi *MenuItem) OnClick() *EventManager {
	return &mi.onClick
}

func (mi *MenuItem) AddSeparator() {
	addMenuItem(mi.hSubMenu, 0, "-", Shortcut{}, nil, false)
}

// AddItem 添加普通菜单项。
func (mi *MenuItem) AddItem(text string, shortcut Shortcut) *MenuItem {
	return addMenuItem(mi.hSubMenu, 0, text, shortcut, nil, false)
}

// AddItemCheckable 添加一个普通菜单项，该菜单项可以带有复选标记。
func (mi *MenuItem) AddItemCheckable(text string, shortcut Shortcut) *MenuItem {
	return addMenuItem(mi.hSubMenu, 0, text, shortcut, nil, true)
}

// AddItemRadio 添加一个普通菜单项，该菜单项可以带有选中标记，并且是单选组的一部分。
func (mi *MenuItem) AddItemRadio(text string, shortcut Shortcut) *MenuItem {
	menuItem := addMenuItem(mi.hSubMenu, 0, text, shortcut, nil, true)
	menuItem.isRadio = true
	return menuItem
}

// AddItemWithBitmap 添加具有快捷键和位图的菜单项。
func (mi *MenuItem) AddItemWithBitmap(text string, shortcut Shortcut, image *Bitmap) *MenuItem {
	return addMenuItem(mi.hSubMenu, 0, text, shortcut, image, false)
}

// AddSubMenu 添加子菜单
func (mi *MenuItem) AddSubMenu(text string) *MenuItem {
	hSubMenu := w32.CreatePopupMenu()
	if hSubMenu == 0 {
		panic("failed CreatePopupMenu")
	}
	return addMenuItem(mi.hSubMenu, hSubMenu, text, Shortcut{}, nil, false)
}

// AddItem 向菜单添加项目，将文本设置为 "-" 用作分隔符。
func addMenuItem(hMenu, hSubMenu w32.HMENU, text string, shortcut Shortcut, image *Bitmap, checkable bool) *MenuItem {
	item := &MenuItem{
		hMenu:     hMenu,
		hSubMenu:  hSubMenu,
		text:      text,
		shortcut:  shortcut,
		image:     image,
		enabled:   true,
		id:        nextMenuItemID,
		checkable: checkable,
		isRadio:   false,
		//visible:  true,
	}
	nextMenuItemID++
	actionsByID[item.id] = item
	menuItems[hMenu] = append(menuItems[hMenu], item)

	var mii w32.MENUITEMINFO
	initMenuItemInfoFromAction(&mii, item)

	index := -1
	if !w32.InsertMenuItem(hMenu, uint32(index), true, &mii) {
		panic("InsertMenuItem failed")
	}
	return item
}

func indexInObserver(a *MenuItem) int {
	var idx int
	for _, mi := range menuItems[a.hMenu] {
		if mi == a {
			return idx
		}
		idx++
	}
	return -1
}

func findMenuItemByID(id int) *MenuItem {
	return actionsByID[uint16(id)]
}

func (mi *MenuItem) update() {
	var mii w32.MENUITEMINFO
	initMenuItemInfoFromAction(&mii, mi)

	if !w32.SetMenuItemInfo(mi.hMenu, uint32(indexInObserver(mi)), true, &mii) {
		panic("SetMenuItemInfo failed")
	}
	if mi.isRadio {
		mi.updateRadioGroup()
	}
}

func (mi *MenuItem) IsSeparator() bool { return mi.text == "-" }
func (mi *MenuItem) SetSeparator()     { mi.text = "-" }

func (mi *MenuItem) Enabled() bool     { return mi.enabled }
func (mi *MenuItem) SetEnabled(b bool) { mi.enabled = b; mi.update() }

func (mi *MenuItem) Checkable() bool     { return mi.checkable }
func (mi *MenuItem) SetCheckable(b bool) { mi.checkable = b; mi.update() }

func (mi *MenuItem) Checked() bool { return mi.checked }
func (mi *MenuItem) SetChecked(b bool) {
	if mi.isRadio {
		radioGroup := radioGroups[mi]
		if radioGroup != nil {
			for _, member := range radioGroup.members {
				member.checked = false
			}
		}

	}
	mi.checked = b
	mi.update()
}

func (mi *MenuItem) Text() string     { return mi.text }
func (mi *MenuItem) SetText(s string) { mi.text = s; mi.update() }

func (mi *MenuItem) Image() *Bitmap     { return mi.image }
func (mi *MenuItem) SetImage(b *Bitmap) { mi.image = b; mi.update() }

func (mi *MenuItem) ToolTip() string     { return mi.toolTip }
func (mi *MenuItem) SetToolTip(s string) { mi.toolTip = s; mi.update() }

func (mi *MenuItem) updateRadioGroup() {
	radioGroup := radioGroups[mi]
	if radioGroup == nil {
		return
	}
	startID := radioGroup.members[0].id
	endID := radioGroup.members[len(radioGroup.members)-1].id
	w32.SelectRadioMenuItem(mi.id, startID, endID, radioGroup.hwnd)
}
