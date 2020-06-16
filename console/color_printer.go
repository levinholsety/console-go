package console

import "fmt"

// Color represents a color in console.
type Color int

func (v Color) String() string {
	switch v {
	case Black:
		return "Black"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Aqua:
		return "Aqua"
	case Red:
		return "Red"
	case Purple:
		return "Purple"
	case Yellow:
		return "Yellow"
	case White:
		return "White"
	case Gray:
		return "Gray"
	case LightBlue:
		return "LightBlue"
	case LightGreen:
		return "LightGreen"
	case LightAqua:
		return "LightAqua"
	case LightRed:
		return "LightRed"
	case LightPurple:
		return "LightPurple"
	case LightYellow:
		return "LightYellow"
	case LightWhite:
		return "LightWhite"
	default:
		return "Unknown"
	}
}

// ColorPrinter can print color text on console.
type ColorPrinter interface {
	SetBackgroundColor(color Color) ColorPrinter
	SetForegroundColor(color Color) ColorPrinter
	ResetColors() ColorPrinter
	Printf(format string, a ...interface{}) (n int, err error)
	Print(a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
}

func (p *colorPrinter) SetBackgroundColor(color Color) ColorPrinter {
	p.bgColor = color
	return p
}

func (p *colorPrinter) SetForegroundColor(color Color) ColorPrinter {
	p.fgColor = color
	return p
}

func (p *colorPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	return p.write(fmt.Sprintf(format, a...))
}

func (p *colorPrinter) Print(a ...interface{}) (n int, err error) {
	return p.write(fmt.Sprint(a...))
}

func (p *colorPrinter) Println(a ...interface{}) (n int, err error) {
	return p.write(fmt.Sprintln(a...))
}
