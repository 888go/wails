package runtime

import (
	"context"

	"github.com/888go/wails/pkg/menu"
)


// ff:菜单设置
// menu:菜单
// ctx:上下文
func X菜单设置(上下文 context.Context, 菜单 *menu.Menu) {
	frontend := getFrontend(上下文)
	frontend.MenuSetApplicationMenu(菜单)
}


// ff:菜单更新
// ctx:上下文
func X菜单更新(上下文 context.Context) {
	frontend := getFrontend(上下文)
	frontend.MenuUpdateApplicationMenu()
}
