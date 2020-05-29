package console

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/levinholsety/common-go/comm"
	"golang.org/x/crypto/ssh/terminal"
)

// SpeedCalculator represents a function to calculate speed with total value and elapsed time.
type SpeedCalculator func(n int64, elapsed time.Duration) string

// NewProgressBar creates an instance of ProgressBar.
func NewProgressBar(total int64) (pb *ProgressBar) {
	pb = &ProgressBar{
		epoch:           time.Now(),
		barLength:       50,
		total:           total,
		progressChannel: make(chan (int64), 10),
	}
	go pb.listen()
	return
}

// ProgressBar can print progress bar in console.
type ProgressBar struct {
	epoch           time.Time
	barLength       int
	total           int64
	current         int64
	speedCalculator SpeedCalculator
	progressChannel chan (int64)
	wg              sync.WaitGroup
}

const interval = time.Millisecond * 1e2

func (p *ProgressBar) listen() {
	ts := time.Now()
	for v := range p.progressChannel {
		if time.Now().Sub(ts) < interval {
			continue
		}
		p.printProgress(v)
		ts = time.Now()
	}
}

// Close closes a progress bar.
func (p *ProgressBar) Close() {
	close(p.progressChannel)
	p.printProgress(p.current)
}

// SetSpeedCalculator sets speed calculator.
func (p *ProgressBar) SetSpeedCalculator(calc SpeedCalculator) {
	p.speedCalculator = calc
}

func (p *ProgressBar) printProgress(current int64) {
	progressBar := make([]byte, p.barLength)
	comm.FillBytes(progressBar[:int(p.current*int64(p.barLength)/p.total)], '=')
	elapsed := time.Now().Sub(p.epoch)
	text := fmt.Sprintf("\r[%s] %d%% %s", string(progressBar), current*100/p.total, comm.FormatTimeDuration(elapsed))
	if p.speedCalculator != nil {
		text += " " + p.speedCalculator(current, elapsed)
	}
	width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return
	}
	if width < len(text) {
		text = text[:width]
	}
	fmt.Print(text)
}

// SetProgress sets current progress.
func (p *ProgressBar) SetProgress(current int64) {
	if current < 0 {
		p.current = 0
	} else if current > p.total {
		p.current = p.total
	} else {
		p.current = current
	}
	p.progressChannel <- p.current
}

// AddProgress adds value to current progress.
func (p *ProgressBar) AddProgress(value int64) {
	p.SetProgress(p.current + value)
}
