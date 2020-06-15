package main

import (
	"os"

	"github.com/levinholsety/console-go/console"
)

var (
	bgColors = []console.Color{
		console.Black,
		console.Blue,
		console.Green,
		console.Aqua,
		console.Red,
		console.Purple,
		console.Yellow,
		console.White,
		console.Gray,
		console.LightBlue,
		console.LightGreen,
		console.LightAqua,
		console.LightRed,
		console.LightPurple,
		console.LightYellow,
		console.LightWhite,
	}
	fgColors = []console.Color{
		console.Black,
		console.Blue,
		console.Green,
		console.Aqua,
		console.Red,
		console.Purple,
		console.Yellow,
		console.White,
		console.Gray,
		console.LightBlue,
		console.LightGreen,
		console.LightAqua,
		console.LightRed,
		console.LightPurple,
		console.LightYellow,
		console.LightWhite,
	}
)

func main() {
	prt := console.NewColorPrinter(os.Stdout)
	for _, fgColor := range fgColors {
		prt.SetForegroundColor(fgColor)
		prt.Printf("%s text on default background\n", fgColor)
	}
	prt = console.NewColorPrinter(os.Stdout)
	for _, color := range bgColors {
		prt.SetBackgroundColor(color)
		prt.Printf("default text on %s background\n", color)
	}
	prt = console.NewColorPrinter(os.Stdout)
	for _, bgColor := range bgColors {
		for _, fgColor := range fgColors {
			if bgColor == fgColor {
				continue
			}
			prt.SetBackgroundColor(bgColor)
			prt.SetForegroundColor(fgColor)
			prt.Printf("%s text on %s background\n", fgColor, bgColor)
		}
	}
}
