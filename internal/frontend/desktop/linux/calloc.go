//go:build linux
// +build linux

package linux

/*
#include <stdlib.h>
*/
import "C"
import "unsafe"

// Calloc 处理 C 数据的分配/释放
type Calloc struct {
	pool []unsafe.Pointer
}

// NewCalloc 创建一个新的分配器
func NewCalloc() Calloc {
	return Calloc{}
}

// String 创建一个新的C风格字符串，并保留对该字符串的引用
func (c Calloc) String(in string) *C.char {
	result := C.CString(in)
	c.pool = append(c.pool, unsafe.Pointer(result))
	return result
}

// Free 释放所有已分配的C语言内存
func (c Calloc) Free() {
	for _, str := range c.pool {
		C.free(str)
	}
	c.pool = []unsafe.Pointer{}
}
