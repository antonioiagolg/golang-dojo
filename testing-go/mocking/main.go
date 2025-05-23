package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
    Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func Countdown(w io.Writer, s Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
    cs := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
 	Countdown(os.Stdout, cs)
}
