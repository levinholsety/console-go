package console

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/levinholsety/common-go/comm"
	"golang.org/x/crypto/ssh/terminal"
)

// SpeedCalculator represents a function to calculate speed with total value and elapsed time.
type SpeedCalculator func(n int64, elapsed time.Duration) string

// NewProgressBar creates an instance of ProgressBar and returns it.
func NewProgressBar(maxValue int64) (bar *ProgressBar) {
	bar = &ProgressBar{
		epoch:     time.Now(),
		barLength: 50,
		MaxValue:  maxValue,
	}
	bar.print()
	return
}

// ProgressBar can print progress bar in console.
type ProgressBar struct {
	prt             ColorPrinter
	epoch           time.Time
	barLength       uint32
	MaxValue        int64
	value           int64
	speedCalculator SpeedCalculator
	percent         uint32
	elapsed         time.Duration
	lock            sync.Mutex
}

// SetColorPrinter sets the color printer with which the current progress bar prints.
func (p *ProgressBar) SetColorPrinter(prt ColorPrinter) {
	p.prt = prt
}

// SetSpeedCalculator sets speed calculator.
func (p *ProgressBar) SetSpeedCalculator(calc SpeedCalculator) {
	p.speedCalculator = calc
}

// SetProgress sets current progress.
func (p *ProgressBar) SetProgress(val int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.value = val
	if p.isChanged() {
		p.print()
	}
}

// AddProgress adds value to current progress.
func (p *ProgressBar) AddProgress(delta int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.value += delta
	if p.isChanged() {
		p.print()
	}
}

func (p *ProgressBar) isChanged() bool {
	if p.value < 0 {
		p.value = 0
	} else if p.value > p.MaxValue {
		p.value = p.MaxValue
	}
	percent := uint32(p.value * 100 / p.MaxValue)
	elapsed := time.Now().Sub(p.epoch)
	if percent == p.percent && elapsed.Milliseconds() == p.elapsed.Milliseconds() {
		return false
	}
	p.percent = percent
	p.elapsed = elapsed
	return true
}

func (p *ProgressBar) print() {
	progressBar := make([]byte, p.barLength)
	comm.FillBytes(progressBar[:p.percent*p.barLength/100], '=')
	text := fmt.Sprintf("\r[%s] %d%% %s", string(progressBar), p.percent, formatDuration(p.elapsed))
	if p.speedCalculator != nil {
		text += " " + p.speedCalculator(p.value, p.elapsed)
	}
	if width, _, err := terminal.GetSize(int(os.Stdout.Fd())); err == nil {
		textLen := len(text)
		if textLen > width {
			text = text[:width]
		} else if textLen < width {
			text += string(bytes.Repeat([]byte{0x20}, width-textLen))
		}
	}
	if p.prt == nil {
		fmt.Print(text)
	} else {
		p.prt.Print(text)
	}
}

const (
	timeDurationFormat         = "%02d:%02d:%02d.%03d"
	timeDurationWithDaysFormat = "%dd " + timeDurationFormat
)

func formatDuration(value time.Duration) string {
	s, ms := div(value.Milliseconds(), 1000)
	m, s := div(s, 60)
	h, m := div(m, 60)
	d, h := div(h, 24)
	if d > 0 {
		return fmt.Sprintf(timeDurationWithDaysFormat, d, h, m, s, ms)
	}
	return fmt.Sprintf(timeDurationFormat, h, m, s, ms)
}

func div(a, b int64) (int64, int64) {
	return a / b, a % b
}
