package runtime

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
)


// ff:菜单设置
// menu:菜单
// ctx:上下文
func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) {
	frontend := getFrontend(ctx)
	frontend.MenuSetApplicationMenu(menu)
}


// ff:菜单更新
// ctx:上下文
func MenuUpdateApplicationMenu(ctx context.Context) {
	frontend := getFrontend(ctx)
	frontend.MenuUpdateApplicationMenu()
}
