package console

import (
	"fmt"
	"strings"
)

type ProgressBar struct {
	Width int
	Total int
}

func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		Width: 50,
		Total: total,
	}
}

func (p *ProgressBar) Progress(current int) {
	progressLen := current * p.Width / p.Total
	progress := strings.Repeat("=", progressLen)
	blanks := strings.Repeat(" ", p.Width-progressLen)
	fmt.Printf("\r[%s%s] %d%%", progress, blanks, current*100/p.Total)
}
