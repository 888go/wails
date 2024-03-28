//go:build obfuscated

package app

// IsObfuscated 返回 true，如果设置了混淆构建标签

// ff:
func IsObfuscated() bool {
	return true
}
