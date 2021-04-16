package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Println("%08.2f\n", a)
	fmt.Println("%08.2g\n", a)
	fmt.Println("$8.5g\n", a)
	fmt.Println("%f\n", c)
}
