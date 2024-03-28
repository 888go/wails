package keys

const (
	NSEventModifierFlagShift   = 1 << 17 // 设置是否按下Shift键。
	NSEventModifierFlagControl = 1 << 18 // 设置是否按下Control键。
	NSEventModifierFlagOption  = 1 << 19 // 设置是否按下Option（在Mac系统中）或Alternate（在其他系统中，通常指Alt键）键。
	NSEventModifierFlagCommand = 1 << 20 // 设置是否按下Command键。
)

var macModifierMap = map[Modifier]int{
	X常量_组合键_Cmd或Ctrl键:   NSEventModifierFlagCommand,
	X常量_组合键_Ctrl键:     NSEventModifierFlagControl,
	X常量_组合键_Option或Alt键: NSEventModifierFlagOption,
	X常量_组合键_Shift键:       NSEventModifierFlagShift,
}


// ff:
// accelerator:
func ToMacModifier(accelerator *Accelerator) int {
	if accelerator == nil {
		return 0
	}
	result := 0
	for _, modifier := range accelerator.X修饰符 {
		result |= macModifierMap[modifier]
	}
	return result
}
