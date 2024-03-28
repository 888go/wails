package menu

type ContextMenu struct {
	ID   string
	Menu *Menu //hs:菜单     
}


// ff:创建上下文菜单
// menu:菜单
// ID:
func NewContextMenu(ID string, menu *Menu) *ContextMenu {
	return &ContextMenu{
		ID:   ID,
		Menu: menu,
	}
}
