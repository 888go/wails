package menumanager

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

// MenuItemMap 保存了 menuIDs 和菜单项之间的映射关系
type MenuItemMap struct {
	idToMenuItemMap map[string]*menu.MenuItem
	menuItemToIDMap map[*menu.MenuItem]string

	// 我们使用一个简单的计数器来记录唯一的菜单ID
	menuIDCounter      int64
	menuIDCounterMutex sync.Mutex
}


// ff:
func NewMenuItemMap() *MenuItemMap {
	result := &MenuItemMap{
		idToMenuItemMap: make(map[string]*menu.MenuItem),
		menuItemToIDMap: make(map[*menu.MenuItem]string),
	}

	return result
}


// ff:
// menu:
func (m *MenuItemMap) AddMenu(menu *menu.Menu) {
	if menu == nil {
		return
	}
	for _, item := range menu.Items {
		m.processMenuItem(item)
	}
}


// ff:
func (m *MenuItemMap) Dump() {
	println("idToMenuItemMap:")
	for key, value := range m.idToMenuItemMap {
		fmt.Printf("  %s\t%p\n", key, value)
	}
	println("\nmenuItemToIDMap")
	for key, value := range m.menuItemToIDMap {
		fmt.Printf("  %p\t%s\n", key, value)
	}
}

// GenerateMenuID 为菜单项生成一个唯一的字符串ID
func (m *MenuItemMap) generateMenuID() string {
	m.menuIDCounterMutex.Lock()
	result := strconv.FormatInt(m.menuIDCounter, 10)
	m.menuIDCounter++
	m.menuIDCounterMutex.Unlock()
	return result
}

func (m *MenuItemMap) processMenuItem(item *menu.MenuItem) {
	if item.SubMenu != nil {
		for _, submenuitem := range item.SubMenu.Items {
			m.processMenuItem(submenuitem)
		}
	}

	// 为这个菜单项创建一个唯一的ID
	menuID := m.generateMenuID()

	// Store references
	m.idToMenuItemMap[menuID] = item
	m.menuItemToIDMap[item] = menuID
}

func (m *MenuItemMap) getMenuItemByID(menuId string) *menu.MenuItem {
	return m.idToMenuItemMap[menuId]
}
