//go:build !devtools

package app

// IsDevtoolsEnabled 返回布尔值，若应启用开发者工具则返回true
// 注意：在调试版本中也会添加devtools标志
func IsDevtoolsEnabled() bool {
	return false
}
