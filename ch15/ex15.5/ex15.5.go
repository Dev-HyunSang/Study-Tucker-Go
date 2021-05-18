package main

import "fmt"

func main() {
	str := "Hello, World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
}
