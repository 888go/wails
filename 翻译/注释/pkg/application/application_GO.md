
<原文开始>
// Application is the main Wails application
<原文结束>

# <翻译开始>
// Application 是 Wails 主应用程序
# <翻译结束>


<原文开始>
// NewWithOptions creates a new Application with the given options
<原文结束>

# <翻译开始>
// NewWithOptions 使用给定的选项创建一个新的Application
# <翻译结束>


<原文开始>
// New creates a new Application with the default options
<原文结束>

# <翻译开始>
// New 创建一个使用默认选项的新 Application
# <翻译结束>


<原文开始>
// SetApplicationMenu sets the application menu
<原文结束>

# <翻译开始>
// 设置应用菜单 将设置应用程序的菜单
# <翻译结束>


<原文开始>
// Run starts the application
<原文结束>

# <翻译开始>
// Run 启动应用程序
# <翻译结束>


<原文开始>
// Quit will shut down the application
<原文结束>

# <翻译开始>
// Quit 将关闭应用程序
# <翻译结束>


<原文开始>
// Bind the given struct to the application
<原文结束>

# <翻译开始>
// 将给定的结构体绑定到应用程序
# <翻译结束>


<原文开始>
	//application, err := app.CreateApp(a.options)
	//if err != nil {
	//	return err
	//}
	//
	//a.application = application
<原文结束>

# <翻译开始>
	// 创建应用，使用a.options作为参数
	// 如果创建过程中发生错误，则将错误赋值给err
	// application, err := app.CreateApp(a.options)
	// 如果err不为nil（即存在错误）
	// if err != nil {
    	// 则直接返回该错误
	//     return err
	// }
	// 将成功创建的应用赋值给a.application变量
	// a.application = application
# <翻译结束>

