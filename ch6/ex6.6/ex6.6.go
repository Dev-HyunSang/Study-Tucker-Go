package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Println("%f + %f == $f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
