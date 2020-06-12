package console

import (
	"fmt"
	"os"
)

// Color represents a color in console.
type Color int

// DefaultColor
const (
	DefaultColor Color = -1
)

// ColorPrinter can print string in specified background color and foreground color.
type ColorPrinter struct {
	file           *os.File
	defaultBgColor Color
	defaultFgColor Color
	bgColor        Color
	fgColor        Color
}

func (p *ColorPrinter) SetFile(file *os.File) *ColorPrinter {
	p.file = file
	return p
}

// Printf prints string to standard output.
func (p *ColorPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	return p.Write([]byte(fmt.Sprintf(format, a...)))
}

// Print prints string to standard output.
func (p *ColorPrinter) Print(a ...interface{}) (n int, err error) {
	return p.Write([]byte(fmt.Sprint(a...)))
}

// Println prints string to standard output.
func (p *ColorPrinter) Println(a ...interface{}) (n int, err error) {
	return p.Write([]byte(fmt.Sprintln(a...)))
}

// // ColorPrint prints string in specified background color and foreground color.
// func ColorPrint(bgColor Color, fgColor Color, a ...interface{}) (int, error) {
// 	return NewColorPrinter(bgColor, fgColor).Print(a...)
// }

// // ColorPrintln prints string in specified background color and foreground color.
// func ColorPrintln(bgColor Color, fgColor Color, a ...interface{}) (int, error) {
// 	return NewColorPrinter(bgColor, fgColor).Println(a...)
// }

// // ColorPrintf prints string in specified background color and foreground color.
// func ColorPrintf(bgColor Color, fgColor Color, format string, a ...interface{}) (int, error) {
// 	return NewColorPrinter(bgColor, fgColor).Printf(format, a...)
// }
