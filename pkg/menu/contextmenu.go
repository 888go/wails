package menu

type ContextMenu struct {
	ID   string
	X菜单 *Menu //hs:菜单     
}


// ff:创建上下文菜单
// menu:菜单
// ID:
func X创建上下文菜单(ID string, 菜单 *Menu) *ContextMenu {
	return &ContextMenu{
		ID:   ID,
		X菜单: 菜单,
	}
}
