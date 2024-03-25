package options

// 以下调试选项在调试构建中会被考虑使用。
type Debug struct {
	// OpenInspectorOnStartup 在应用程序启动时打开检查器。
	OpenInspectorOnStartup bool
}
