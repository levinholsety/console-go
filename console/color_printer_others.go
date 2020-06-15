// +build !windows

package console

import (
	"fmt"
	"os"
	"strings"
)

// Colors
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Aqua
	White
	Gray
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightPurple
	LightAqua
	LightWhite
)

// NewColorPrinter creates a new ColorPrinter instance.
func NewColorPrinter(file *os.File) ColorPrinter {
	return &colorPrinter{
		file: file,
	}
}

type colorPrinter struct {
	file    *os.File
	bgColor Color
	fgColor Color
}

func (p *colorPrinter) SetBackgroundColor(color Color) ColorPrinter {
	p.bgColor = color
	return p
}

func (p *colorPrinter) SetForegroundColor(color Color) ColorPrinter {
	p.fgColor = color
	return p
}

func (p *colorPrinter) bgColorCode() string {
	if p.bgColor < Gray {
		return fmt.Sprintf("%d", p.bgColor+40)
	}
	return fmt.Sprintf("%d;1", p.bgColor+32)
}

func (p *colorPrinter) fgColorCode() string {
	if p.fgColor < Gray {
		return fmt.Sprintf("%d", p.fgColor+30)
	}
	return fmt.Sprintf("%d;1", p.fgColor+22)
}

func (p *colorPrinter) write(text string) (n int, err error) {
	return p.file.Write([]byte(fmt.Sprintf("\033[%s;%sm%s\033[0m", p.bgColorCode(), p.fgColorCode(), strings.ReplaceAll(text, "\n", "\033[0m\n"))))
}
