package runtime

import "context"

func X剪贴板取文本(上下文 context.Context) (string, error) {
	appFrontend := getFrontend(上下文)
	return appFrontend.ClipboardGetText()
}

func X剪贴板置文本(上下文 context.Context, 文本 string) error {
	appFrontend := getFrontend(上下文)
	return appFrontend.ClipboardSetText(文本)
}
