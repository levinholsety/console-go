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

// ExecuteWithProgressBar creates an instance of ProgressBar.
func ExecuteWithProgressBar(task func(bar *ProgressBar) error, maxValue int64) (err error) {
	bar := &ProgressBar{
		epoch:           time.Now(),
		barLength:       50,
		maxValue:        maxValue,
		interval:        time.Millisecond * 10,
		progressChannel: make(chan int64, 10),
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go bar.listen(wg)
	err = func() error {
		defer close(bar.progressChannel)
		return task(bar)
	}()
	wg.Wait()
	return
}

// ProgressBar can print progress bar in console.
type ProgressBar struct {
	epoch           time.Time
	barLength       int
	maxValue        int64
	current         int64
	speedCalculator SpeedCalculator
	progressChannel chan int64
	interval        time.Duration
}

func (p *ProgressBar) listen(wg *sync.WaitGroup) {
	ts := time.Now()
	for v := range p.progressChannel {
		if time.Now().Sub(ts) < p.interval {
			continue
		}
		p.print(v)
		ts = time.Now()
	}
	p.print(p.current)
	wg.Done()
}

func (p *ProgressBar) print(current int64) {
	progressBar := make([]byte, p.barLength)
	comm.FillBytes(progressBar[:int(p.current*int64(p.barLength)/p.maxValue)], '=')
	elapsed := time.Now().Sub(p.epoch)
	text := fmt.Sprintf("\r[%s] %d%% %s", string(progressBar), current*100/p.maxValue, comm.FormatTimeDuration(elapsed))
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

// SetSpeedCalculator sets speed calculator.
func (p *ProgressBar) SetSpeedCalculator(calc SpeedCalculator) {
	p.speedCalculator = calc
}

// SetProgress sets current progress.
func (p *ProgressBar) SetProgress(current int64) {
	if current < 0 {
		p.current = 0
	} else if current > p.maxValue {
		p.current = p.maxValue
	} else {
		p.current = current
	}
	p.progressChannel <- p.current
}

// AddProgress adds value to current progress.
func (p *ProgressBar) AddProgress(value int64) {
	p.SetProgress(p.current + value)
}
