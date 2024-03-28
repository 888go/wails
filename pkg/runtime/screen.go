package runtime

import (
	"context"

	"github.com/888go/wails/internal/frontend"
)

type Screen = frontend.Screen

// ScreenGetAllScreens 返回所有屏幕

// ff:取屏幕信息
// ctx:上下文
func X取屏幕信息(上下文 context.Context) ([]Screen, error) {
	appFrontend := getFrontend(上下文)
	return appFrontend.ScreenGetAll()
}
