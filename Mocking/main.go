package main

import (
	"os"
	"time"
)

func main() {
	sleeper := ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	CountDown(os.Stdout, &sleeper)
}
