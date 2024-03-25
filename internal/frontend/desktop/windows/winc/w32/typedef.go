//go:build windows

/*
 * Copyright (C) 2019 Tad Vizbaras. All Rights Reserved.
 * Copyright (C) 2010-2012 The W32 Authors. All Rights Reserved.
 */

package w32

import (
	"fmt"
	"unsafe"
)

// From MSDN: Windows Data Types
// http://msdn.microsoft.com/en-us/library/s3f49ktz.aspx
// http://msdn.microsoft.com/en-us/library/windows/desktop/aa383751.aspx
// ATOM                  WORD
// BOOL                  int32
// BOOLEAN               byte
// BYTE                  byte
// CCHAR                 int8
// CHAR                  int8
// COLORREF              DWORD
// DWORD                 uint32
// DWORDLONG             ULONGLONG
// DWORD_PTR             ULONG_PTR
// DWORD32               uint32
// DWORD64               uint64
// FLOAT                 float32
// HACCEL                HANDLE
// HALF_PTR              struct{} // ???
// HANDLE                PVOID
// HBITMAP               HANDLE
// HBRUSH                HANDLE
// HCOLORSPACE           HANDLE
// HCONV                 HANDLE
// HCONVLIST             HANDLE
// HCURSOR               HANDLE
// HDC                   HANDLE
// HDDEDATA              HANDLE
// HDESK                 HANDLE
// HDROP                 HANDLE
// HDWP                  HANDLE
// HENHMETAFILE          HANDLE
// HFILE                 HANDLE
// HFONT                 HANDLE
// HGDIOBJ               HANDLE
// HGLOBAL               HANDLE
// HHOOK                 HANDLE
// HICON                 HANDLE
// HINSTANCE             HANDLE
// HKEY                  HANDLE
// HKL                   HANDLE
// HLOCAL                HANDLE
// HMENU                 HANDLE
// HMETAFILE             HANDLE
// HMODULE               HANDLE
// HPALETTE              HANDLE
// HPEN                  HANDLE
// HRESULT               int32
// HRGN                  HANDLE
// HSZ                   HANDLE
// HWINSTA               HANDLE
// HWND                  HANDLE
// INT                   int32
// INT_PTR               uintptr
// INT8                  int8
// INT16                 int16
// INT32                 int32
// INT64                 int64
// LANGID                WORD
// LCID                  DWORD
// LCTYPE                DWORD
// LGRPID                DWORD
// LONG                  int32
// LONGLONG              int64
// LONG_PTR              uintptr
// LONG32                int32
// LONG64                int64
// LPARAM                LONG_PTR
// LPBOOL                *BOOL
// LPBYTE                *BYTE
// LPCOLORREF            *COLORREF
// LPCSTR                *int8
// LPCTSTR               LPCWSTR
// LPCVOID               unsafe.Pointer
// LPCWSTR               *WCHAR
// LPDWORD               *DWORD
// LPHANDLE              *HANDLE
// LPINT                 *INT
// LPLONG                *LONG
// LPSTR                 *CHAR
// LPTSTR                LPWSTR
// LPVOID                unsafe.Pointer
// LPWORD                *WORD
// LPWSTR                *WCHAR
// LRESULT               LONG_PTR
// PBOOL                 *BOOL
// PBOOLEAN              *BOOLEAN
// PBYTE                 *BYTE
// PCHAR                 *CHAR
// PCSTR                 *CHAR
// PCTSTR                PCWSTR
// PCWSTR                *WCHAR
// PDWORD                *DWORD
// PDWORDLONG            *DWORDLONG
// PDWORD_PTR            *DWORD_PTR
// PDWORD32              *DWORD32
// PDWORD64              *DWORD64
// PFLOAT                *FLOAT
// PHALF_PTR             *HALF_PTR
// PHANDLE               *HANDLE
// PHKEY                 *HKEY
// PINT_PTR              *INT_PTR
// PINT8                 *INT8
// PINT16                *INT16
// PINT32                *INT32
// PINT64                *INT64
// PLCID                 *LCID
// PLONG                 *LONG
// PLONGLONG             *LONGLONG
// PLONG_PTR             *LONG_PTR
// PLONG32               *LONG32
// PLONG64               *LONG64
// POINTER_32            struct{} // ???
// POINTER_64            struct{} // ???
// POINTER_SIGNED        uintptr
// POINTER_UNSIGNED      uintptr
// PSHORT                *SHORT
// PSIZE_T               *SIZE_T
// PSSIZE_T              *SSIZE_T
// PSTR                  *CHAR
// PTBYTE                *TBYTE
// PTCHAR                *TCHAR
// PTSTR                 PWSTR
// PUCHAR                *UCHAR
// PUHALF_PTR            *UHALF_PTR
// PUINT                 *UINT
// PUINT_PTR             *UINT_PTR
// PUINT8                *UINT8
// PUINT16               *UINT16
// PUINT32               *UINT32
// PUINT64               *UINT64
// PULONG                *ULONG
// PULONGLONG            *ULONGLONG
// PULONG_PTR            *ULONG_PTR
// PULONG32              *ULONG32
// PULONG64              *ULONG64
// PUSHORT               *USHORT
// PVOID                 unsafe.Pointer
// PWCHAR                *WCHAR
// PWORD                 *WORD
// PWSTR                 *WCHAR
// QWORD                 uint64
// SC_HANDLE             HANDLE
// SC_LOCK               LPVOID
// SERVICE_STATUS_HANDLE HANDLE
// SHORT                 int16
// SIZE_T                ULONG_PTR
// SSIZE_T               LONG_PTR
// TBYTE                 WCHAR
// TCHAR                 WCHAR
// UCHAR                 uint8
// UHALF_PTR             struct{} // ???
// UINT                  uint32
// UINT_PTR              uintptr
// UINT8                 uint8
// UINT16                uint16
// UINT32                uint32
// UINT64                uint64
// ULONG                 uint32
// ULONGLONG             uint64
// ULONG_PTR             uintptr
// ULONG32               uint32
// ULONG64               uint64
// USHORT                uint16
// USN                   LONGLONG
// WCHAR                 uint16
// WORD                  uint16
// WPARAM                UINT_PTR
type (
	ATOM            = uint16
	BOOL            = int32
	COLORREF        = uint32
	DWM_FRAME_COUNT = uint64
	WORD            = uint16
	DWORD           = uint32
	HACCEL          = HANDLE
	HANDLE          = uintptr
	HBITMAP         = HANDLE
	HBRUSH          = HANDLE
	HCURSOR         = HANDLE
	HDC             = HANDLE
	HDROP           = HANDLE
	HDWP            = HANDLE
	HENHMETAFILE    = HANDLE
	HFONT           = HANDLE
	HGDIOBJ         = HANDLE
	HGLOBAL         = HANDLE
	HGLRC           = HANDLE
	HHOOK           = HANDLE
	HICON           = HANDLE
	HIMAGELIST      = HANDLE
	HINSTANCE       = HANDLE
	HKEY            = HANDLE
	HKL             = HANDLE
	HMENU           = HANDLE
	HMODULE         = HANDLE
	HMONITOR        = HANDLE
	HPEN            = HANDLE
	HRESULT         = int32
	HRGN            = HANDLE
	HRSRC           = HANDLE
	HTHUMBNAIL      = HANDLE
	HWND            = HANDLE
	LPARAM          = uintptr
	LPCVOID         = unsafe.Pointer
	LRESULT         = uintptr
	PVOID           = unsafe.Pointer
	QPC_TIME        = uint64
	ULONG_PTR       = uintptr
	SIZE_T          = ULONG_PTR
	WPARAM          = uintptr
	UINT            = uint
)

