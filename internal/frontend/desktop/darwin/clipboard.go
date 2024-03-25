//go:build darwin

package darwin

import (
	"os/exec"
)

// ff:剪贴板取文本
func (f *Frontend) ClipboardGetText() (string, error) {
	pasteCmd := exec.Command("pbpaste")
	out, err := pasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// ff:剪贴板置文本
// text:文本
func (f *Frontend) ClipboardSetText(text string) error {
	copyCmd := exec.Command("pbcopy")
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}
