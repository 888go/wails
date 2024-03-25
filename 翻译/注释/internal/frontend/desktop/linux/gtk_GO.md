
<原文开始>
//export handleMenuItemClick
<原文结束>

# <翻译开始>
//export handleMenuItemClick
// 导出handleMenuItemClick函数，供其他语言（如C）调用
# <翻译结束>


<原文开始>
	// Make sure to execute the final callback on a new goroutine otherwise if the callback e.g. tries to open a dialog, the
	// main thread will get blocked and so the message loop blocks. As a result the app will block and shows a
	// "not responding" dialog.
<原文结束>

# <翻译开始>
// 确保在新的goroutine上执行最终的回调函数，否则如果回调函数（例如）尝试打开一个对话框，主线程将会被阻塞，因此消息循环也会阻塞。其结果是应用会被阻塞并显示一个“无响应”的对话框。
# <翻译结束>

