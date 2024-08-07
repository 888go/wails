//go:build linux
// +build linux

package linux

/*
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0

#include "gtk/gtk.h"

static GtkCheckMenuItem *toGtkCheckMenuItem(void *pointer) { return (GTK_CHECK_MENU_ITEM(pointer)); }

extern void blockClick(GtkWidget* menuItem, gulong handler_id);
extern void unblockClick(GtkWidget* menuItem, gulong handler_id);
*/
import "C"
import (
	"unsafe"

	"github.com/888go/wails/pkg/menu"
)


// ff:
// label:
func GtkMenuItemWithLabel(label string) *C.GtkWidget {
	cLabel := C.CString(label)
	result := C.gtk_menu_item_new_with_label(cLabel)
	C.free(unsafe.Pointer(cLabel))
	return result
}


// ff:
// label:
func GtkCheckMenuItemWithLabel(label string) *C.GtkWidget {
	cLabel := C.CString(label)
	result := C.gtk_check_menu_item_new_with_label(cLabel)
	C.free(unsafe.Pointer(cLabel))
	return result
}


// ff:
// group:
// label:
func GtkRadioMenuItemWithLabel(label string, group *C.GSList) *C.GtkWidget {
	cLabel := C.CString(label)
	result := C.gtk_radio_menu_item_new_with_label(group, cLabel)
	C.free(unsafe.Pointer(cLabel))
	return result
}

//export handleMenuItemClick
// 导出handleMenuItemClick函数，供其他语言（如C）调用
func handleMenuItemClick(gtkWidget unsafe.Pointer) {
// 确保在新的goroutine上执行最终的回调函数，否则如果回调函数（例如）尝试打开一个对话框，主线程将会被阻塞，因此消息循环也会阻塞。其结果是应用会被阻塞并显示一个“无响应”的对话框。

	item := gtkSignalToMenuItem[(*C.GtkWidget)(gtkWidget)]
	switch item.Type {
	case menu.CheckboxType:
		item.Checked = !item.Checked
		checked := C.int(0)
		if item.Checked {
			checked = C.int(1)
		}
		for _, gtkCheckbox := range gtkCheckboxCache[item] {
			handler := gtkSignalHandlers[gtkCheckbox]
			C.blockClick(gtkCheckbox, handler)
			C.gtk_check_menu_item_set_active(C.toGtkCheckMenuItem(unsafe.Pointer(gtkCheckbox)), checked)
			C.unblockClick(gtkCheckbox, handler)
		}
		go item.Click(&menu.CallbackData{MenuItem: item})
	case menu.RadioType:
		gtkRadioItems := gtkRadioMenuCache[item]
		active := C.gtk_check_menu_item_get_active(C.toGtkCheckMenuItem(gtkWidget))
		if int(active) == 1 {
			for _, gtkRadioItem := range gtkRadioItems {
				handler := gtkSignalHandlers[gtkRadioItem]
				C.blockClick(gtkRadioItem, handler)
				C.gtk_check_menu_item_set_active(C.toGtkCheckMenuItem(unsafe.Pointer(gtkRadioItem)), 1)
				C.unblockClick(gtkRadioItem, handler)
			}
			item.Checked = true
			go item.Click(&menu.CallbackData{MenuItem: item})
		} else {
			item.Checked = false
		}
	default:
		go item.Click(&menu.CallbackData{MenuItem: item})
	}
}
