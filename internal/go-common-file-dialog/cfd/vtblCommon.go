//go:build windows
// +build windows

package cfd

type comDlgFilterSpec struct {
	pszName *int16
	pszSpec *int16
}

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type iModalWindowVtbl struct {
	iUnknownVtbl
	Show uintptr // 对 hwndOwner（HWND 类型）这个窗口句柄调用此函数，将返回 HRESULT 类型的结果。
}

type iFileDialogVtbl struct {
	iModalWindowVtbl
	SetFileTypes        uintptr // 函数定义：(cFileTypes UINT 类型, rgFilterSpec *COMDLG_FILTERSPEC 类型) HRESULT 返回值类型
// 
// 参数：
// - cFileTypes: UINT 类型，表示文件类型数量
// - rgFilterSpec：指向 COMDLG_FILTERSPEC 结构体的指针，用于指定文件过滤器规范（如：显示在打开或保存对话框中的文件类型列表）
// 返回值：
// - HRESULT：一个表示函数调用结果的返回值类型，在Windows API中常用，用来表示成功、失败及错误信息
	SetFileTypeIndex    uintptr // 函数：(iFileType UINT) HRESULT
// 
// 参数：
// - iFileType：UINT 类型，表示文件类型
// 
// 返回值：
// - HRESULT：一个返回结果值，通常用于表示 Windows API 调用的成功或失败状态。在 COM（Component Object Model）编程中，HRESULT 是一种标准的错误代码格式。
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr // 对FILEOPENDIALOGOPTIONS类型的fos变量进行方法扩展，返回一个HRESULT值
	GetOptions          uintptr // 对结构体FILEOPENDIALOGOPTIONS的指针pfos进行方法封装，返回一个HRESULT类型的值
// ```go
// (pfos *FILEOPENDIALOGOPTIONS) HRESULT
// 函数接收一个指向FILEOPENDIALOGOPTIONS结构体的指针pfos，
// 执行相关操作后，返回一个HRESULT类型的值，该类型在Windows API中通常用于表示函数调用的结果状态。
	SetDefaultFolder    uintptr // 对于IShellItem类型的指针psi，定义一个HRESULT函数
	SetFolder           uintptr // 对于IShellItem类型的指针psi，定义一个HRESULT函数
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr // 函数声明：（pszName LPCWSTR 类型的参数）返回 HRESULT 类型
// 在 Go 语言中，这段代码并不是一个完整的函数声明，它似乎是 C++/COM 接口中的函数指针定义在 Golang 中的一种表示方式。
// 如果按照 C++ 的语境翻译，该注释的大致含义为：
// 定义一个函数，其参数为类型为 LPCWSTR 的 pszName，返回值类型为 HRESULT。
// 其中：
// LPCWSTR 是 Long Pointer to Constant Wide String（长指针常量宽字符串），在 Windows 环境下用于表示 Unicode 字符串；
// HRESULT 是 COM（Component Object Model）编程模型中广泛使用的一种返回值类型，表示调用方法或函数的结果状态。
	GetFileName         uintptr
	SetTitle            uintptr // 函数（pszTitle LPCWSTR） HRESULT
// （注：该注释仅描述了函数签名，没有提供更多上下文信息，故翻译可能不够精确）
// 
// 函数名：未给出
// 
// 参数：
// - pszTitle：类型为LPCWSTR的参数，表示一个宽字符指针，常用于存储标题字符串（在Windows编程环境下常见）
// 
// 返回值：
// - HRESULT：在COM和Windows API中常用的一种返回值类型，表示函数执行结果。其具体含义根据不同的函数实现有所不同，一般用来表示成功或失败以及错误的具体信息。
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr // 对函数的注释翻译：
// 函数功能：为指向IShellItem接口指针的指针(ppsi)提供一个HRESULT返回值的方法
// 原始注释：
// func (ppsi **IShellItem) HRESULT
// 在Windows编程中，`IShellItem`是一个用于表示文件系统或虚拟项的COM接口。这个函数声明可能是在定义一个接收者为指向`IShellItem`接口指针的指针的方法，并且该方法返回一个HRESULT，通常用来表示函数调用的成功或失败以及其他的错误信息。但在给出的代码片段中，缺少了方法的具体实现和名称。
	AddPlace            uintptr
	SetDefaultExtension uintptr // 函数（pszDefaultExtension LPCWSTR 类型参数）返回 HRESULT
// 
// 这个注释是对一个Go语言函数的描述，该函数接受一个名为 pszDefaultExtension 的 LPCWSTR 类型参数，并返回 HRESULT 类型的结果。LPCWSTR 是 Windows API 中常用的类型，表示指向宽字符（wchar_t）的长指针，通常用于处理Unicode字符串；而 HRESULT 是 COM（Component Object Model）和一些 Windows API 中使用的一种返回值类型，用来表示函数调用的结果状态。
// 翻译成中文后，可能如下：
// 函数（pszDefaultExtension：LPCWSTR 类型参数）
// 返回：HRESULT 类型
// 
// 本函数接收一个名为 pszDefaultExtension 的 LPCWSTR 类型参数（通常是一个 Unicode 字符串指针），并返回一个 HRESULT 类型的结果，用于表示函数调用的状态或错误信息。
	// 这只能在回调函数中使用。
	Close           uintptr
	SetClientGuid   uintptr // 函数定义：为（REFGUID 类型的）guid 参数返回 HRESULT
	ClearClientData uintptr
	SetFilter       uintptr
}
