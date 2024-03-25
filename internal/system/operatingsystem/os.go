package operatingsystem

// OS 包含关于操作系统的相关信息
type OS struct {
	ID      string
	Name    string
	Version string
}

// Info 获取当前平台的信息
func Info() (*OS, error) {
	return platformInfo()
}
