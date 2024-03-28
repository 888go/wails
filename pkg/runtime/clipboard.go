package runtime

import "context"


// ff:剪贴板取文本
// ctx:上下文
func ClipboardGetText(ctx context.Context) (string, error) {
	appFrontend := getFrontend(ctx)
	return appFrontend.ClipboardGetText()
}


// ff:剪贴板置文本
// text:文本
// ctx:上下文
func ClipboardSetText(ctx context.Context, text string) error {
	appFrontend := getFrontend(ctx)
	return appFrontend.ClipboardSetText(text)
}
