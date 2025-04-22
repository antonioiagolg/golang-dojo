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

type DefaultSleeper struct {
    timeInSeconds int
}

func (ds *DefaultSleeper) Sleep() {
    time.Sleep(time.Second * time.Duration(ds.timeInSeconds))
}

func Countdown(w io.Writer, s Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
    ds := &DefaultSleeper{timeInSeconds: 1}
	Countdown(os.Stdout, ds)
}
