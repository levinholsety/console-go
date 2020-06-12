package console

import (
	"os"
	"sync"
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
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTextAttribute = kernel32.NewProc("SetConsoleTextAttribute")
)

func setConsoleTextAttribute(hConsoleOutput uintptr, wAttributes uintptr) {
	procSetConsoleTextAttribute.Call(hConsoleOutput, wAttributes)
}

func NewColorPrinter(file *os.File, fgColor Color) (prt *ColorPrinter) {
	prt = &ColorPrinter{
		file:    file,
		fgColor: fgColor,
	}
	info := new(windows.ConsoleScreenBufferInfo)
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(file.Fd()), info); err == nil {
		prt.defaultBgColor = Color(info.Attributes & 0b11110000)
		prt.defaultFgColor = Color(info.Attributes & 0b1111)
	} else {
		prt.defaultBgColor = Black << 4
		prt.defaultFgColor = White
	}
	prt.bgColor = prt.defaultBgColor
	if prt.fgColor == DefaultColor {
		prt.fgColor = prt.defaultFgColor
	}
	return
}

var lock = new(sync.Mutex)

func (p *ColorPrinter) SetBackgroundColor(color Color) *ColorPrinter {
	if color == DefaultColor {
		p.bgColor = p.defaultBgColor
	} else {
		p.bgColor = color << 4
	}
	return p
}

func (p *ColorPrinter) SetForegroundColor(color Color) *ColorPrinter {
	if color == DefaultColor {
		p.fgColor = p.defaultFgColor
	} else {
		p.fgColor = color
	}
	return p
}

func (p *ColorPrinter) Write(value []byte) (n int, err error) {
	lock.Lock()
	defer lock.Unlock()
	setConsoleTextAttribute(p.file.Fd(), uintptr(p.bgColor|p.fgColor))
	n, err = p.file.Write(value)
	setConsoleTextAttribute(p.file.Fd(), uintptr(p.defaultBgColor|p.defaultFgColor))
	return
}
