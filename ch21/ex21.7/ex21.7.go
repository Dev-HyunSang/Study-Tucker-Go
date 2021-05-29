package main

import (
	"fmt"
	"os"
)

type Writer func(string)

func WriteHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	WriteHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
