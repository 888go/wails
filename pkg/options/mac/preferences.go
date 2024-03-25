package mac

import "github.com/leaanthony/u"

var (
	Enabled  = u.True
	Disabled = u.False
)

// Preferences 允许设置 webkit 的偏好设置
type Preferences struct {
// 一个布尔值，表示是否按下tab键会改变焦点至链接和表单控件。
// 默认设置为false。
	TabFocusesLinks u.Bool
// 一个布尔值，表示是否允许人们选择或以其他方式与文本进行交互。
// 默认设置为true。
	TextInteractionEnabled u.Bool
// 一个布尔值，表示 web 视图是否可以全屏显示内容。
// 默认设置为 false
	FullscreenEnabled u.Bool
}
