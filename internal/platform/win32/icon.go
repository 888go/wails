//go:build windows

package win32

import (
	"unsafe"
)

func CreateIconFromResourceEx(presbits uintptr, dwResSize uint32, isIcon bool, version uint32, cxDesired int, cyDesired int, flags uint) (uintptr, error) {
	icon := 0
	if isIcon {
		icon = 1
	}
	r, _, err := procCreateIconFromResourceEx.Call(
		presbits,
		uintptr(dwResSize),
		uintptr(icon),
		uintptr(version),
		uintptr(cxDesired),
		uintptr(cyDesired),
		uintptr(flags),
	)

	if r == 0 {
		return 0, err
	}
	return r, nil
}

// CreateHIconFromPNG 从PNG文件创建一个HICON
func CreateHIconFromPNG(pngData []byte) (HICON, error) {
	icon, err := CreateIconFromResourceEx(
		uintptr(unsafe.Pointer(&pngData[0])),
		uint32(len(pngData)),
		true,
		0x00030000,
		0,
		0,
		LR_DEFAULTSIZE)
	return HICON(icon), err
}
