package runtime

import (
	"context"

	"github.com/888go/wails/internal/frontend"
)

type Screen = frontend.Screen

// ScreenGetAllScreens 返回所有屏幕
func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	appFrontend := getFrontend(ctx)
	return appFrontend.ScreenGetAll()
}
