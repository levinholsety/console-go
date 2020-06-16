package main

import (
	"os"
	"time"

	"github.com/levinholsety/console-go/console"
)

func main() {
	pg := console.NewProgressBar(100)
	pg.SetColorPrinter(console.NewColorPrinter(os.Stdout).SetForegroundColor(console.Gray))
	for i := int64(0); i < pg.MaxValue; i++ {
		pg.AddProgress(1)
		time.Sleep(time.Millisecond * 30)
	}
}
