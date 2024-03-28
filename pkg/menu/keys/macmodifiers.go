package keys

const (
	NSEventModifierFlagShift   = 1 << 17 // 设置是否按下Shift键。
	NSEventModifierFlagControl = 1 << 18 // 设置是否按下Control键。
	NSEventModifierFlagOption  = 1 << 19 // 设置是否按下Option（在Mac系统中）或Alternate（在其他系统中，通常指Alt键）键。
	NSEventModifierFlagCommand = 1 << 20 // 设置是否按下Command键。
)

var macModifierMap = map[Modifier]int{
	CmdOrCtrlKey:   NSEventModifierFlagCommand,
	ControlKey:     NSEventModifierFlagControl,
	OptionOrAltKey: NSEventModifierFlagOption,
	ShiftKey:       NSEventModifierFlagShift,
}


// ff:
// accelerator:
func ToMacModifier(accelerator *Accelerator) int {
	if accelerator == nil {
		return 0
	}
	result := 0
	for _, modifier := range accelerator.Modifiers {
		result |= macModifierMap[modifier]
	}
	return result
}
