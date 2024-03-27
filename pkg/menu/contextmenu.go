package menu

type ContextMenu struct {
	ID   string
	X菜单 *Menu
}

func X创建上下文菜单(ID string, 菜单 *Menu) *ContextMenu {
	return &ContextMenu{
		ID:   ID,
		X菜单: 菜单,
	}
}
