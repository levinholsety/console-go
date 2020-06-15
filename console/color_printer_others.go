// +build !windows

package console

import (
	"fmt"
	"os"
	"strings"
)

// Colors
const (
	Black Color = iota + 1
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
	if p.bgColor >= Black && p.bgColor <= White {
		return fmt.Sprintf("%d", p.bgColor+39)
	}
	if p.bgColor >= Gray && p.bgColor <= LightWhite {
		return fmt.Sprintf("%d;1", p.bgColor+31)
	}
	return ""
}

func (p *colorPrinter) fgColorCode() string {
	if p.fgColor >= Black && p.fgColor <= White {
		return fmt.Sprintf("%d", p.fgColor+29)
	}
	if p.fgColor >= Gray && p.fgColor <= LightWhite {
		return fmt.Sprintf("%d;1", p.fgColor+21)
	}
	return ""
}

func (p *colorPrinter) write(text string) (n int, err error) {
	bgCode := p.bgColorCode()
	fgCode := p.fgColorCode()
	var code string
	if len(bgCode) > 0 && len(fgCode) > 0 {
		code = bgCode + ";" + fgCode
	} else if len(bgCode) > 0 {
		code = bgCode
	} else if len(fgCode) > 0 {
		code = fgCode
	} else {
		code = "0"
	}
	return p.file.Write([]byte(fmt.Sprintf("\033[%sm%s\033[0m", code, strings.ReplaceAll(text, "\n", fmt.Sprintf("\033[0m\n\033[%sm", code)))))
}
