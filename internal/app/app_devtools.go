//go:build devtools

package app

// 注意：在调试版本构建中也会添加devtools标志
func IsDevtoolsEnabled() bool {
	return true
}
