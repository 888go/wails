package menumanager

import "github.com/888go/wails/pkg/menu"


// ff:
// applicationMenu:
func (m *Manager) SetApplicationMenu(applicationMenu *menu.Menu) error {
	if applicationMenu == nil {
		return nil
	}

	m.applicationMenu = applicationMenu

	// Reset the menu map
	m.applicationMenuItemMap = NewMenuItemMap()

	// 将菜单添加到菜单映射中
	m.applicationMenuItemMap.AddMenu(applicationMenu)

	return m.processApplicationMenu()
}


// ff:
func (m *Manager) GetApplicationMenuJSON() string {
	return m.applicationMenuJSON
}


// ff:
func (m *Manager) GetProcessedApplicationMenu() *WailsMenu {
	return m.processedApplicationMenu
}

// UpdateApplicationMenu 重新处理应用程序菜单以获取结构变化等信息
// 返回更新后菜单的 JSON 表示

// ff:
func (m *Manager) UpdateApplicationMenu() (string, error) {
	m.applicationMenuItemMap = NewMenuItemMap()
	m.applicationMenuItemMap.AddMenu(m.applicationMenu)
	err := m.processApplicationMenu()
	return m.applicationMenuJSON, err
}

func (m *Manager) processApplicationMenu() error {
	// Process the menu
	m.processedApplicationMenu = NewWailsMenu(m.applicationMenuItemMap, m.applicationMenu)
	m.processRadioGroups(m.processedApplicationMenu, m.applicationMenuItemMap)
	applicationMenuJSON, err := m.processedApplicationMenu.AsJSON()
	if err != nil {
		return err
	}
	m.applicationMenuJSON = applicationMenuJSON
	return nil
}
