package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type spyCountDownOperation struct {
	calls []string
}

func (s *spyCountDownOperation) Sleep() {
	s.calls = append(s.calls, sleep)
}
func (s *spyCountDownOperation) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

// Implements Sleeper interface
type spyTime struct {
	durationSlept time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("Print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &spyCountDownOperation{}
		CountDown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep befor every print", func(t *testing.T) {
		spySleepPrinter := &spyCountDownOperation{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := spyTime{}

	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
