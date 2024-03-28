package keys

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {

	i := is.New(t)

	type args struct {
		Input    string
		Expected *Accelerator
	}

	gooddata := []args{
		{"CmdOrCtrl+A", X组合按键Cmd或Ctrl("A")},
		{"SHIFT+.", X组合按键Shift(".")},
		{"CTRL+plus", X组合按键Ctrl键("+")},
		{"CTRL+SHIFT+escApe", X组合按键("escape", X常量_组合键_Ctrl键, X常量_组合键_Shift键)},
		{";", X按键(";")},
		{"OptionOrAlt+Page Down", X组合按键Option或Alt键("Page Down")},
	}
	for _, tt := range gooddata {
		result, err := Parse(tt.Input)
		i.NoErr(err)
		i.Equal(result, tt.Expected)
	}
	baddata := []string{"CmdOrCrl+A", "SHIT+.", "CTL+plus", "CTRL+SHIF+esApe", "escap", "Sper+Tab", "OptionOrAlt"}
	for _, d := range baddata {
		result, err := Parse(d)
		i.True(err != nil)
		i.Equal(result, nil)
	}
}
