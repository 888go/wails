//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

var DefaultBackgroundBrush = NewSystemColorBrush(w32.COLOR_BTNFACE)

type Brush struct {
	hBrush   w32.HBRUSH
	logBrush w32.LOGBRUSH
}


// ff:
// color:
func NewSolidColorBrush(color Color) *Brush {
	lb := w32.LOGBRUSH{LbStyle: w32.BS_SOLID, LbColor: w32.COLORREF(color)}
	hBrush := w32.CreateBrushIndirect(&lb)
	if hBrush == 0 {
		panic("Faild to create solid color brush")
	}

	return &Brush{hBrush, lb}
}


// ff:
// colorIndex:
func NewSystemColorBrush(colorIndex int) *Brush {
	// lb := w32.LOGBRUSH{LbStyle: w32.BS_SOLID, LbColor: w32.COLORREF(colorIndex)}
// 创建一个w32.LOGBRUSH结构体变量lb，其中：
// LbStyle字段设置为w32.BS_SOLID，表示画刷样式为实心填充；
// LbColor字段设置为w32.COLORREF类型的colorIndex，用于指定画刷颜色。
	lb := w32.LOGBRUSH{LbStyle: w32.BS_NULL}
	hBrush := w32.GetSysColorBrush(colorIndex)
	if hBrush == 0 {
		panic("GetSysColorBrush failed")
	}
	return &Brush{hBrush, lb}
}


// ff:
// color:
func NewHatchedColorBrush(color Color) *Brush {
	lb := w32.LOGBRUSH{LbStyle: w32.BS_HATCHED, LbColor: w32.COLORREF(color)}
	hBrush := w32.CreateBrushIndirect(&lb)
	if hBrush == 0 {
		panic("Faild to create solid color brush")
	}

	return &Brush{hBrush, lb}
}


// ff:
func NewNullBrush() *Brush {
	lb := w32.LOGBRUSH{LbStyle: w32.BS_NULL}
	hBrush := w32.CreateBrushIndirect(&lb)
	if hBrush == 0 {
		panic("Failed to create null brush")
	}

	return &Brush{hBrush, lb}
}


// ff:
func (br *Brush) GetHBRUSH() w32.HBRUSH {
	return br.hBrush
}


// ff:
func (br *Brush) GetLOGBRUSH() *w32.LOGBRUSH {
	return &br.logBrush
}


// ff:
func (br *Brush) Dispose() {
	if br.hBrush != 0 {
		w32.DeleteObject(w32.HGDIOBJ(br.hBrush))
		br.hBrush = 0
	}
}
