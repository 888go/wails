package runtime

import "context"


// ff:剪贴板取文本
// ctx:上下文
func X剪贴板取文本(上下文 context.Context) (string, error) {
	appFrontend := getFrontend(上下文)
	return appFrontend.ClipboardGetText()
}


// ff:剪贴板置文本
// text:文本
// ctx:上下文
func X剪贴板置文本(上下文 context.Context, 文本 string) error {
	appFrontend := getFrontend(上下文)
	return appFrontend.ClipboardSetText(文本)
}
