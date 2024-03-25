//go:build windows
// +build windows

package webview2runtime

import (
	_ "embed"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

// Info 包含关于 webview2 运行时安装的所有信息。
type Info struct {
	Location        string
	Name            string
	Version         string
	SilentUninstall string
}

// IsOlderThan 函数在已安装版本比给定所需版本旧时返回 true。
// 如果出现错误，将返回错误信息。

// ff:
// requiredVersion:
func (i *Info) IsOlderThan(requiredVersion string) (bool, error) {
	var mod = syscall.NewLazyDLL("WebView2Loader.dll")
	var CompareBrowserVersions = mod.NewProc("CompareBrowserVersions")
	v1, err := syscall.UTF16PtrFromString(i.Version)
	if err != nil {
		return false, err
	}
	v2, err := syscall.UTF16PtrFromString(requiredVersion)
	if err != nil {
		return false, err
	}
	var result int = 9
	_, _, err = CompareBrowserVersions.Call(uintptr(unsafe.Pointer(v1)), uintptr(unsafe.Pointer(v2)), uintptr(unsafe.Pointer(&result)))
	if result < -1 || result > 1 {
		return false, err
	}
	return result == -1, nil
}

func downloadBootstrapper() (string, error) {
	bootstrapperURL := `https://go.microsoft.com/fwlink/p/?LinkId=2124703`
	installer := filepath.Join(os.TempDir(), `MicrosoftEdgeWebview2Setup.exe`)

	// Download installer
	out, err := os.Create(installer)
	defer out.Close()
	if err != nil {
		return "", err
	}
	resp, err := http.Get(bootstrapperURL)
	defer resp.Body.Close()
	if err != nil {
		err = out.Close()
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return installer, nil
}

// InstallUsingEmbeddedBootstrapper 会从微软下载引导程序并运行它，以安装
// 运行时的最新版本。
// 如果安装程序运行成功，则返回 true。
// 如果发生错误，则返回错误。

// ff:
func InstallUsingEmbeddedBootstrapper() (bool, error) {
	installer, err := WriteInstaller(os.TempDir())
	if err != nil {
		return false, err
	}
	result, err := runInstaller(installer)
	if err != nil {
		return false, err
	}

	return result, os.Remove(installer)

}

// InstallUsingBootstrapper 会从微软提取嵌入式的引导程序并运行它，以安装
// 运行时的最新版本。
// 如果安装程序成功运行，则返回 true。
// 如果出现错误，则返回错误。

// ff:
func InstallUsingBootstrapper() (bool, error) {

	installer, err := downloadBootstrapper()
	if err != nil {
		return false, err
	}

	result, err := runInstaller(installer)
	if err != nil {
		return false, err
	}

	return result, os.Remove(installer)

}

func runInstaller(installer string) (bool, error) {
	// 代码来源：https://stackoverflow.com/a/10385867
// （该注释表明了以下Golang代码片段来源于Stack Overflow网站上的一个回答，链接为https://stackoverflow.com/a/10385867）
	cmd := exec.Command(installer)
	if err := cmd.Start(); err != nil {
		return false, err
	}
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus() == 0, nil
			}
		}
	}
	return true, nil
}

// Confirm 将向用户显示一条消息以及“确定”和“取消”按钮。
// 如果用户选择了“确定”，则返回 true。
// 如果出现错误，则返回错误。

// ff:
// title:
// caption:
func Confirm(caption string, title string) (bool, error) {
	var flags uint = 0x00000001 // MB_OKCANCEL
	result, err := MessageBox(caption, title, flags)
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// Error 将错误消息发送给用户。
// 如果发生错误，将返回一个错误。

// ff:
// title:
// caption:
func Error(caption string, title string) error {
	var flags uint = 0x00000010 // MB_ICONERROR
	_, err := MessageBox(caption, title, flags)
	return err
}

// MessageBox 函数以给定的标题和提示信息向用户显示对话框。
// 可以通过参数 Flags 自定义对话框样式。
// 如果发生错误，将返回一个错误。

// ff:
// flags:
// title:
// caption:
func MessageBox(caption string, title string, flags uint) (int, error) {
	captionUTF16, err := syscall.UTF16PtrFromString(caption)
	if err != nil {
		return -1, err
	}
	titleUTF16, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return -1, err
	}
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(0),
		uintptr(unsafe.Pointer(captionUTF16)),
		uintptr(unsafe.Pointer(titleUTF16)),
		uintptr(flags))

	return int(ret), nil
}

// OpenInstallerDownloadWebpage 将会在 WebView2 下载页面上打开浏览器

// ff:
func OpenInstallerDownloadWebpage() error {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://developer.microsoft.com/en-us/microsoft-edge/webview2/")
	return cmd.Run()
}