// 参考文献：《Windows桌面开发文档》，网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd162805.aspx
type POINT struct {
	X, Y int32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/dd162897.aspx （MSDN官网关于某个Windows桌面开发API的详细说明）
type RECT struct {
	Left, Top, Right, Bottom int32
}

func (r *RECT) String() string {
	return fmt.Sprintf("RECT (%p): Left: %d, Top: %d, Right: %d, Bottom: %d", r, r.Left, r.Top, r.Right, r.Bottom)
}

// 参考文献：《Windows桌面编程》，MSDN官方文档，网址为：http://msdn.microsoft.com/en-us/library/windows/desktop/ms633577.aspx
type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type TPMPARAMS struct {
	CbSize    uint32
	RcExclude RECT
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms644958.aspx
// （由于您提供的代码片段中并没有具体的代码，故此处仅为参考链接的中文翻译说明，该链接为Microsoft官方文档，介绍了Windows桌面应用程序开发中关于“进程和线程管理”的相关API函数。）
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-minmaxinfo
// 此处是微软官方文档链接，指向Windows API中关于MINMAXINFO结构体的详细说明。
// winuser是Windows用户界面API的一个命名空间，ns-winuser表示该链接描述的是winuser命名空间下的一个类型定义。
// MINMAXINFO结构体在Windows编程中用于存储窗口最小化、最大化以及还原时的相关信息。
type MINMAXINFO struct {
	PtReserved     POINT
	PtMaxSize      POINT
	PtMaxPosition  POINT
	PtMinTrackSize POINT
	PtMaxTrackSize POINT
}

// 参考文档：《Windows桌面开发》, MSDN在线图书馆，网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145037.aspx
type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [LF_FACESIZE]uint16
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms646839.aspx （Windows 操作系统桌面版 API 文档）
type OPENFILENAME struct {
	StructSize      uint32
	Owner           HWND
	Instance        HINSTANCE
	Filter          *uint16
	CustomFilter    *uint16
	MaxCustomFilter uint32
	FilterIndex     uint32
	File            *uint16
	MaxFile         uint32
	FileTitle       *uint16
	MaxFileTitle    uint32
	InitialDir      *uint16
	Title           *uint16
	Flags           uint32
	FileOffset      uint16
	FileExtension   uint16
	DefExt          *uint16
	CustData        uintptr
	FnHook          uintptr
	TemplateName    *uint16
	PvReserved      unsafe.Pointer
	DwReserved      uint32
	FlagsEx         uint32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/bb773205.aspx
type BROWSEINFO struct {
	Owner        HWND
	Root         *uint16
	DisplayName  *uint16
	Title        *uint16
	Flags        uint32
	CallbackFunc uintptr
	LParam       uintptr
	Image        int32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/aa373931.aspx 
// （由于您提供的Golang代码片段中没有具体的代码内容，故此处仅对引用的文档链接进行了翻译。原链接为微软官方文档关于某个Windows桌面开发相关的技术说明。）
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms221627.aspx 
// （由于没有提供具体的代码行，此处仅为参考链接的中文翻译，该链接为微软官方文档关于某个Windows桌面开发的技术指南。）
type VARIANT struct {
	VT         uint16 //  2
	WReserved1 uint16 //  4
	WReserved2 uint16 //  6
	WReserved3 uint16 //  8
	Val        int64  // 16
}

// 参考网址：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms221416.aspx
type DISPPARAMS struct {
	Rgvarg            uintptr
	RgdispidNamedArgs uintptr
	CArgs             uint32
	CNamedArgs        uint32
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/ms221133.aspx
// （由于该注释仅为一个参考链接，未针对具体代码进行解释，故无法给出更详尽的中文翻译。但根据提供的网址链接，可知这是微软官方文档中关于Windows桌面应用程序开发的一个技术说明页面。）
type EXCEPINFO struct {
	WCode             uint16
	WReserved         uint16
	BstrSource        *uint16
	BstrDescription   *uint16
	BstrHelpFile      *uint16
	DwHelpContext     uint32
	PvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int32
}

// 参考网址：https://msdn.microsoft.com/en-us/library/windows/desktop/dd145035.aspx
// （由于您提供的代码片段中没有具体的代码，此处仅为网页链接的注释翻译，该链接指向的是微软官方文档关于某个Windows桌面开发相关的API或指南。）
type LOGBRUSH struct {
	LbStyle uint32
	LbColor COLORREF
	LbHatch uintptr
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/dd183565.aspx 
// （由于您给出的代码片段中并没有具体的代码内容，故此注释仅翻译了引用的参考链接地址，该链接为微软官方文档关于某个Windows桌面开发相关的技术说明。）
type DEVMODE struct {
	DmDeviceName       [CCHDEVICENAME]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint32
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor            int16
	DmDuplex           int16
	DmYResolution      int16
	DmTTOption         int16
	DmCollate          int16
	DmFormName         [CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint32
	DmPelsWidth        uint32
	DmPelsHeight       uint32
	DmDisplayFlags     uint32
	DmDisplayFrequency uint32
	DmICMMethod        uint32
	DmICMIntent        uint32
	DmMediaType        uint32
	DmDitherType       uint32
	DmReserved1        uint32
	DmReserved2        uint32
	DmPanningWidth     uint32
	DmPanningHeight    uint32
}

// 参考文档：《Windows桌面平台开发指南》,位于msdn.microsoft.com/en-us/library/windows/desktop/dd183376.aspx页面
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd162938.aspx
// （由于您提供的代码片段中没有具体的代码内容，这里仅对提供的URL链接进行了翻译，该链接为微软官方文档关于某个Windows桌面开发相关的页面。）
type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

// 参考网址：https://msdn.microsoft.com/zh-cn/library/windows/desktop/dd183375.aspx
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

// 参考文献：《Windows桌面应用程序开发》，位于微软MSDN官网，网址为：http://msdn.microsoft.com/en-us/library/windows/desktop/dd183371.aspx
type BITMAP struct {
	BmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       unsafe.Pointer
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd183567.aspx
type DIBSECTION struct {
	DsBm        BITMAP
	DsBmih      BITMAPINFOHEADER
	DsBitfields [3]uint32
	DshSection  HANDLE
	DsOffset    uint32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/dd162607.aspx
type ENHMETAHEADER struct {
	IType          uint32
	NSize          uint32
	RclBounds      RECT
	RclFrame       RECT
	DSignature     uint32
	NVersion       uint32
	NBytes         uint32
	NRecords       uint32
	NHandles       uint16
	SReserved      uint16
	NDescription   uint32
	OffDescription uint32
	NPalEntries    uint32
	SzlDevice      SIZE
	SzlMillimeters SIZE
	CbPixelFormat  uint32
	OffPixelFormat uint32
	BOpenGL        uint32
	SzlMicrometers SIZE
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145106.aspx
// （由于您提供的代码片段中并没有具体的代码，故这里仅对注释进行翻译，该注释为一个参考链接，指向微软开发者网络（MSDN）上关于某个Windows桌面开发相关主题的技术文档。）
type SIZE struct {
	CX, CY int32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145132.aspx
// （由于您提供的代码片段中没有具体的代码，此处仅为链接的中文描述翻译，该链接指向的是微软官方文档关于某个Windows桌面开发的技术指南。）
type TEXTMETRIC struct {
	TmHeight           int32
	TmAscent           int32
	TmDescent          int32
	TmInternalLeading  int32
	TmExternalLeading  int32
	TmAveCharWidth     int32
	TmMaxCharWidth     int32
	TmWeight           int32
	TmOverhang         int32
	TmDigitizedAspectX int32
	TmDigitizedAspectY int32
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
}

// 参考文献：[微软官方文档]（https://msdn.microsoft.com/en-us/library/windows/desktop/dd183574.aspx）
// 这段注释引用了微软官方的Windows桌面开发文档，由于没有提供具体的代码内容，所以这里仅对网址进行翻译解释，该网址指向的是关于某个Windows桌面开发相关的技术说明或API接口文档。
type DOCINFO struct {
	CbSize       int32
	LpszDocName  *uint16
	LpszOutput   *uint16
	LpszDatatype *uint16
	FwType       uint32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/bb775514.aspx
// （由于没有提供具体的代码，此处仅对引用的网址进行了翻译）
// 该网址为微软开发者网络（MSDN）上关于Windows桌面应用程序开发的一个页面文档，
// 文档内容涉及某种功能或API的详细说明和使用指南。
type NMHDR struct {
	HwndFrom HWND
	IdFrom   uintptr
	Code     uint32
}

// 参考文档：https://msdn.microsoft.com/zh-cn/library/windows/desktop/bb774743.aspx
type LVCOLUMN struct {
	Mask       uint32
	Fmt        int32
	Cx         int32
	PszText    *uint16
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
}

// 参考文献：《Windows桌面编程》，MSDN官方文档，网址：http://msdn.microsoft.com/en-us/library/windows/desktop/bb774760.aspx
type LVITEM struct {
	Mask       uint32
	IItem      int32
	ISubItem   int32
	State      uint32
	StateMask  uint32
	PszText    *uint16
	CchTextMax int32
	IImage     int32
	LParam     uintptr
	IIndent    int32
	IGroupId   int32
	CColumns   uint32
	PuColumns  uint32
}

type LVFINDINFO struct {
	Flags       uint32
	PszText     *uint16
	LParam      uintptr
	Pt          POINT
	VkDirection uint32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/bb774754.aspx
type LVHITTESTINFO struct {
	Pt       POINT
	Flags    uint32
	IItem    int32
	ISubItem int32
	IGroup   int32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/bb774771.aspx
type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    uintptr
	UKeyFlags uint32
}

type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey uint16
	Flags uint32
}

// 参考网址：https://msdn.microsoft.com/en-us/library/windows/desktop/bb774773.aspx
type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    uintptr
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/bb774780.aspx
// （由于您提供的Golang代码片段中没有具体的代码内容，这里仅对提供的URL链接进行了翻译，该链接为微软官方文档关于某个Windows桌面开发相关的技术文章。）
type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/bb775507.aspx
// （由于您提供的代码片段中并没有具体的代码，故这里仅对URL进行了翻译。）
// 该网址是微软开发者网络（MSDN）上关于Windows桌面应用程序开发的一个页面链接，内容可能与某个API函数、编程接口或技术指南相关，但没有具体代码所以无法给出更精确的注释翻译。
type INITCOMMONCONTROLSEX struct {
	DwSize uint32
	DwICC  uint32
}

// 参考文献：《Windows桌面应用程序开发》，MSDN官方文档，网址：http://msdn.microsoft.com/en-us/library/windows/desktop/bb760256.aspx
type TOOLINFO struct {
	CbSize     uint32
	UFlags     uint32
	Hwnd       HWND
	UId        uintptr
	Rect       RECT
	Hinst      HINSTANCE
	LpszText   *uint16
	LParam     uintptr
	LpReserved unsafe.Pointer
}

// 参考文献：[微软开发者网络]（https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms645604.aspx）
type TRACKMOUSEEVENT struct {
	CbSize      uint32
	DwFlags     uint32
	HwndTrack   HWND
	DwHoverTime uint32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms534067.aspx
type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms534068.aspx
type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd162768.aspx
// （由于您提供的Golang代码片段中没有具体的代码内容，这里仅对URL进行了翻译，原链接为微软官方文档关于某个Windows桌面开发相关的技术说明。）
type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/aa363646.aspx
// （由于您提供的代码片段中没有具体的代码，此处仅为该URL的中文解释：此链接是微软官方文档关于Windows桌面开发的一个页面，内容涉及某个API函数或功能的详细介绍和用法，具体主题需访问该链接查看。）
type EVENTLOGRECORD struct {
	Length              uint32
	Reserved            uint32
	RecordNumber        uint32
	TimeGenerated       uint32
	TimeWritten         uint32
	EventID             uint32
	EventType           uint16
	NumStrings          uint16
	EventCategory       uint16
	ReservedFlags       uint16
	ClosingRecordNumber uint32
	StringOffset        uint32
	UserSidLength       uint32
	UserSidOffset       uint32
	DataLength          uint32
	DataOffset          uint32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms685996.aspx
type SERVICE_STATUS struct {
	DwServiceType             uint32
	DwCurrentState            uint32
	DwControlsAccepted        uint32
	DwWin32ExitCode           uint32
	DwServiceSpecificExitCode uint32
	DwCheckPoint              uint32
	DwWaitHint                uint32
}

/* -------------------------
    Undocumented API
------------------------- */

type ACCENT_STATE DWORD

const (
	ACCENT_DISABLED                   ACCENT_STATE = 0
	ACCENT_ENABLE_GRADIENT            ACCENT_STATE = 1
	ACCENT_ENABLE_TRANSPARENTGRADIENT ACCENT_STATE = 2
	ACCENT_ENABLE_BLURBEHIND          ACCENT_STATE = 3
	ACCENT_ENABLE_ACRYLICBLURBEHIND   ACCENT_STATE = 4 // RS4 1803
	ACCENT_ENABLE_HOSTBACKDROP        ACCENT_STATE = 5 // RS5 1809
	ACCENT_INVALID_STATE              ACCENT_STATE = 6
)

type ACCENT_POLICY struct {
	AccentState   ACCENT_STATE
	AccentFlags   DWORD
	GradientColor DWORD
	AnimationId   DWORD
}

type WINDOWCOMPOSITIONATTRIBDATA struct {
	Attrib WINDOWCOMPOSITIONATTRIB
	PvData PVOID
	CbData SIZE_T
}

type WINDOWCOMPOSITIONATTRIB DWORD

const (
	WCA_UNDEFINED                     WINDOWCOMPOSITIONATTRIB = 0
	WCA_NCRENDERING_ENABLED           WINDOWCOMPOSITIONATTRIB = 1
	WCA_NCRENDERING_POLICY            WINDOWCOMPOSITIONATTRIB = 2
	WCA_TRANSITIONS_FORCEDISABLED     WINDOWCOMPOSITIONATTRIB = 3
	WCA_ALLOW_NCPAINT                 WINDOWCOMPOSITIONATTRIB = 4
	WCA_CAPTION_BUTTON_BOUNDS         WINDOWCOMPOSITIONATTRIB = 5
	WCA_NONCLIENT_RTL_LAYOUT          WINDOWCOMPOSITIONATTRIB = 6
	WCA_FORCE_ICONIC_REPRESENTATION   WINDOWCOMPOSITIONATTRIB = 7
	WCA_EXTENDED_FRAME_BOUNDS         WINDOWCOMPOSITIONATTRIB = 8
	WCA_HAS_ICONIC_BITMAP             WINDOWCOMPOSITIONATTRIB = 9
	WCA_THEME_ATTRIBUTES              WINDOWCOMPOSITIONATTRIB = 10
	WCA_NCRENDERING_EXILED            WINDOWCOMPOSITIONATTRIB = 11
	WCA_NCADORNMENTINFO               WINDOWCOMPOSITIONATTRIB = 12
	WCA_EXCLUDED_FROM_LIVEPREVIEW     WINDOWCOMPOSITIONATTRIB = 13
	WCA_VIDEO_OVERLAY_ACTIVE          WINDOWCOMPOSITIONATTRIB = 14
	WCA_FORCE_ACTIVEWINDOW_APPEARANCE WINDOWCOMPOSITIONATTRIB = 15
	WCA_DISALLOW_PEEK                 WINDOWCOMPOSITIONATTRIB = 16
	WCA_CLOAK                         WINDOWCOMPOSITIONATTRIB = 17
	WCA_CLOAKED                       WINDOWCOMPOSITIONATTRIB = 18
	WCA_ACCENT_POLICY                 WINDOWCOMPOSITIONATTRIB = 19
	WCA_FREEZE_REPRESENTATION         WINDOWCOMPOSITIONATTRIB = 20
	WCA_EVER_UNCLOAKED                WINDOWCOMPOSITIONATTRIB = 21
	WCA_VISUAL_OWNER                  WINDOWCOMPOSITIONATTRIB = 22
	WCA_HOLOGRAPHIC                   WINDOWCOMPOSITIONATTRIB = 23
	WCA_EXCLUDED_FROM_DDA             WINDOWCOMPOSITIONATTRIB = 24
	WCA_PASSIVEUPDATEMODE             WINDOWCOMPOSITIONATTRIB = 25
	WCA_USEDARKMODECOLORS             WINDOWCOMPOSITIONATTRIB = 26
	WCA_CORNER_STYLE                  WINDOWCOMPOSITIONATTRIB = 27
	WCA_PART_COLOR                    WINDOWCOMPOSITIONATTRIB = 28
	WCA_DISABLE_MOVESIZE_FEEDBACK     WINDOWCOMPOSITIONATTRIB = 29
	WCA_LAST                          WINDOWCOMPOSITIONATTRIB = 30
)

// ------------------------- 
// （这条注释未包含任何具体信息，只是一条分割线，用于在代码中进行视觉上的分隔或标记某个区域的开始或结束。）

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms684225.aspx
// （由于您提供的代码片段中没有具体的代码内容，故此处仅对URL进行了翻译，该链接为微软开发者网络（MSDN）上关于Windows桌面应用程序开发的一个页面。）
type MODULEENTRY32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  *uint8
	ModBaseSize  uint32
	HModule      HMODULE
	SzModule     [MAX_MODULE_NAME32 + 1]uint16
	SzExePath    [MAX_PATH]uint16
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms724284.aspx
type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

// 参考文档：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms682119.aspx
// （由于您提供的代码片段中没有具体的代码，故此处仅提供了参考文档的中文翻译。该链接为微软官方文档，描述了Windows桌面应用程序开发中的某个API函数或功能，但未给出具体是哪个函数，因此无法提供更精确的注释翻译。）
type COORD struct {
	X, Y int16
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms686311.aspx（该链接为微软官方文档，描述了某个Windows桌面API的接口或功能）
type SMALL_RECT struct {
	Left, Top, Right, Bottom int16
}

// 参考文献：《MSDN.microsoft.com》，Windows桌面应用程序开发，函数功能描述页面
// 链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms682093.aspx
type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/bb773244.aspx
// （该注释仅提供了一个参考链接，没有具体针对某段代码进行解释，故不涉及翻译具体注释内容。原始链接为微软开发者网络（MSDN）上关于某个Windows桌面开发相关主题的技术文档。）
type MARGINS struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

// 参考链接：https://msdn.microsoft.com/zh-cn/library/windows/desktop/aa969500.aspx
type DWM_BLURBEHIND struct {
	DwFlags                uint32
	fEnable                BOOL
	hRgnBlur               HRGN
	fTransitionOnMaximized BOOL
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/aa969501.aspx
// （由于您提供的代码片段中没有具体的代码，故此仅为参考网址的中文翻译说明）
// 该注释表示这是一个参考资料，指向的是MSDN（微软开发者网络）上关于某个Windows桌面开发相关主题的技术文档。文档地址为：http://msdn.microsoft.com/en-us/library/windows/desktop/aa969501.aspx ，具体内容需访问链接查看。
type DWM_PRESENT_PARAMETERS struct {
	cbSize             uint32
	fQueue             BOOL
	cRefreshStart      DWM_FRAME_COUNT
	cBuffer            uint32
	fUseSourceRate     BOOL
	rateSource         UNSIGNED_RATIO
	cRefreshesPerFrame uint32
	eSampling          DWM_SOURCE_FRAME_SAMPLING
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/aa969502.aspx
// （由于没有提供具体的代码行，此处仅为参考链接的中文翻译说明，该链接为微软官方文档关于Windows桌面开发的一个技术指南页面。）
type DWM_THUMBNAIL_PROPERTIES struct {
	dwFlags               uint32
	rcDestination         RECT
	rcSource              RECT
	opacity               byte
	fVisible              BOOL
	fSourceClientAreaOnly BOOL
}

// 参考文献：《Windows桌面应用程序开发》，MSDN官方文档，网址为：http://msdn.microsoft.com/en-us/library/windows/desktop/aa969503.aspx
type DWM_TIMING_INFO struct {
	cbSize                 uint32
	rateRefresh            UNSIGNED_RATIO
	qpcRefreshPeriod       QPC_TIME
	rateCompose            UNSIGNED_RATIO
	qpcVBlank              QPC_TIME
	cRefresh               DWM_FRAME_COUNT
	cDXRefresh             uint32
	qpcCompose             QPC_TIME
	cFrame                 DWM_FRAME_COUNT
	cDXPresent             uint32
	cRefreshFrame          DWM_FRAME_COUNT
	cFrameSubmitted        DWM_FRAME_COUNT
	cDXPresentSubmitted    uint32
	cFrameConfirmed        DWM_FRAME_COUNT
	cDXPresentConfirmed    uint32
	cRefreshConfirmed      DWM_FRAME_COUNT
	cDXRefreshConfirmed    uint32
	cFramesLate            DWM_FRAME_COUNT
	cFramesOutstanding     uint32
	cFrameDisplayed        DWM_FRAME_COUNT
	qpcFrameDisplayed      QPC_TIME
	cRefreshFrameDisplayed DWM_FRAME_COUNT
	cFrameComplete         DWM_FRAME_COUNT
	qpcFrameComplete       QPC_TIME
	cFramePending          DWM_FRAME_COUNT
	qpcFramePending        QPC_TIME
	cFramesDisplayed       DWM_FRAME_COUNT
	cFramesComplete        DWM_FRAME_COUNT
	cFramesPending         DWM_FRAME_COUNT
	cFramesAvailable       DWM_FRAME_COUNT
	cFramesDropped         DWM_FRAME_COUNT
	cFramesMissed          DWM_FRAME_COUNT
	cRefreshNextDisplayed  DWM_FRAME_COUNT
	cRefreshNextPresented  DWM_FRAME_COUNT
	cRefreshesDisplayed    DWM_FRAME_COUNT
	cRefreshesPresented    DWM_FRAME_COUNT
	cRefreshStarted        DWM_FRAME_COUNT
	cPixelsReceived        uint64
	cPixelsDrawn           uint64
	cBuffersEmpty          DWM_FRAME_COUNT
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd389402.aspx
// （由于您提供的代码片段中没有具体的代码内容，故此处仅对引用的MSDN文档链接进行了翻译，该链接指向的是一个关于Windows桌面开发的技术文档。）
type MilMatrix3x2D struct {
	S_11, S_12, S_21, S_22 float64
	DX, DY                 float64
}

// 参考文档：https://msdn.microsoft.com/zh-cn/library/windows/desktop/aa969505.aspx
type UNSIGNED_RATIO struct {
	uiNumerator   uint32
	uiDenominator uint32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms632603.aspx
// （由于您提供的代码片段中没有具体的代码，此处仅为参考链接的翻译说明。该链接是微软官方文档关于Windows桌面应用程序开发的一个页面，描述了某个API函数或功能的详细信息。）
type CREATESTRUCT struct {
	CreateParams uintptr
	Instance     HINSTANCE
	Menu         HMENU
	Parent       HWND
	Cy, Cx       int32
	Y, X         int32
	Style        int32
	Name         *uint16
	Class        *uint16
	dwExStyle    uint32
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145065.aspx
// （由于您提供的代码片段中没有具体的代码内容，此处仅翻译了注释部分，该注释表明这是一个参考链接，指向的是MSDN（微软开发者网络）上关于某个Windows桌面开发相关主题的技术文档。）
type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

type WINDOWINFO struct {
	CbSize          DWORD
	RcWindow        RECT
	RcClient        RECT
	DwStyle         DWORD
	DwExStyle       DWORD
	DwWindowStatus  DWORD
	CxWindowBorders UINT
	CyWindowBorders UINT
	AtomWindowType  ATOM
	WCreatorVersion WORD
}

type MONITOR_DPI_TYPE int32

const (
	MDT_EFFECTIVE_DPI MONITOR_DPI_TYPE = 0
	MDT_ANGULAR_DPI   MONITOR_DPI_TYPE = 1
	MDT_RAW_DPI       MONITOR_DPI_TYPE = 2
	MDT_DEFAULT       MONITOR_DPI_TYPE = 0
)

func (w *WINDOWINFO) isStyle(style DWORD) bool {
	return w.DwStyle&style == style
}

func (w *WINDOWINFO) IsPopup() bool {
	return w.isStyle(WS_POPUP)
}

func (m *MONITORINFO) Dump() {
	fmt.Printf("MONITORINFO (%p)\n", m)
	fmt.Printf("  CbSize   : %d\n", m.CbSize)
	fmt.Printf("  RcMonitor: %s\n", &m.RcMonitor)
	fmt.Printf("  RcWork   : %s\n", &m.RcWork)
	fmt.Printf("  DwFlags  : %d\n", m.DwFlags)
}

// 参考文档：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145066.aspx
// （由于没有提供具体的代码行，此处仅翻译了URL链接的注释，该链接指向的是微软官方关于某个Windows桌面开发相关的技术文档。）
type MONITORINFOEX struct {
	MONITORINFO
	SzDevice [CCHDEVICENAME]uint16
}

// 参考文档：https://msdn.microsoft.com/zh-cn/library/windows/desktop/dd368826.aspx
type PIXELFORMATDESCRIPTOR struct {
	Size                   uint16
	Version                uint16
	DwFlags                uint32
	IPixelType             byte
	ColorBits              byte
	RedBits, RedShift      byte
	GreenBits, GreenShift  byte
	BlueBits, BlueShift    byte
	AlphaBits, AlphaShift  byte
	AccumBits              byte
	AccumRedBits           byte
	AccumGreenBits         byte
	AccumBlueBits          byte
	AccumAlphaBits         byte
	DepthBits, StencilBits byte
	AuxBuffers             byte
	ILayerType             byte
	Reserved               byte
	DwLayerMask            uint32
	DwVisibleMask          uint32
	DwDamageMask           uint32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms646270(v=vs.85).aspx
// （该注释未对具体代码进行解释，仅提供了一个MSDN（微软开发者网络）的官方开发文档链接，该文档详细介绍了Windows桌面应用程序编程接口（API）中的某个函数或方法的功能和用法。）
type INPUT struct {
	Type uint32
	Mi   MOUSEINPUT
	Ki   KEYBDINPUT
	Hi   HARDWAREINPUT
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms646273(v=vs.85).aspx
// （由于没有提供具体的代码行，此处仅为引用链接的中文解释）
// 该链接指向微软MSDN（Microsoft Developer Network）上关于Windows桌面API的一个文档页面，
// 文档标题为“GetThreadDescription 函数”，内容涵盖了该函数在Windows操作系统中的使用方法和详细说明，
// 针对Visual Studio 8.5版本，描述了如何获取线程的描述信息。
type MOUSEINPUT struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// 参考微软官方文档：https://msdn.microsoft.com/zh-cn/library/windows/desktop/ms646271(v=vs.85).aspx
type KEYBDINPUT struct {
	WVk         uint16
	WScan       uint16
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms646269(v=vs.85).aspx
// （由于没有提供具体的代码行，此注释仅提供了参考的MSDN文档链接，该链接为Windows桌面API函数的官方文档，描述了某个函数的功能、参数和使用方法。）
type HARDWAREINPUT struct {
	UMsg    uint32
	WParamL uint16
	WParamH uint16
}

type KbdInput struct {
	typ uint32
	ki  KEYBDINPUT
}

type MouseInput struct {
	typ uint32
	mi  MOUSEINPUT
}

type HardwareInput struct {
	typ uint32
	hi  HARDWAREINPUT
}

// 参考微软官方MSDN文档：https://msdn.microsoft.com/en-us/library/windows/desktop/ms724950(v=vs.85).aspx
// （由于没有提供具体的代码行，此注释是对链接的描述，指向的是Windows桌面开发相关的一个MSDN页面资源，该页面内容与Windows API的某个功能或接口有关。）
type SYSTEMTIME struct {
	Year         uint16
	Month        uint16
	DayOfWeek    uint16
	Day          uint16
	Hour         uint16
	Minute       uint16
	Second       uint16
	Milliseconds uint16
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/ms644967(v=vs.85).aspx
// 该注释并未给出具体代码，它提供了一个 MSDN（微软开发者网络）上的官方文档链接，内容是关于Windows桌面应用程序开发的一个API函数的详细说明。翻译后的注释如下：
// ```go
// 链接指向的是MSDN上关于Windows桌面版开发的一个文档资源，
// 文档地址为：http://msdn.microsoft.com/en-us/library/windows/desktop/ms644967(v=vs.85).aspx
type KBDLLHOOKSTRUCT struct {
	VkCode      DWORD
	ScanCode    DWORD
	Flags       DWORD
	Time        DWORD
	DwExtraInfo ULONG_PTR
}

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    POINT
	PtMaxPosition    POINT
	RcNormalPosition RECT
}

type SCROLLINFO struct {
	CbSize    uint32
	FMask     uint32
	NMin      int32
	NMax      int32
	NPage     uint32
	NPos      int32
	NTrackPos int32
}
