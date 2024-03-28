package runtime

import (
	"context"
)

// BrowserOpenURL 使用系统默认浏览器打开指定的url
func BrowserOpenURL(ctx context.Context, url string) {
	appFrontend := getFrontend(ctx)
	appFrontend.BrowserOpenURL(url)
}
