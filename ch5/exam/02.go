package main

import "fmt"

func main() {
	var a int
	var b int

	// fmt.Scanln(a, b) / 기존 문제의 코드
	fmt.Scanln(&a, &b) // 문제를 해결한 후의 코드
	fmt.Println(a, b)
}
