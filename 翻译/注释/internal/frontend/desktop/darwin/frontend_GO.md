
<原文开始>
// this should be initialized as early as possible to handle first instance launch
<原文结束>

# <翻译开始>
// 这个应当尽早初始化，以便处理首次实例启动
# <翻译结束>


<原文开始>
	//if strings.HasPrefix(message, "systemevent:") {
	//	f.processSystemEvent(message)
	//	return
	//}
<原文结束>

# <翻译开始>
// 如果字符串message以"systemevent:"开头 {
//     f.processSystemEvent(message) //调用处理系统事件的方法
//     return //结束当前函数执行
// }
# <翻译结束>


<原文开始>
// Callback from a method call
<原文结束>

# <翻译开始>
// 从方法调用返回的回调函数
# <翻译结束>


<原文开始>
//func (f *Frontend) processSystemEvent(message string) {
//	sl := strings.Split(message, ":")
//	if len(sl) != 2 {
//		f.logger.Error("Invalid system message: %s", message)
//		return
//	}
//	switch sl[1] {
//	case "fullscreen":
//		f.mainWindow.DisableSizeConstraints()
//	case "unfullscreen":
//		f.mainWindow.EnableSizeConstraints()
//	default:
//		f.logger.Error("Unknown system message: %s", message)
//	}
//}
<原文结束>

# <翻译开始>
// 对于函数 (f *Frontend) processSystemEvent(message string) {
// 将接收到的消息字符串message以":"为分隔符进行切割
// 如果切割后的子串数量不为2，则
// 记录错误日志，输出无效的系统消息，并直接返回
// 根据第二个子串（下标为1）的值进行判断并执行相应操作
// 若第二个子串为 "fullscreen"
// 则禁用主窗口大小的约束限制
// 若第二个子串为 "unfullscreen"
// 则启用主窗口大小的约束限制
// 若上述情况均不符合
// 记录错误日志，输出未知的系统消息
// }
# <翻译结束>


<原文开始>
//export processURLRequest
<原文结束>

# <翻译开始>
//export processURLRequest
// 导出processURLRequest函数（供C语言调用）
// 在Golang的cgo中，`//export`关键字用于声明一个Go函数，表示该函数可供C代码通过C ABI（应用程序二进制接口）进行调用。因此，这段注释翻译为：
// 声明导出函数processURLRequest，以便C代码能够调用
# <翻译结束>


<原文开始>
//export HandleCustomProtocol
<原文结束>

# <翻译开始>
//export HandleCustomProtocol
// 导出HandleCustomProtocol函数，以便在C语言或其他外部环境中调用
// （该注释表明这个Go函数是用于被C语言或者其他需要与Go互操作的环境调用的，通过`//export`标记，Go构建工具会生成相应的导出符号，使得该函数可以在其他语言中被调用。）
# <翻译结束>

