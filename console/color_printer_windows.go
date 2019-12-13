package console

import (
	"fmt"
	"os"
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
	defaultTextAttributes = func() uint16 {
		info := new(windows.ConsoleScreenBufferInfo)
		if err := windows.GetConsoleScreenBufferInfo(windows.Handle(stdout), info); err != nil {
			err = fmt.Errorf("color printer init failed: %w", err)
			fmt.Println("error:", err)
			return 0x07
		}
		return info.Attributes
	}()
)

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
)

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uintptr) error {
	r, _, err := procSetConsoleTextAttribute.Call(hConsoleOutput, wAttributes)
	if r == 0 {
		return fmt.Errorf("set printer color failed: %w", err)
	}
	return nil
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

func (p *ColorPrinter) Write(value []byte) (n int, err error) {
	if err := setConsoleTextAttribute(stdout, uintptr(p.bgColor|p.fgColor)); err != nil {
		fmt.Println("error:", err)
	}
	if n, err = os.Stdout.Write(value); err != nil {
		return
	}
	setConsoleTextAttribute(stdout, uintptr(defaultTextAttributes))
	return
}
