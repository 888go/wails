//go:build !obfuscated

package app

// IsObfuscated 返回 false，如果未设置混淆构建标签

// ff:
func IsObfuscated() bool {
	return false
}
