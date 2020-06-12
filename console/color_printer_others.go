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
func NewColorPrinter(file *os.File, fgColor Color) *ColorPrinter {
	if fgColor < 0 {
		fgColor = 0
	}
	return &ColorPrinter{
		file:    file,
		fgColor: fgColor,
	}
}

func (p *ColorPrinter) SetBackgroundColor(color Color) *ColorPrinter {
	if color < 0 {
		p.bgColor = 0
	} else {
		p.bgColor = color + 10
	}
	return p
}

func (p *ColorPrinter) SetForegroundColor(color Color) *ColorPrinter {
	if color < 0 {
		p.fgColor = 0
	} else {
		p.fgColor = color
	}
	return p
}

func (p *ColorPrinter) Write(value []byte) (n int, err error) {
	return p.file.Write([]byte(fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", p.bgColor, p.fgColor, string(value))))
}
