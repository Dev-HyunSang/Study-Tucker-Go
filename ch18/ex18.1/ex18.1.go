package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("Slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
}

// 오류 발생, 슬라이스 길이를 초과함, 런 타임 에러
