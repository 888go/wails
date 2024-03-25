package menu

import (
	"sync"

	"github.com/888go/wails/pkg/menu/keys"
)

// MenuItem 表示菜单中包含的一个菜单项
type MenuItem struct {
	// Label 是作为菜单文本显示的内容
	X显示名称 string
	// Role 是一种预定义的菜单类型
	X项角色 Role
	// Accelerator 保存了一个键绑定的表示形式
	X快捷键 *keys.Accelerator
	// MenuItem的类型，例如：复选框、文本、分隔符、单选按钮、子菜单
	X常量_菜单项类型 Type
	// Disabled 使项目不可选择
	X是否禁用 bool
	// Hidden 确保该菜单项不会在菜单中显示
	X是否隐藏 bool
	// Checked 表示项目是否被选中（仅用于 Checkbox 和 Radio 类型）
	X是否选中 bool
// Submenu 包含一个菜单项列表，这些菜单项将作为子菜单显示
// SubMenu []*MenuItem `json:"SubMenu,omitempty"`
	X子菜单 *Menu

	// 菜单被点击时的回调函数
	X单击回调函数 Callback
	/*
		// Text Colour
		RGBA string

		// Font
		FontSize int
		FontName string

		// Image - 基于base64编码的图像数据
		Image string

		// MacTemplateImage 表示在Mac系统中，这个图片是一个模板图片
		MacTemplateImage bool

		// MacAlternate 表示该条目是上一个菜单项的替代项
		MacAlternate bool

		// Tooltip
		Tooltip string
	*/
	// 这个用于保存菜单项的父级。
	parent *MenuItem

	// 用于在删除元素时进行锁定
	removeLock sync.Mutex
}

// Parent 返回菜单项的父级元素。
// 若该菜单项是顶级菜单，则返回 nil。
func (m *MenuItem) X取父菜单() *MenuItem {
	return m.parent
}

// Append 尝试将给定的菜单项添加到此菜单项的子菜单项中。如果此菜单项不是子菜单，则此方法将不会添加该菜单项，而是直接返回 false。
func (m *MenuItem) X加入子菜单(菜单项 *MenuItem) bool {
	if !m.isSubMenu() {
		return false
	}
	菜单项.parent = m
	m.X子菜单.X加入(菜单项)
	return true
}

// Prepend 尝试将给定的菜单项添加到
// 该菜单项的子菜单项前面。如果此菜单项不是一个
// 子菜单，那么该方法将不会添加该项，并且
// 简单地返回 false。
func (m *MenuItem) X加入子菜单最前(菜单项 *MenuItem) bool {
	if !m.isSubMenu() {
		return false
	}
	菜单项.parent = m
	m.X子菜单.X加入子菜单最前(菜单项)
	return true
}

func (m *MenuItem) X删除() {
	// 遍历我的父节点的子节点
	m.X取父菜单().removeChild(m)
}

func (m *MenuItem) removeChild(item *MenuItem) {
	m.removeLock.Lock()
	for index, child := range m.X子菜单.Items {
		if item == child {
			m.X子菜单.Items = append(m.X子菜单.Items[:index], m.X子菜单.Items[index+1:]...)
		}
	}
	m.removeLock.Unlock()
}

// InsertAfter尝试在父菜单中将给定的项添加到本项之后。如果不存在父菜单（即我们是顶级菜单），则返回false
func (m *MenuItem) X插入当前后面(菜单项 *MenuItem) bool {
	// 我们需要找到我的父级
	if m.parent == nil {
		return false
	}

	// 让我的父级插入这个项目
	return m.parent.insertNewItemAfterGivenItem(m, 菜单项)
}

// InsertBefore尝试在父菜单中将给定的项插入到当前项之前。如果不存在父菜单（即我们是一个顶级菜单），则返回false
func (m *MenuItem) X插入当前前面(菜单项 *MenuItem) bool {
	// 我们需要找到我的父级
	if m.parent == nil {
		return false
	}

	// 让我的父级插入这个项目
	return m.parent.insertNewItemBeforeGivenItem(m, 菜单项)
}

// insertNewItemAfterGivenItem 将在本项的子菜单中，将给定的项目插入到给定目标之后。如果本项不是子菜单，则说明出现了问题 :/
func (m *MenuItem) insertNewItemAfterGivenItem(target *MenuItem,
	newItem *MenuItem,
) bool {
	if !m.isSubMenu() {
		return false
	}

	// 查找目标的索引
	targetIndex := m.getItemIndex(target)
	if targetIndex == -1 {
		return false
	}

	// 向切片中插入元素
	return m.insertItemAtIndex(targetIndex+1, newItem)
}

// insertNewItemBeforeGivenItem将在当前子菜单中将给定的项目插入到给定目标之前。如果我们不是子菜单，则说明出现了错误:/
func (m *MenuItem) insertNewItemBeforeGivenItem(target *MenuItem,
	newItem *MenuItem,
) bool {
	if !m.isSubMenu() {
		return false
	}

	// 查找目标的索引
	targetIndex := m.getItemIndex(target)
	if targetIndex == -1 {
		return false
	}

	// 向切片中插入元素
	return m.insertItemAtIndex(targetIndex, newItem)
}

