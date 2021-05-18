package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n"

	// 백쿼트로 묶으면 특수 문자가 동작하지 않습니다!
	str2 := `Go is "awsome!"!\nGo is simple and\t'powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}
