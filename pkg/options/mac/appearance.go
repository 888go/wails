package mac

// AppearanceType 是 Cocoa 窗口的一种外观类型
type AppearanceType string

const (
	// DefaultAppearance 使用默认的系统值
	DefaultAppearance AppearanceType = ""
	// NSAppearanceNameAqua - 标准的浅色系统外观。
	NSAppearanceNameAqua AppearanceType = "NSAppearanceNameAqua"
	// NSAppearanceNameDarkAqua - 深色系统外观的标准名称。
	NSAppearanceNameDarkAqua AppearanceType = "NSAppearanceNameDarkAqua"
	// NSAppearanceNameVibrantLight - 明亮活力外观
	NSAppearanceNameVibrantLight AppearanceType = "NSAppearanceNameVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastAqua - 一种标准浅色系统外观的高对比度版本。
	NSAppearanceNameAccessibilityHighContrastAqua AppearanceType = "NSAppearanceNameAccessibilityHighContrastAqua"
	// NSAppearanceNameAccessibilityHighContrastDarkAqua - 一种高对比度的标准深色系统外观版本。
	NSAppearanceNameAccessibilityHighContrastDarkAqua AppearanceType = "NSAppearanceNameAccessibilityHighContrastDarkAqua"
	// NSAppearanceNameAccessibilityHighContrastVibrantLight - 一种高对比度的浅色活力外观。
	NSAppearanceNameAccessibilityHighContrastVibrantLight AppearanceType = "NSAppearanceNameAccessibilityHighContrastVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastVibrantDark - 深色高对比度生动外观的高对比度版本。
	NSAppearanceNameAccessibilityHighContrastVibrantDark AppearanceType = "NSAppearanceNameAccessibilityHighContrastVibrantDark"
)
