package colour

import (
	"fmt"
	"strings"

	"github.com/wzshiming/ctc"
)

var ColourEnabled = true


// ff:
// text:
// col:
func Col(col ctc.Color, text string) string {
	if !ColourEnabled {
		return text
	}
	return fmt.Sprintf("%s%s%s", col, text, ctc.Reset)
}


// ff:
// text:
func Yellow(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightYellow, text)
}


// ff:
// text:
func Red(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightRed, text)
}


// ff:
// text:
func Blue(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightBlue, text)
}


// ff:
// text:
func Green(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightGreen, text)
}


// ff:
// text:
func Cyan(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightCyan, text)
}


// ff:
// text:
func Magenta(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightMagenta, text)
}


// ff:
// text:
func White(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightWhite, text)
}


// ff:
// text:
func Black(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBrightBlack, text)
}


// ff:
// text:
func DarkYellow(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundYellow, text)
}


// ff:
// text:
func DarkRed(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundRed, text)
}


// ff:
// text:
func DarkBlue(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBlue, text)
}


// ff:
// text:
func DarkGreen(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundGreen, text)
}


// ff:
// text:
func DarkCyan(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundCyan, text)
}


// ff:
// text:
func DarkMagenta(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundMagenta, text)
}


// ff:
// text:
func DarkWhite(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundWhite, text)
}


// ff:
// text:
func DarkBlack(text string) string {
	if !ColourEnabled {
		return text
	}
	return Col(ctc.ForegroundBlack, text)
}

var rainbowCols = []func(string) string{Red, Yellow, Green, Cyan, Blue, Magenta}


// ff:
// text:
func Rainbow(text string) string {
	if !ColourEnabled {
		return text
	}
	var builder strings.Builder

	for i := 0; i < len(text); i++ {
		fn := rainbowCols[i%len(rainbowCols)]
		builder.WriteString(fn(text[i : i+1]))
	}

	return builder.String()
}