func (m *MenuItem) isSubMenu() bool {
	return m.X常量_菜单项类型 == X常量_菜单项类型_子菜单
}

// getItemIndex 返回给定目标相对于此菜单的索引
func (m *MenuItem) getItemIndex(target *MenuItem) int {
	// 这个方法应当只在子菜单上调用
	if !m.isSubMenu() {
		return -1
	}

	// hunt down that bad boy
	for index, item := range m.X子菜单.Items {
		if item == target {
			return index
		}
	}

	return -1
}

// insertItemAtIndex 尝试在指定索引处将给定的项目插入到子菜单中
// 来源：https://stackoverflow.com/a/61822301
func (m *MenuItem) insertItemAtIndex(index int, target *MenuItem) bool {
	// 如果索引越界，则返回 false
	if index > len(m.X子菜单.Items) {
		return false
	}

	// Save parent reference
	target.parent = m

	// 如果索引是最后一个项目，则进行常规追加
	if index == len(m.X子菜单.Items) {
		m.X子菜单.Items = append(m.X子菜单.Items, target)
		return true
	}

	m.X子菜单.Items = append(m.X子菜单.Items[:index+1], m.X子菜单.Items[index:]...)
	m.X子菜单.Items[index] = target
	return true
}

func (m *MenuItem) X设置显示名称(名称 string) {
	if m.X显示名称 == 名称 {
		return
	}
	m.X显示名称 = 名称
}

func (m *MenuItem) X是否为分隔符() bool {
	return m.X常量_菜单项类型 == X常量_菜单项类型_分隔符
}

func (m *MenuItem) X是否为复选框() bool {
	return m.X常量_菜单项类型 == X常量_菜单项类型_复选框
}

func (m *MenuItem) X设置禁用() *MenuItem {
	m.X是否禁用 = true
	return m
}

func (m *MenuItem) X取消禁用() *MenuItem {
	m.X是否禁用 = false
	return m
}

func (m *MenuItem) X绑定单击事件(回调函数 Callback) *MenuItem {
	m.X单击回调函数 = 回调函数
	return m
}

func (m *MenuItem) X设置快捷键(快捷键 *keys.Accelerator) *MenuItem {
	m.X快捷键 = 快捷键
	return m
}

func (m *MenuItem) X设置选中(选中 bool) *MenuItem {
	m.X是否选中 = 选中
	if m.X常量_菜单项类型 != X常量_菜单项类型_单选框 {
		m.X常量_菜单项类型 = X常量_菜单项类型_复选框
	}
	return m
}

func (m *MenuItem) X设置隐藏() *MenuItem {
	m.X是否隐藏 = true
	return m
}

func (m *MenuItem) X取消隐藏() *MenuItem {
	m.X是否隐藏 = false
	return m
}

func (m *MenuItem) X是否为菜单项() bool {
	return m.X常量_菜单项类型 == X常量_菜单项类型_单选框
}

func X创建文本菜单项(显示名称 string) *MenuItem {
	return &MenuItem{
		X常量_菜单项类型:  X常量_菜单项类型_文本,
		X显示名称: 显示名称,
	}
}

// Text 是一个辅助函数，用于创建基本的文本菜单项
func X创建文本菜单项2(显示名称 string, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	return &MenuItem{
		X显示名称:       显示名称,
		X常量_菜单项类型:        X常量_菜单项类型_文本,
		X快捷键: 快捷键,
		X单击回调函数:       单击回调函数,
	}
}

// Separator 提供一个菜单分隔符
func X创建分隔符菜单项() *MenuItem {
	return &MenuItem{
		X常量_菜单项类型: X常量_菜单项类型_分隔符,
	}
}

// Radio 是一个辅助工具，用于创建带快捷键的基本单选菜单项
func X创建单选框菜单项(显示名称 string, 选中 bool, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	return &MenuItem{
		X显示名称:       显示名称,
		X常量_菜单项类型:        X常量_菜单项类型_单选框,
		X是否选中:     选中,
		X快捷键: 快捷键,
		X单击回调函数:       单击回调函数,
	}
}

// Checkbox 是一个辅助函数，用于创建基本的复选框菜单项
func X创建复选框菜单项(显示名称 string, 选中 bool, 快捷键 *keys.Accelerator, 单击回调函数 Callback) *MenuItem {
	return &MenuItem{
		X显示名称:       显示名称,
		X常量_菜单项类型:        X常量_菜单项类型_复选框,
		X是否选中:     选中,
		X快捷键: 快捷键,
		X单击回调函数:       单击回调函数,
	}
}

// SubMenu 是一个用于创建子菜单的辅助函数
func X创建子菜单(显示名称 string, 子菜单 *Menu) *MenuItem {
	result := &MenuItem{
		X显示名称:   显示名称,
		X子菜单: 子菜单,
		X常量_菜单项类型:    X常量_菜单项类型_子菜单,
	}

	子菜单.setParent(result)

	return result
}
