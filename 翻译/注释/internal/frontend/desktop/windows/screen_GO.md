
<原文开始>
	// Checks to make sure all the fields are the same.
	// A cleaner way would be to check identity of devices. but I couldn't find a way of doing that using the win32 API
<原文结束>

# <翻译开始>
// 检查确保所有字段都相同。
// 一个更简洁的方法是检查设备的身份。但我没找到使用win32 API实现该功能的方法。
# <翻译结束>


<原文开始>
	// Adapted from winc.utils.getMonitorInfo TODO: add this to win32
	// See docs for
	//https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmonitorinfoa
<原文结束>

# <翻译开始>
// 该段代码改编自 winc.utils.getMonitorInfo，待办：将其添加至 win32 库中
// 参考文档：
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getmonitorinfoa
// （注：此网址为微软官方文档，关于 Windows API 函数 GetMonitorInfoA 的说明）
# <翻译结束>


<原文开始>
// adapted from https://stackoverflow.com/a/23492886/4188138
<原文结束>

# <翻译开始>
// 该代码段改编自 StackOverflow 网站上的回答：https://stackoverflow.com/a/23492886/4188138
# <翻译结束>


<原文开始>
	// see docs for the following pages to better understand this function
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nc-winuser-monitorenumproc
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-monitorinfo
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-monitorfromwindow
<原文结束>

# <翻译开始>
// 为了更好地理解此函数，请查看以下页面的文档
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors // （中文版链接，可能需要手动添加或替换）
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nc-winuser-monitorenumproc
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/ns-winuser-monitorinfo
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-monitorfromwindow
// 上述链接分别对应：
// - `EnumDisplayMonitors` 函数的说明
// - `MONITORENUMPROC` 函数指针类型的说明
// - `MONITORINFO` 结构体的定义
// - `MonitorFromWindow` 函数的说明
# <翻译结束>


<原文开始>
// not sure what the consequences of returning false are, so let's just return true and handle it ourselves
<原文结束>

# <翻译开始>
// 不确定返回false会带来什么后果，所以我们先返回true并自行处理这个问题
# <翻译结束>


<原文开始>
	// the reason we need a container is that we have don't know how many times this function will be called
	// this "append" call could potentially do an allocation and rewrite the pointer to monitors. So we save the pointer in screenContainer.monitors
	// and retrieve the values after all EnumProc calls
	// If EnumProc is multi-threaded, this could be problematic. Although, I don't think it is.
<原文结束>

# <翻译开始>
// 我们需要一个容器的原因是我们并不知道这个函数会被调用多少次
// 这个 "append" 调用可能潜在地进行内存分配并重写 monitors 指针。所以我们把指针保存在 screenContainer.monitors 中
// 并在所有 EnumProc 调用结束后获取其值
// 如果 EnumProc 是多线程的，这可能会存在问题。尽管我认为并不是这样。
# <翻译结束>


<原文开始>
// let's keep screenContainer.errors the same size as screenContainer.monitors in case we want to match them up later if necessary
<原文结束>

# <翻译开始>
// 让我们保持 screenContainer.errors 与 screenContainer.monitors 的大小相同，以便后续如有必要时可以将它们对应匹配起来
# <翻译结束>


<原文开始>
// TODO fix hack of container sharing by having a proper data sharing mechanism between windows and the runtime
<原文结束>

# <翻译开始>
// TODO：通过在Windows与运行时之间建立合适的数据共享机制，修复容器共享的临时解决方案
# <翻译结束>

