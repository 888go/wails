//go:build windows

/*
 * Based on code originally from https://github.com/atotto/clipboard. Copyright (c) 2013 Ato Araki. All rights reserved.
 */

package win32

import (
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

const (
	cfUnicodetext = 13
	gmemMoveable  = 0x0002
)

// waitOpenClipboard 尝试打开剪贴板，最多等待1秒钟以完成此操作。
func waitOpenClipboard() error {
	started := time.Now()
	limit := started.Add(time.Second)
	var r uintptr
	var err error
	for time.Now().Before(limit) {
		r, _, err = procOpenClipboard.Call(0)
		if r != 0 {
			return nil
		}
		time.Sleep(time.Millisecond)
	}
	return err
}

func GetClipboardText() (string, error) {
// LockOSThread 确保整个方法从开始到结束都在同一个线程上执行（实际上是锁定了 goroutine 的线程归属）。
// 否则，如果 goroutine 在执行过程中切换了线程（这是常见情况），OpenClipboard 和 CloseClipboard 将在两个不同的线程上发生，这将导致剪贴板死锁。
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	if formatAvailable, _, err := procIsClipboardFormatAvailable.Call(cfUnicodetext); formatAvailable == 0 {
		return "", err
	}
	err := waitOpenClipboard()
	if err != nil {
		return "", err
	}

	h, _, err := procGetClipboardData.Call(cfUnicodetext)
	if h == 0 {
		_, _, _ = procCloseClipboard.Call()
		return "", err
	}

	l, _, err := kernelGlobalLock.Call(h)
	if l == 0 {
		_, _, _ = procCloseClipboard.Call()
		return "", err
	}

	text := syscall.UTF16ToString((*[1 << 20]uint16)(unsafe.Pointer(l))[:])

	r, _, err := kernelGlobalUnlock.Call(h)
	if r == 0 {
		_, _, _ = procCloseClipboard.Call()
		return "", err
	}

	closed, _, err := procCloseClipboard.Call()
	if closed == 0 {
		return "", err
	}
	return text, nil
}

func SetClipboardText(text string) error {
// LockOSThread 确保整个方法从开始到结束都在同一个线程上执行（实际上是锁定了 goroutine 的线程归属）。
// 否则，如果 goroutine 在执行过程中切换了线程（这是常见情况），OpenClipboard 和 CloseClipboard 将在两个不同的线程上发生，这将导致剪贴板死锁。
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := waitOpenClipboard()
	if err != nil {
		return err
	}

	r, _, err := procEmptyClipboard.Call(0)
	if r == 0 {
		_, _, _ = procCloseClipboard.Call()
		return err
	}

	data, err := syscall.UTF16FromString(text)
	if err != nil {
		return err
	}

// "如果hMem参数标识了一个内存对象，则该对象必须使用带有GMEM_MOVEABLE标志的函数分配。"
	h, _, err := kernelGlobalAlloc.Call(gmemMoveable, uintptr(len(data)*int(unsafe.Sizeof(data[0]))))
	if h == 0 {
		_, _, _ = procCloseClipboard.Call()
		return err
	}
	defer func() {
		if h != 0 {
			kernelGlobalFree.Call(h)
		}
	}()

	l, _, err := kernelGlobalLock.Call(h)
	if l == 0 {
		_, _, _ = procCloseClipboard.Call()
		return err
	}

	r, _, err = kernelLstrcpy.Call(l, uintptr(unsafe.Pointer(&data[0])))
	if r == 0 {
		_, _, _ = procCloseClipboard.Call()
		return err
	}

	r, _, err = kernelGlobalUnlock.Call(h)
	if r == 0 {
		if err.(syscall.Errno) != 0 {
			_, _, _ = procCloseClipboard.Call()
			return err
		}
	}

	r, _, err = procSetClipboardData.Call(cfUnicodetext, h)
	if r == 0 {
		_, _, _ = procCloseClipboard.Call()
		return err
	}
	h = 0 // 抑制延迟清理
	closed, _, err := procCloseClipboard.Call()
	if closed == 0 {
		return err
	}
	return nil
}
