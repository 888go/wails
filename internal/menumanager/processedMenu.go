package menumanager

import (
	"encoding/json"

	"github.com/888go/wails/pkg/menu"
	"github.com/888go/wails/pkg/menu/keys"
)

type ProcessedMenuItem struct {
	ID string
	// Label 是作为菜单文本显示的内容
	Label string `json:",omitempty"`
// Role 是一种预定义的菜单类型
// Role menu.Role `json:",omitempty"` // 这一行表示Role字段使用menu包中的Role类型，并在序列化为JSON时，如果该字段值为空，则忽略（omitempty）
// Accelerator 用于保存一个键绑定的表示
	Accelerator *keys.Accelerator `json:",omitempty"`
	// MenuItem的类型，例如：复选框、文本、分隔符、单选按钮、子菜单
	Type menu.Type
	// Disabled 使项目不可选择
	Disabled bool `json:",omitempty"`
	// Hidden 确保该菜单项不会在菜单中显示
	Hidden bool `json:",omitempty"`
	// Checked 表示项目是否被选中（仅用于 Checkbox 和 Radio 类型）
	Checked bool `json:",omitempty"`
// Submenu 包含一个菜单项列表，这些菜单项将作为子菜单显示
// SubMenu []*MenuItem `json:"SubMenu,omitempty"`
	SubMenu *ProcessedMenu `json:",omitempty"`
	/*
		// Colour
		RGBA string `json:",omitempty"`

		// Font
		FontSize int    `json:",omitempty"`
		FontName string `json:",omitempty"`

		// Image - 基于base64编码的图像数据
		Image            string `json:",omitempty"`
		MacTemplateImage bool   `json:", omitempty"`
		MacAlternate     bool   `json:", omitempty"`

		// Tooltip
		Tooltip string `json:",omitempty"`

		// Styled label
		StyledLabel []*ansi.StyledText `json:",omitempty"`
	*/
}


// ff:
// menuItem:
// menuItemMap:
func NewProcessedMenuItem(menuItemMap *MenuItemMap, menuItem *menu.MenuItem) *ProcessedMenuItem {
	ID := menuItemMap.menuItemToIDMap[menuItem]

// 解析ANSI文本
//var styledLabel []*ansi.StyledText // 声明一个指向ansi.StyledText结构体的指针切片，用于存储解析后的文本
//tempLabel := menuItem.Label        // 将menuItem的Label属性赋值给临时变量tempLabel
//if strings.Contains(tempLabel, "\033[") {  // 检查tempLabel是否包含ANSI转义序列（"\033["）
//	parsedLabel, err := ansi.Parse(menuItem.Label)  // 使用ansi包中的Parse函数解析原始文本，得到styledLabel和可能的错误信息err
//	if err == nil {                            // 如果没有错误发生
//		styledLabel = parsedLabel              // 将解析后的文本赋值给styledLabel变量
//	}
//}

	result := &ProcessedMenuItem{
		ID:    ID,
		Label: menuItem.X显示名称,
		// 角色:             menuItem.Role,
		Accelerator: menuItem.X快捷键,
		Type:        menuItem.X常量_菜单项类型,
		Disabled:    menuItem.X是否禁用,
		Hidden:      menuItem.X是否隐藏,
		Checked:     menuItem.X是否选中,
		SubMenu:     nil,
// 背景颜色:             menuItem的背景颜色,
// 字体大小:         menuItem的字体大小,
// 字体名称:         menuItem的字体名称,
// 图像:            menuItem的图像,
// Mac模板图像:     menuItem的Mac系统模板图像,
// Mac替代项:       menuItem在Mac系统中的替代项,
// 工具提示:          menuItem的工具提示文本,
// 样式化标签:      styledLabel
	}

	if menuItem.X子菜单 != nil {
		result.SubMenu = NewProcessedMenu(menuItemMap, menuItem.X子菜单)
	}

	return result
}

type ProcessedMenu struct {
	Items []*ProcessedMenuItem
}


// ff:
// menu:
// menuItemMap:
func NewProcessedMenu(menuItemMap *MenuItemMap, menu *menu.Menu) *ProcessedMenu {
	result := &ProcessedMenu{}
	if menu != nil {
		for _, item := range menu.Items {
			processedMenuItem := NewProcessedMenuItem(menuItemMap, item)
			result.Items = append(result.Items, processedMenuItem)
		}
	}

	return result
}

// WailsMenu是原始菜单，其中还包含了从菜单数据中提取出来的单选组
type WailsMenu struct {
	Menu              *ProcessedMenu
	RadioGroups       []*RadioGroup
	currentRadioGroup []string
}

// RadioGroup 用于存储同一组单选按钮的所有成员
type RadioGroup struct {
	Members []string
	Length  int
}


// ff:
// menu:
// menuItemMap:
func NewWailsMenu(menuItemMap *MenuItemMap, menu *menu.Menu) *WailsMenu {
	result := &WailsMenu{}

	// Process the menus
	result.Menu = NewProcessedMenu(menuItemMap, menu)

	// 处理单选组
	result.processRadioGroups()

	return result
}


// ff:
func (w *WailsMenu) AsJSON() (string, error) {
	menuAsJSON, err := json.Marshal(w)
	if err != nil {
		return "", err
	}
	return string(menuAsJSON), nil
}

func (w *WailsMenu) processRadioGroups() {
	// 遍历顶级菜单
	for _, item := range w.Menu.Items {
		// Process MenuItem
		w.processMenuItem(item)
	}

	w.finaliseRadioGroup()
}

func (w *WailsMenu) processMenuItem(item *ProcessedMenuItem) {
	switch item.Type {

	// 我们需要递归子菜单
	case menu.X常量_菜单项类型_子菜单:

		// 结束当前任何无线电组，因为它们不会向下传递到子菜单
		w.finaliseRadioGroup()

		// 处理每个子菜单项
		for _, subitem := range item.SubMenu.Items {
			w.processMenuItem(subitem)
		}
	case menu.X常量_菜单项类型_单选框:
		// 将项目添加到无线电组
		w.currentRadioGroup = append(w.currentRadioGroup, item.ID)
	default:
		w.finaliseRadioGroup()
	}
}

func (w *WailsMenu) finaliseRadioGroup() {
	// 如果我们正在处理一个单选组，修正相关的引用
	if len(w.currentRadioGroup) > 0 {

		// Create new radiogroup
		group := &RadioGroup{
			Members: w.currentRadioGroup,
			Length:  len(w.currentRadioGroup),
		}
		w.RadioGroups = append(w.RadioGroups, group)

		// Empty the radio group
		w.currentRadioGroup = []string{}
	}
}
