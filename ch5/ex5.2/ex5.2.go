package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Println("%5d, %5d\n", a, b)
	fmt.Println("%05d, %05d\n", a, b)
	fmt.Println("%-5d, %-05d\n", a, b)

	fmt.Println("%5d, %5d\n", c, c)
	fmt.Println("%05d, %05d\n", c, c)
	fmt.Println("%-5d, %-05d\n", c, c)
}
