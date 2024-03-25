package mac

// 定义 ActivationPolicy 类型为 int
//
// 下面是 ActivationPolicy 的常量枚举值：
// NSApplicationActivationPolicyRegular    表示常规激活策略，ActivationPolicy 值为 0
// NSApplicationActivationPolicyAccessory  表示辅助激活策略，ActivationPolicy 值为 1
// NSApplicationActivationPolicyProhibited 表示禁止激活策略，ActivationPolicy 值为 2
// type ActivationPolicy int
// 这段 Go 代码定义了一个名为 `ActivationPolicy` 的整数类型，并列举了三种该类型的预定义常量，这些常量用于表示应用程序的不同激活策略。

type AboutInfo struct {
	Title   string
	Message string
	Icon    []byte
}

// Options 是 Mac 系统特有的选项
type Options struct {
	TitleBar             *TitleBar
	Appearance           AppearanceType
	WebviewIsTransparent bool
	WindowIsTranslucent  bool
	Preferences          *Preferences
	// ActivationPolicy 激活策略
	About      *AboutInfo
	OnFileOpen func(filePath string) `json:"-"`
	OnUrlOpen  func(filePath string) `json:"-"`
	// URLHandlers 是一个映射，其键为字符串类型，值为接收一个字符串参数的函数类型。这个映射用于处理与特定URL路径关联的处理器函数。
}
