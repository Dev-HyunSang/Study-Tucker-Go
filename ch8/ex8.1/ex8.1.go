package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	C = 20 // Error -  상수는 대입문 좌변에 올 수 없음.
	fmt.Println(&C)
}
