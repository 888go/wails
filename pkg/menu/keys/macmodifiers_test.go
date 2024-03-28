package keys

import "testing"

func TestToMacModifier(t *testing.T) {

	tests := []struct {
		name        string
		accelerator *Accelerator
		want        int
	}{
		// TODO: Add test cases.
		{"nil", nil, 0},
		{"empty", &Accelerator{}, 0},
		{"key", &Accelerator{X名称: "p"}, 0},
		{"cmd", X组合按键Cmd或Ctrl(""), NSEventModifierFlagCommand},
		{"ctrl", X组合按键Ctrl键(""), NSEventModifierFlagControl},
		{"shift", X组合按键Shift(""), NSEventModifierFlagShift},
		{"option", X组合按键Option或Alt键(""), NSEventModifierFlagOption},
		{"cmd+ctrl", X组合按键("", X常量_组合键_Cmd或Ctrl键, X常量_组合键_Ctrl键), NSEventModifierFlagCommand | NSEventModifierFlagControl},
		{"cmd+ctrl+shift", X组合按键("", X常量_组合键_Cmd或Ctrl键, X常量_组合键_Ctrl键, X常量_组合键_Shift键), NSEventModifierFlagCommand | NSEventModifierFlagControl | NSEventModifierFlagShift},
		{"cmd+ctrl+shift+option", X组合按键("", X常量_组合键_Cmd或Ctrl键, X常量_组合键_Ctrl键, X常量_组合键_Shift键, X常量_组合键_Option或Alt键), NSEventModifierFlagCommand | NSEventModifierFlagControl | NSEventModifierFlagShift | NSEventModifierFlagOption},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMacModifier(tt.accelerator); got != tt.want {
				t.Errorf("ToMacModifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
