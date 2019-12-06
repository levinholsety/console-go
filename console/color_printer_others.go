// +build !windows

package console

import (
	"fmt"
	"os"
)

// Colors
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Aqua
	White
)

// Colors
const (
	Gray Color = iota + 30
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightPurple
	LightAqua
	LightWhite
)

// NewColorPrinter creates a new ColorPrinter instance.
func NewColorPrinter(bgColor Color, fgColor Color) *ColorPrinter {
	if bgColor < 0 {
		bgColor = 0
	} else {
		bgColor += 10
	}
	if fgColor < 0 {
		fgColor = 0
	}
	return &ColorPrinter{
		bgColor: int(bgColor),
		fgColor: int(fgColor),
	}
}

func (p *ColorPrinter) Write(value []byte) (n int, err error) {
	return os.Stdout.Write([]byte(fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", p.bgColor, p.fgColor, string(value))))
}
