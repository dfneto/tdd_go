package main

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(buffer *bytes.Buffer, name string) {
	fmt.Fprintf(buffer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
