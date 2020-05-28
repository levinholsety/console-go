package console

import (
	"fmt"
	"strings"
	"time"

	"github.com/levinholsety/common-go/comm"
)

// ProgressBar can print progress bar in console.
type ProgressBar struct {
	epoch time.Time
	Width int
	Total int
}

// NewProgressBar creates an instance of ProgressBar.
func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		epoch: time.Now(),
		Width: 50,
		Total: total,
	}
}

// Progress changes progress.
func (p *ProgressBar) Progress(current int) {
	progressLen := current * p.Width / p.Total
	progress := strings.Repeat("=", progressLen)
	blanks := strings.Repeat(" ", p.Width-progressLen)
	fmt.Printf("\r[%s%s] %d%% %s", progress, blanks, current*100/p.Total, comm.FormatTimeDuration(time.Now().Sub(p.epoch)))
}
