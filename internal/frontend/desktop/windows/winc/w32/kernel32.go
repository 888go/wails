//go:build windows

/*
 * Copyright (C) 2019 Tad Vizbaras. All Rights Reserved.
 * Copyright (C) 2010-2012 The W32 Authors. All Rights Reserved.
 */
package w32

import (
	"syscall"
	"unsafe"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetModuleHandle            = modkernel32.NewProc("GetModuleHandleW")
	procMulDiv                     = modkernel32.NewProc("MulDiv")
	procGetConsoleWindow           = modkernel32.NewProc("GetConsoleWindow")
	procGetCurrentThread           = modkernel32.NewProc("GetCurrentThread")
	procGetCurrentThreadId         = modkernel32.NewProc("GetCurrentThreadId")
	procGetLogicalDrives           = modkernel32.NewProc("GetLogicalDrives")
	procGetLogicalDriveStrings     = modkernel32.NewProc("GetLogicalDriveStringsW")
	procGetUserDefaultLCID         = modkernel32.NewProc("GetUserDefaultLCID")
	procLstrlen                    = modkernel32.NewProc("lstrlenW")
	procLstrcpy                    = modkernel32.NewProc("lstrcpyW")
	procGlobalAlloc                = modkernel32.NewProc("GlobalAlloc")
	procGlobalFree                 = modkernel32.NewProc("GlobalFree")
	procGlobalLock                 = modkernel32.NewProc("GlobalLock")
	procGlobalUnlock               = modkernel32.NewProc("GlobalUnlock")
	procMoveMemory                 = modkernel32.NewProc("RtlMoveMemory")
	procFindResource               = modkernel32.NewProc("FindResourceW")
	procSizeofResource             = modkernel32.NewProc("SizeofResource")
	procLockResource               = modkernel32.NewProc("LockResource")
	procLoadResource               = modkernel32.NewProc("LoadResource")
	procGetLastError               = modkernel32.NewProc("GetLastError")
	procOpenProcess                = modkernel32.NewProc("OpenProcess")
	procTerminateProcess           = modkernel32.NewProc("TerminateProcess")
	procCloseHandle                = modkernel32.NewProc("CloseHandle")
	procCreateToolhelp32Snapshot   = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procModule32First              = modkernel32.NewProc("Module32FirstW")
	procModule32Next               = modkernel32.NewProc("Module32NextW")
	procGetSystemTimes             = modkernel32.NewProc("GetSystemTimes")
	procGetConsoleScreenBufferInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleTextAttribute    = modkernel32.NewProc("SetConsoleTextAttribute")
	procGetDiskFreeSpaceEx         = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procGetProcessTimes            = modkernel32.NewProc("GetProcessTimes")
	procSetSystemTime              = modkernel32.NewProc("SetSystemTime")
	procGetSystemTime              = modkernel32.NewProc("GetSystemTime")
)


// ff:
// modulename:
func GetModuleHandle(modulename string) HINSTANCE {
	var mn uintptr
	if modulename == "" {
		mn = 0
	} else {
		mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
	}
	ret, _, _ := procGetModuleHandle.Call(mn)
	return HINSTANCE(ret)
}


// ff:
// denominator:
// numerator:
// number:
func MulDiv(number, numerator, denominator int) int {
	ret, _, _ := procMulDiv.Call(
		uintptr(number),
		uintptr(numerator),
		uintptr(denominator))

	return int(ret)
}


// ff:
func GetConsoleWindow() HWND {
	ret, _, _ := procGetConsoleWindow.Call()

	return HWND(ret)
}


// ff:
func GetCurrentThread() HANDLE {
	ret, _, _ := procGetCurrentThread.Call()

	return HANDLE(ret)
}


// ff:
func GetCurrentThreadId() HANDLE {
	ret, _, _ := procGetCurrentThreadId.Call()

	return HANDLE(ret)
}


// ff:
func GetLogicalDrives() uint32 {
	ret, _, _ := procGetLogicalDrives.Call()

	return uint32(ret)
}


// ff:
func GetUserDefaultLCID() uint32 {
	ret, _, _ := procGetUserDefaultLCID.Call()

	return uint32(ret)
}


// ff:
// lpString:
func Lstrlen(lpString *uint16) int {
	ret, _, _ := procLstrlen.Call(uintptr(unsafe.Pointer(lpString)))

	return int(ret)
}


// ff:
// lpString:
// buf:
func Lstrcpy(buf []uint16, lpString *uint16) {
	procLstrcpy.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(lpString)))
}


// ff:
// dwBytes:
// uFlags:
func GlobalAlloc(uFlags uint, dwBytes uint32) HGLOBAL {
	ret, _, _ := procGlobalAlloc.Call(
		uintptr(uFlags),
		uintptr(dwBytes))

	if ret == 0 {
		panic("GlobalAlloc failed")
	}

	return HGLOBAL(ret)
}


// ff:
// hMem:
func GlobalFree(hMem HGLOBAL) {
	ret, _, _ := procGlobalFree.Call(uintptr(hMem))

	if ret != 0 {
		panic("GlobalFree failed")
	}
}


// ff:
// hMem:
func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := procGlobalLock.Call(uintptr(hMem))

	if ret == 0 {
		panic("GlobalLock failed")
	}

	return unsafe.Pointer(ret)
}


