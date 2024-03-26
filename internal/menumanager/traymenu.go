package menumanager

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/leaanthony/go-ansi-parser"

	"github.com/pkg/errors"
	"github.com/888go/wails/pkg/menu"
)

var (
	trayMenuID      int
	trayMenuIDMutex sync.Mutex
)

func generateTrayID() string {
	var idStr string
	trayMenuIDMutex.Lock()
	idStr = strconv.Itoa(trayMenuID)
	trayMenuID++
	trayMenuIDMutex.Unlock()
	return idStr
}

type TrayMenu struct {
	ID               string
	Label            string
	FontSize         int
	FontName         string
	Disabled         bool
	Tooltip          string `json:",omitempty"`
	Image            string
	MacTemplateImage bool
	RGBA             string
	menuItemMap      *MenuItemMap
	menu             *menu.Menu
	ProcessedMenu    *WailsMenu
	trayMenu         *menu.TrayMenu
	StyledLabel      []*ansi.StyledText `json:",omitempty"`
}

func (t *TrayMenu) AsJSON() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewTrayMenu(trayMenu *menu.TrayMenu) *TrayMenu {
	// Parse ANSI text
	var styledLabel []*ansi.StyledText
	tempLabel := trayMenu.X显示名称
	if strings.Contains(tempLabel, "\033[") {
		parsedLabel, err := ansi.Parse(tempLabel)
		if err == nil {
			styledLabel = parsedLabel
		}
	}

	result := &TrayMenu{
		Label:            trayMenu.X显示名称,
		FontName:         trayMenu.X字体名称,
		FontSize:         trayMenu.X字体大小,
		Disabled:         trayMenu.X是否禁用,
		Tooltip:          trayMenu.X提示,
		Image:            trayMenu.X图标名称,
		MacTemplateImage: trayMenu.Mac模板图标,
		menu:             trayMenu.X菜单,
		RGBA:             trayMenu.RGBA,
		menuItemMap:      NewMenuItemMap(),
		trayMenu:         trayMenu,
		StyledLabel:      styledLabel,
	}

	result.menuItemMap.AddMenu(trayMenu.X菜单)
	result.ProcessedMenu = NewWailsMenu(result.menuItemMap, result.menu)

	return result
}

func (m *Manager) OnTrayMenuOpen(id string) {
	trayMenu, ok := m.trayMenus[id]
	if !ok {
		return
	}
	if trayMenu.trayMenu.X打开回调函数 == nil {
		return
	}
	go trayMenu.trayMenu.X打开回调函数()
}

func (m *Manager) OnTrayMenuClose(id string) {
	trayMenu, ok := m.trayMenus[id]
	if !ok {
		return
	}
	if trayMenu.trayMenu.X关闭回调函数 == nil {
		return
	}
	go trayMenu.trayMenu.X关闭回调函数()
}

func (m *Manager) AddTrayMenu(trayMenu *menu.TrayMenu) (string, error) {
	newTrayMenu := NewTrayMenu(trayMenu)

	// Hook up a new ID
	trayID := generateTrayID()
	newTrayMenu.ID = trayID

	// Save the references
	m.trayMenus[trayID] = newTrayMenu
	m.trayMenuPointers[trayMenu] = trayID

	return newTrayMenu.AsJSON()
}

func (m *Manager) GetTrayID(trayMenu *menu.TrayMenu) (string, error) {
	trayID, exists := m.trayMenuPointers[trayMenu]
	if !exists {
		return "", fmt.Errorf("Unable to find menu ID for tray menu!")
	}
	return trayID, nil
}

// SetTrayMenu 更新或创建一个菜单
func (m *Manager) SetTrayMenu(trayMenu *menu.TrayMenu) (string, error) {
	trayID, trayMenuKnown := m.trayMenuPointers[trayMenu]
	if !trayMenuKnown {
		return m.AddTrayMenu(trayMenu)
	}

	// 创建更新后的托盘菜单
	updatedTrayMenu := NewTrayMenu(trayMenu)
	updatedTrayMenu.ID = trayID

	// Save the reference
	m.trayMenus[trayID] = updatedTrayMenu

	return updatedTrayMenu.AsJSON()
}

func (m *Manager) GetTrayMenus() ([]string, error) {
	result := []string{}
	for _, trayMenu := range m.trayMenus {
		JSON, err := trayMenu.AsJSON()
		if err != nil {
			return nil, err
		}
		result = append(result, JSON)
	}

	return result, nil
}

func (m *Manager) UpdateTrayMenuLabel(trayMenu *menu.TrayMenu) (string, error) {
	trayID, trayMenuKnown := m.trayMenuPointers[trayMenu]
	if !trayMenuKnown {
		return "", fmt.Errorf("[UpdateTrayMenuLabel] unknown tray id for tray %s", trayMenu.X显示名称)
	}

	type LabelUpdate struct {
		ID               string
		Label            string `json:",omitempty"`
		FontName         string `json:",omitempty"`
		FontSize         int
		RGBA             string `json:",omitempty"`
		Disabled         bool
		Tooltip          string `json:",omitempty"`
		Image            string `json:",omitempty"`
		MacTemplateImage bool
		StyledLabel      []*ansi.StyledText `json:",omitempty"`
	}

	// Parse ANSI text
	var styledLabel []*ansi.StyledText
	tempLabel := trayMenu.X显示名称
	if strings.Contains(tempLabel, "\033[") {
		parsedLabel, err := ansi.Parse(tempLabel)
		if err == nil {
			styledLabel = parsedLabel
		}
	}

	update := &LabelUpdate{
		ID:               trayID,
		Label:            trayMenu.X显示名称,
		FontName:         trayMenu.X字体名称,
		FontSize:         trayMenu.X字体大小,
		Disabled:         trayMenu.X是否禁用,
		Tooltip:          trayMenu.X提示,
		Image:            trayMenu.X图标名称,
		MacTemplateImage: trayMenu.Mac模板图标,
		RGBA:             trayMenu.RGBA,
		StyledLabel:      styledLabel,
	}

	data, err := json.Marshal(update)
	if err != nil {
		return "", errors.Wrap(err, "[UpdateTrayMenuLabel] ")
	}

	return string(data), nil
}

func (m *Manager) GetContextMenus() ([]string, error) {
	result := []string{}
	for _, contextMenu := range m.contextMenus {
		JSON, err := contextMenu.AsJSON()
		if err != nil {
			return nil, err
		}
		result = append(result, JSON)
	}

	return result, nil
}
