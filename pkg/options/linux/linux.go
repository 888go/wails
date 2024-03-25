package linux

// WebviewGpuPolicy 用于确定 webview 硬件加速策略的值。
type WebviewGpuPolicy int

const (
	// WebviewGpuPolicyAlways 硬件加速始终启用。
	WebviewGpuPolicyAlways WebviewGpuPolicy = iota
	// WebviewGpuPolicyOnDemand 当网页内容请求时，硬件加速启用/禁用。
	WebviewGpuPolicyOnDemand
	// WebviewGpuPolicyNever 硬件加速始终禁用。
	WebviewGpuPolicyNever
)

// 以下选项针对Linux系统构建时特定
type Options struct {
// Icon 设置代表窗口的图标。当窗口被最小化（也称为图标化）时，会使用这个图标。
	Icon []byte

	// WindowIsTranslucent 在启用时将窗口背景设置为透明
	WindowIsTranslucent bool

	// Messages 是可以自定义的消息
	Messages *Messages

// WebviewGpuPolicy 用于确定 webview 的硬件加速策略。
//   - WebviewGpuPolicyAlways   // 始终启用硬件加速
//   - WebviewGpuPolicyOnDemand // 按需启用硬件加速
//   - WebviewGpuPolicyNever    // 从不启用硬件加速
// 由于 https://github.com/wailsapp/wails/issues/2977 中的问题，如果在调用 wails.Run() 时 options.Linux 为 nil，
// 则默认将 WebviewGpuPolicy 设置为 WebviewGpuPolicyNever。
// 客户端代码可以通过传递非空的 Options 并按需设置 WebviewGpuPolicy 来覆盖此默认行为。
	WebviewGpuPolicy WebviewGpuPolicy

// ProgramName 用于通过 GTK 的 g_set_prgname() 函数设置程序名称，以便在窗口管理器中显示。
// 此名称不应进行本地化处理。[参见文档]
//
// 当创建 .desktop 文件时，如果 .desktop 文件的 Name 属性与可执行文件名不同，
// 这个值有助于窗口分组和桌面图标功能。
//
// [参见文档]: https://docs.gtk.org/glib/func.set_prgname.html
	ProgramName string
}

type Messages struct {
	WebKit2GTKMinRequired string
}


// ff:
func DefaultMessages() *Messages {
	return &Messages{
		WebKit2GTKMinRequired: "This application requires at least WebKit2GTK %s to be installed.",
	}
}
