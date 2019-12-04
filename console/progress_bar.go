package console

import (
	"fmt"
	"strings"
	"time"

	"github.com/levinholsety/common-go/timeutil"
)

type ProgressBar struct {
	epoch time.Time
	Width int
	Total int
}

func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		epoch: time.Now(),
		Width: 50,
		Total: total,
	}
}

func (p *ProgressBar) Progress(current int) {
	progressLen := current * p.Width / p.Total
	progress := strings.Repeat("=", progressLen)
	blanks := strings.Repeat(" ", p.Width-progressLen)
	fmt.Printf("\r[%s%s] %d%% %s", progress, blanks, current*100/p.Total, timeutil.FormatDuration(time.Now().Sub(p.epoch)))
}
