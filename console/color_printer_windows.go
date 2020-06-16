package console

import (
	"os"
	"strings"
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
		p.defFgColor, p.defBgColor = Color(info.Attributes&0b1111), Color(info.Attributes>>4)
	} else {
		p.defFgColor, p.defBgColor = White, Black
	}
	p.fgColor, p.bgColor = p.defFgColor, p.defBgColor
	return p
}

type colorPrinter struct {
	file       *os.File
	defFgColor Color
	defBgColor Color
	fgColor    Color
	bgColor    Color
}

func (p *colorPrinter) ResetColors() ColorPrinter {
	p.fgColor, p.bgColor = p.defFgColor, p.defBgColor
	return p
}

var lock = new(sync.Mutex)

func (p *colorPrinter) attributes() uintptr {
	return uintptr(p.fgColor | p.bgColor<<4)
}

func (p *colorPrinter) defAttributes() uintptr {
	return uintptr(p.defFgColor | p.defBgColor<<4)
}

func (p *colorPrinter) write(text string) (n int, err error) {
	lock.Lock()
	defer lock.Unlock()
	lines := strings.Split(text, "\n")
	lastIndex := len(lines) - 1
	for i, line := range lines {
		var count int
		setConsoleTextAttribute(p.file.Fd(), p.attributes())
		count, err = p.file.Write([]byte(line))
		if err != nil {
			return
		}
		n += count
		setConsoleTextAttribute(p.file.Fd(), p.defAttributes())
		if i < lastIndex {
			count, err = p.file.Write([]byte{'\n'})
			if err != nil {
				return
			}
			n += count
		}
	}
	return
}