// ff:
// hMem:
func GlobalUnlock(hMem HGLOBAL) bool {
	ret, _, _ := procGlobalUnlock.Call(uintptr(hMem))

	return ret != 0
}


// ff:
// length:
// source:
// destination:
func MoveMemory(destination, source unsafe.Pointer, length uint32) {
	procMoveMemory.Call(
		uintptr(unsafe.Pointer(destination)),
		uintptr(source),
		uintptr(length))
}


// ff:
// HRSRC:
// lpType:
// lpName:
// hModule:
func FindResource(hModule HMODULE, lpName, lpType *uint16) (HRSRC, error) {
	ret, _, _ := procFindResource.Call(
		uintptr(hModule),
		uintptr(unsafe.Pointer(lpName)),
		uintptr(unsafe.Pointer(lpType)))

	if ret == 0 {
		return 0, syscall.GetLastError()
	}

	return HRSRC(ret), nil
}


// ff:
// hResInfo:
// hModule:
func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
	ret, _, _ := procSizeofResource.Call(
		uintptr(hModule),
		uintptr(hResInfo))

	if ret == 0 {
		panic("SizeofResource failed")
	}

	return uint32(ret)
}


// ff:
// hResData:
func LockResource(hResData HGLOBAL) unsafe.Pointer {
	ret, _, _ := procLockResource.Call(uintptr(hResData))

	if ret == 0 {
		panic("LockResource failed")
	}

	return unsafe.Pointer(ret)
}


// ff:
// hResInfo:
// hModule:
func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
	ret, _, _ := procLoadResource.Call(
		uintptr(hModule),
		uintptr(hResInfo))

	if ret == 0 {
		panic("LoadResource failed")
	}

	return HGLOBAL(ret)
}


// ff:
func GetLastError() uint32 {
	ret, _, _ := procGetLastError.Call()
	return uint32(ret)
}


// ff:
// processId:
// inheritHandle:
// desiredAccess:
func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) HANDLE {
	inherit := 0
	if inheritHandle {
		inherit = 1
	}

	ret, _, _ := procOpenProcess.Call(
		uintptr(desiredAccess),
		uintptr(inherit),
		uintptr(processId))
	return HANDLE(ret)
}


// ff:
// uExitCode:
// hProcess:
func TerminateProcess(hProcess HANDLE, uExitCode uint) bool {
	ret, _, _ := procTerminateProcess.Call(
		uintptr(hProcess),
		uintptr(uExitCode))
	return ret != 0
}


// ff:
// object:
func CloseHandle(object HANDLE) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(object))
	return ret != 0
}


// ff:
// processId:
// flags:
func CreateToolhelp32Snapshot(flags, processId uint32) HANDLE {
	ret, _, _ := procCreateToolhelp32Snapshot.Call(
		uintptr(flags),
		uintptr(processId))

	if ret <= 0 {
		return HANDLE(0)
	}

	return HANDLE(ret)
}


// ff:
// me:
// snapshot:
func Module32First(snapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32First.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(me)))

	return ret != 0
}


// ff:
// me:
// snapshot:
func Module32Next(snapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32Next.Call(
		uintptr(snapshot),
		uintptr(unsafe.Pointer(me)))

	return ret != 0
}


// ff:
// lpUserTime:
// lpKernelTime:
// lpIdleTime:
func GetSystemTimes(lpIdleTime, lpKernelTime, lpUserTime *FILETIME) bool {
	ret, _, _ := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(lpIdleTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))

	return ret != 0
}


// ff:
// lpUserTime:
// lpKernelTime:
// lpExitTime:
// lpCreationTime:
// hProcess:
func GetProcessTimes(hProcess HANDLE, lpCreationTime, lpExitTime, lpKernelTime, lpUserTime *FILETIME) bool {
	ret, _, _ := procGetProcessTimes.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpCreationTime)),
		uintptr(unsafe.Pointer(lpExitTime)),
		uintptr(unsafe.Pointer(lpKernelTime)),
		uintptr(unsafe.Pointer(lpUserTime)))

	return ret != 0
}


// ff:
// hConsoleOutput:
func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE) *CONSOLE_SCREEN_BUFFER_INFO {
	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}


// ff:
// wAttributes:
// hConsoleOutput:
func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}


// ff:
// dirName:
func GetDiskFreeSpaceEx(dirName string) (r bool,
	freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64) {
	ret, _, _ := procGetDiskFreeSpaceEx.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dirName))),
		uintptr(unsafe.Pointer(&freeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalNumberOfBytes)),
		uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)))
	return ret != 0,
		freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes
}


// ff:
func GetSystemTime() *SYSTEMTIME {
	var time SYSTEMTIME
	procGetSystemTime.Call(
		uintptr(unsafe.Pointer(&time)))
	return &time
}


// ff:
// time:
func SetSystemTime(time *SYSTEMTIME) bool {
	ret, _, _ := procSetSystemTime.Call(
		uintptr(unsafe.Pointer(time)))
	return ret != 0
}


// ff:
// lpBuffer:
// nBufferLength:
func GetLogicalDriveStrings(nBufferLength uint32, lpBuffer *uint16) uint32 {
	ret, _, _ := procGetLogicalDriveStrings.Call(
		uintptr(nBufferLength),
		uintptr(unsafe.Pointer(lpBuffer)),
		0)

	return uint32(ret)
}
