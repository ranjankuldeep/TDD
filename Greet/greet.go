package main

import (
	"fmt"
	"io"
)

// This is refer to as dependecy injection.
func Greet(writer io.Writer, str string) {
	fmt.Fprintf(writer, "hey %s", str)
}
