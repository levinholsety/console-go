// +build !windows

package console

import (
	"fmt"
	"os"
	"strings"
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

const (
	Gray Color = (iota + 30) | (1 << 8)
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightPurple
	LightAqua
	LightWhite
)

func (v Color) fgCode() string {
	color := v & 0b11111111
	highlight := v >> 8
	if highlight == 0 {
		return fmt.Sprintf("%d", color)
	}
	return fmt.Sprintf("%d;%d", color, highlight)
}

func (v Color) bgCode() string {
	color := v&0b11111111 + 10
	highlight := v >> 8
	if highlight == 0 {
		return fmt.Sprintf("%d", color)
	}
	return fmt.Sprintf("%d;%d", color, highlight)
}

// NewColorPrinter creates a new ColorPrinter instance.
func NewColorPrinter(file *os.File) ColorPrinter {
	return &colorPrinter{
		file: file,
	}
}

type colorPrinter struct {
	file    *os.File
	fgColor Color
	bgColor Color
}

func (p *colorPrinter) ResetColors() ColorPrinter {
	p.fgColor, p.bgColor = 0, 0
	return p
}

const codeReset = "\033[0m"

func (p *colorPrinter) colorCode() (code string) {
	fgCode := p.fgColor.fgCode()
	bgCode := p.bgColor.bgCode()
	if len(fgCode) > 0 && len(bgCode) > 0 {
		code = fgCode + ";" + bgCode
	} else if len(fgCode) > 0 {
		code = fgCode
	} else if len(bgCode) > 0 {
		code = bgCode
	} else {
		code = "0"
	}
	code = fmt.Sprintf("\033[%sm", code)
	return
}

func (p *colorPrinter) write(text string) (n int, err error) {
	codeColor := p.colorCode()
	text = strings.ReplaceAll(text, "\n", codeReset+"\n"+codeColor)
	text = codeColor + text + codeReset
	return p.file.Write([]byte(text))
}
