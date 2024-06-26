package menumanager

import (
	"fmt"

	"github.com/888go/wails/pkg/menu"
)

type Manager struct {
	// The application menu.
	applicationMenu          *menu.Menu
	applicationMenuJSON      string
	processedApplicationMenu *WailsMenu

	// 我们的应用程序菜单映射
	applicationMenuItemMap *MenuItemMap

	// Context menus
	contextMenus        map[string]*ContextMenu
	contextMenuPointers map[*menu.ContextMenu]string

	// Tray menu stores
	trayMenus        map[string]*TrayMenu
	trayMenuPointers map[*menu.TrayMenu]string

	// Radio groups
	radioGroups map[*menu.MenuItem][]*menu.MenuItem
}


// ff:
func NewManager() *Manager {
	return &Manager{
		applicationMenuItemMap: NewMenuItemMap(),
		contextMenus:           make(map[string]*ContextMenu),
		contextMenuPointers:    make(map[*menu.ContextMenu]string),
		trayMenus:              make(map[string]*TrayMenu),
		trayMenuPointers:       make(map[*menu.TrayMenu]string),
		radioGroups:            make(map[*menu.MenuItem][]*menu.MenuItem),
	}
}

func (m *Manager) getMenuItemByID(menuMap *MenuItemMap, menuId string) *menu.MenuItem {
	return menuMap.idToMenuItemMap[menuId]
}


// ff:
// parentID:
// menuType:
// data:
// menuID:
func (m *Manager) ProcessClick(menuID string, data string, menuType string, parentID string) error {
	var menuItemMap *MenuItemMap

	switch menuType {
	case "ApplicationMenu":
		menuItemMap = m.applicationMenuItemMap
	case "ContextMenu":
		contextMenu := m.contextMenus[parentID]
		if contextMenu == nil {
			return fmt.Errorf("unknown context menu: %s", parentID)
		}
		menuItemMap = contextMenu.menuItemMap
	case "TrayMenu":
		trayMenu := m.trayMenus[parentID]
		if trayMenu == nil {
			return fmt.Errorf("unknown tray menu: %s", parentID)
		}
		menuItemMap = trayMenu.menuItemMap
	default:
		return fmt.Errorf("unknown menutype: %s", menuType)
	}

	// Get the menu item
	menuItem := menuItemMap.getMenuItemByID(menuID)
	if menuItem == nil {
		return fmt.Errorf("Cannot process menuid %s - unknown", menuID)
	}

	// 这个菜单项是否是复选框？
	if menuItem.X常量_菜单项类型 == menu.X常量_菜单项类型_复选框 {
		// Toggle state
		menuItem.X是否选中 = !menuItem.X是否选中
	}

	if menuItem.X常量_菜单项类型 == menu.X常量_菜单项类型_单选框 {
		println("Toggle radio")
		// Get my radio group
		for _, radioMenuItem := range m.radioGroups[menuItem] {
			radioMenuItem.X是否选中 = (radioMenuItem == menuItem)
		}
	}

	if menuItem.X单击回调函数 == nil {
		// No callback
		return fmt.Errorf("No callback for menu '%s'", menuItem.X显示名称)
	}

	// 创建新的Callback结构体
	callbackData := &menu.CallbackData{
		MenuItem: menuItem,
		// ContextData: data,
	}

	// Call back!
	go menuItem.X单击回调函数(callbackData)

	return nil
}

func (m *Manager) processRadioGroups(processedMenu *WailsMenu, itemMap *MenuItemMap) {
	for _, group := range processedMenu.RadioGroups {
		radioGroupMenuItems := []*menu.MenuItem{}
		for _, member := range group.Members {
			item := m.getMenuItemByID(itemMap, member)
			radioGroupMenuItems = append(radioGroupMenuItems, item)
		}
		for _, radioGroupMenuItem := range radioGroupMenuItems {
			m.radioGroups[radioGroupMenuItem] = radioGroupMenuItems
		}
	}
}
