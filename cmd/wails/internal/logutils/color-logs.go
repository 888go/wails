package logutils

import (
	"fmt"

	"github.com/888go/wails/internal/colour"
)


// ff:
// args:
// message:
func LogGreen(message string, args ...interface{}) {
	if len(message) == 0 {
		return
	}
	text := fmt.Sprintf(message, args...)
	println(colour.Green(text))
}


// ff:
// args:
// message:
func LogRed(message string, args ...interface{}) {
	if len(message) == 0 {
		return
	}
	text := fmt.Sprintf(message, args...)
	println(colour.Red(text))
}


// ff:
// args:
// message:
func LogDarkYellow(message string, args ...interface{}) {
	if len(message) == 0 {
		return
	}
	text := fmt.Sprintf(message, args...)
	println(colour.DarkYellow(text))
}
