//go:build windows

/*
 * Copyright (C) 2019 Tad Vizbaras. All Rights Reserved.
 * Copyright (C) 2010-2012 The W32 Authors. All Rights Reserved.
 */
package w32

type pIUnknownVtbl struct {
	pQueryInterface uintptr
	pAddRef         uintptr
	pRelease        uintptr
}

type IUnknown struct {
	lpVtbl *pIUnknownVtbl
}


// ff:
// id:
func (this *IUnknown) QueryInterface(id *GUID) *IDispatch {
	return ComQueryInterface(this, id)
}


// ff:
func (this *IUnknown) AddRef() int32 {
	return ComAddRef(this)
}


// ff:
func (this *IUnknown) Release() int32 {
	return ComRelease(this)
}
