// +build !windows

package console

import (
	"fmt"
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

// Printf prints string to standard output.
func (p *ColorPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf("\x1b[%d;%dm%s\x1b[0m", p.bgColor, p.fgColor, fmt.Sprintf(format, a...))
}
