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

// NewColorPrinter creates an instance of ColorPrinter and returns it.
func NewColorPrinter(file *os.File) ColorPrinter {
	p := &colorPrinter{
		file: file,
	}
	info := new(windows.ConsoleScreenBufferInfo)
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(file.Fd()), info); err == nil {
		p.defBgColor, p.defFgColor = Color(info.Attributes>>4), Color(info.Attributes&0b1111)
	} else {
		p.defBgColor, p.defFgColor = Black, White
	}
	p.bgColor, p.fgColor = p.defBgColor, p.defFgColor
	return p
}

type colorPrinter struct {
	file       *os.File
	defBgColor Color
	defFgColor Color
	bgColor    Color
	fgColor    Color
}

func (p *colorPrinter) SetBackgroundColor(color Color) ColorPrinter {
	p.bgColor = color
	return p
}

func (p *colorPrinter) SetForegroundColor(color Color) ColorPrinter {
	p.fgColor = color
	return p
}

var lock = new(sync.Mutex)

func (p *colorPrinter) attributes() uintptr {
	return uintptr(p.bgColor<<4 | p.fgColor)
}

func (p *colorPrinter) defAttributes() uintptr {
	return uintptr(p.defBgColor<<4 | p.defFgColor)
}

func (p *colorPrinter) write(text string) (n int, err error) {
	lock.Lock()
	defer lock.Unlock()
	setConsoleTextAttribute(p.file.Fd(), p.attributes())
	n, err = p.file.Write([]byte(text))
	setConsoleTextAttribute(p.file.Fd(), p.defAttributes())
	return
}
