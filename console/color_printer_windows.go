package console

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

// Colors
const (
	Black Color = iota
	Blue
	Green
	Aqua
	Red
	Purple
	Yellow
	White
	Gray
	LightBlue
	LightGreen
	LightAqua
	LightRed
	LightPurple
	LightYellow
	LightWhite
)

var (
	stdout                = uintptr(syscall.Stdout)
	defaultTextAttributes uint16
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
)

func init() {
	info := new(windows.ConsoleScreenBufferInfo)
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(stdout), info); err != nil {
		panic(err)
	}
	defaultTextAttributes = info.Attributes
}

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uintptr) {
	r, _, err := procSetConsoleTextAttribute.Call(hConsoleOutput, wAttributes)
	if r == 0 {
		panic(err)
	}
}

// NewColorPrinter creates a new ColorPrinter instance.
func NewColorPrinter(bgColor Color, fgColor Color) *ColorPrinter {
	if bgColor < 0 {
		bgColor = Color(defaultTextAttributes) & 0o11110000
	} else {
		bgColor <<= 4
	}
	if fgColor < 0 {
		fgColor = Color(defaultTextAttributes) & 0o1111
	}
	return &ColorPrinter{
		bgColor: int(bgColor),
		fgColor: int(fgColor),
	}
}

// Printf prints string to standard output.
func (p *ColorPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	setConsoleTextAttribute(stdout, uintptr(p.bgColor|p.fgColor))
	n, err = fmt.Printf(format, a...)
	setConsoleTextAttribute(stdout, uintptr(defaultTextAttributes))
	return
}
