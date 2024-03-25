package build

// DesktopBuilder 用于构建桌面应用程序
type DesktopBuilder struct {
	*BaseBuilder
}

func newDesktopBuilder(options *Options) *DesktopBuilder {
	return &DesktopBuilder{
		BaseBuilder: NewBaseBuilder(options),
	}
}
