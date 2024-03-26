package runtime

import (
	"context"
)

// BrowserOpenURL 使用系统默认浏览器打开指定的url
func X默认浏览器打开url(上下文 context.Context, url string) {
	appFrontend := getFrontend(上下文)
	appFrontend.BrowserOpenURL(url)
}
