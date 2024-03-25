//go:build windows

package windows

import (
	"testing"

	"github.com/888go/wails/internal/frontend"
	"golang.org/x/sys/windows"
)

func Test_calculateMessageDialogFlags(t *testing.T) {
	tests := []struct {
		name    string
		options frontend.MessageDialogOptions
		want    uint32
	}{
		{
			name: "Test Info Dialog",
			options: frontend.MessageDialogOptions{
				X对话框类型: frontend.X常量_对话框_信息,
			},
			want: windows.MB_OK | windows.MB_ICONINFORMATION,
		},
		{
			name: "Test Error Dialog",
			options: frontend.MessageDialogOptions{
				X对话框类型: frontend.X常量_对话框_错误,
			},
			want: windows.MB_ICONERROR | windows.MB_OK,
		},
		{
			name: "Test Question Dialog",
			options: frontend.MessageDialogOptions{
				X对话框类型: frontend.X常量_对话框_问题,
			},
			want: windows.MB_YESNO,
		},
		{
			name: "Test Question Dialog with default cancel",
			options: frontend.MessageDialogOptions{
				X对话框类型:          frontend.X常量_对话框_问题,
				X默认按钮: "No",
			},
			want: windows.MB_YESNO | windows.MB_DEFBUTTON2,
		},
		{
			name: "Test Question Dialog with default cancel (lowercase)",
			options: frontend.MessageDialogOptions{
				X对话框类型:          frontend.X常量_对话框_问题,
				X默认按钮: "no",
			},
			want: windows.MB_YESNO | windows.MB_DEFBUTTON2,
		},
		{
			name: "Test Warning Dialog",
			options: frontend.MessageDialogOptions{
				X对话框类型: frontend.X常量_对话框_警告,
			},
			want: windows.MB_OK | windows.MB_ICONWARNING,
		},
		{
			name: "Test Error Dialog",
			options: frontend.MessageDialogOptions{
				X对话框类型: frontend.X常量_对话框_错误,
			},
			want: windows.MB_ICONERROR | windows.MB_OK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMessageDialogFlags(tt.options); got != tt.want {
				t.Errorf("calculateMessageDialogFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
