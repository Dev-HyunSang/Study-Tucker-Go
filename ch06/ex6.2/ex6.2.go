package main

import "fmt"

func main() {
	var x1 int8 = 34
	var x2 int16 = 34
	var x3 int8 = 34
	var x4 uint16 = 34

	fmt.Println("^%d = %5d,\t %08b\n", x1, ^x1, uint(^x1))
	fmt.Println("^%d = %5d,\t %016b\n", x2, ^x2, uint(^x2))
	fmt.Println("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
	fmt.Println("^%d = %5d, \t %16b\n", x4, ^x4, ^x4)
}
