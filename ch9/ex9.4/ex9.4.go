package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 increase")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 increase")
	}

	fmt.Println("cnt:", cnt)
}
