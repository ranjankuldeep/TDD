package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "cheers")

	got := buf.String()
	want := "hey cheers"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
