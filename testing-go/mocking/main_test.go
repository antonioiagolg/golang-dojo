package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	Calls int
}

type SpyTimer struct {
	timeSlept time.Duration
}

func (st *SpyTimer) Sleep(t time.Duration) {
	st.timeSlept = t
}

type CountdownOrderExecution struct {
	Calls []string
}

func (coe *CountdownOrderExecution) Write(b []byte) (int, error) {
	coe.Calls = append(coe.Calls, write)
	return 0, nil
}

func (coe *CountdownOrderExecution) Sleep() {
	coe.Calls = append(coe.Calls, sleep)
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("Checking the correct number of call on sleep", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)
		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if sleeper.Calls != 3 {
			t.Errorf("got %d but want 3 calls", sleeper.Calls)
		}
	})

    t.Run("Checking the order of calls", func(t *testing.T) {
        coe := &CountdownOrderExecution{}
        Countdown(coe, coe)
        want := []string{
            "write", // 3
            "sleep",
            "write", // 2
            "sleep",
            "write", // 1
            "sleep",
            "write", // Go!
        }

        if !reflect.DeepEqual(want, coe.Calls) {
            t.Error("Wrong call sequence")
        }
    })

}

func TestConfigurableSleep(t *testing.T) {
	timeSlept := 5 * time.Second
	spyTimer := &SpyTimer{}
	
	configSleeper := &ConfigurableSleeper{duration: timeSlept, sleep: spyTimer.Sleep}
	configSleeper.Sleep()

	if spyTimer.timeSlept != timeSlept {
		t.Errorf("Wrong time slept")
	}

}
