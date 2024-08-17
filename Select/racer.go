package main

import (
	"fmt"
	"net/http"
	"time"
)

const defaultSleep = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, defaultSleep)
}

func ConfigurableRacer(a, b string, sleepDuration time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(sleepDuration):
		return "", fmt.Errorf("time limit exceeded")
	}
}

// func measureResponseTime(a string) time.Duration {
// 	startA := time.Now()
// 	http.Get(a)
// 	return time.Since(startA)
// }

// ping returns empty channel
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(chan struct{}) {
		http.Get(url)
		close(ch)
	}(ch)
	return ch
}
