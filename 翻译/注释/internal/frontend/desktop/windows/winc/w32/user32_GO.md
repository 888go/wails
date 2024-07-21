
<原文开始>
//procSysColorBrush            = moduser32.NewProc("GetSysColorBrush")
<原文结束>

# <翻译开始>
// procSysColorBrush 是一个指向用户32模块中名为"GetSysColorBrush"的新进程的指针
# <翻译结束>


<原文开始>
//procSetMenu                  = moduser32.NewProc("SetMenu")
<原文结束>

# <翻译开始>
// procSetMenu 是一个指向 "SetMenu" 函数的指针，该函数位于 moduser32 包中
// moduser32.NewProc 用于从用户32模块动态加载并创建指向 "SetMenu" API 函数的指针
// procSetMenu                  = moduser32.NewProc("SetMenu") 
// 这一行代码定义了一个名为 procSetMenu 的变量，它通过 moduser32 包获取并指向 Windows API 中的 "SetMenu" 函数。
# <翻译结束>


<原文开始>
	//procDrawMenuBar     = moduser32.NewProc("DrawMenuBar")
	//procInsertMenuItem                = moduser32.NewProc("InsertMenuItemW") // FIXIT:
<原文结束>

# <翻译开始>
	// procDrawMenuBar = moduser32.NewProc("DrawMenuBar") 	// 创建并获取名为"DrawMenuBar"的用户32模块（moduser32）中的新过程
	// procInsertMenuItem = moduser32.NewProc("InsertMenuItemW") 	// 创建并获取名为"InsertMenuItemW"的用户32模块中的新过程 	// 注意：待修复或进一步处理
# <翻译结束>


<原文开始>
// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
<原文结束>

# <翻译开始>
// 这是微软官方文档链接，指向Windows Win32 API中关于CheckMenuRadioItem函数的说明：
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
// （该链接描述的是Windows API中的一个函数，用于在菜单栏中切换单选按钮式的菜单项。）
# <翻译结束>

