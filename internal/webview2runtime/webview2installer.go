package webview2runtime

import (
	_ "embed"
	"os"
	"path/filepath"
)

//go:embed MicrosoftEdgeWebview2Setup.exe
var setupexe []byte

// WriteInstallerToFile 将安装程序文件写入给定的文件。
func WriteInstallerToFile(targetFile string) error {
	return os.WriteFile(targetFile, setupexe, 0o755)
}

// WriteInstaller 将安装程序 exe 文件写入给定的目录，并返回该文件的路径。
func WriteInstaller(targetPath string) (string, error) {
	installer := filepath.Join(targetPath, `MicrosoftEdgeWebview2Setup.exe`)
	return installer, WriteInstallerToFile(installer)
}
