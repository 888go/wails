package runtime

import (
	"context"

	"github.com/wailsapp/wails/v2/internal/frontend"
)

type Screen = frontend.Screen

// ScreenGetAllScreens 返回所有屏幕

// ff:取屏幕信息
// ctx:上下文
func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	appFrontend := getFrontend(ctx)
	return appFrontend.ScreenGetAll()
}
