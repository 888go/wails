package keys

import (
	"strings"

	"github.com/leaanthony/slicer"
)

var modifierStringMap = map[string]map[Modifier]string{
	"windows": {
		X常量_组合键_Cmd或Ctrl键:   "Ctrl",
		X常量_组合键_Ctrl键:     "Ctrl",
		X常量_组合键_Option或Alt键: "Alt",
		X常量_组合键_Shift键:       "Shift",
		// SuperKey:       "Win",
	},
	"darwin": {
		X常量_组合键_Cmd或Ctrl键:   "Cmd",
		X常量_组合键_Ctrl键:     "Ctrl",
		X常量_组合键_Option或Alt键: "Option",
		X常量_组合键_Shift键:       "Shift",
		// SuperKey:       "Cmd",
	},
	"linux": {
		X常量_组合键_Cmd或Ctrl键:   "Ctrl",
		X常量_组合键_Ctrl键:     "Ctrl",
		X常量_组合键_Option或Alt键: "Alt",
		X常量_组合键_Shift键:       "Shift",
		// SuperKey:       "Super", 
// 翻译：// 超级键：      "Super"
	},
}


// ff:
// platform:
// accelerator:
func Stringify(accelerator *Accelerator, platform string) string {
	result := slicer.String()
	for _, modifier := range accelerator.X修饰符 {
		result.Add(modifierStringMap[platform][modifier])
	}
	result.Deduplicate()
	result.Add(strings.ToUpper(accelerator.X名称))
	return result.Join("+")
}
