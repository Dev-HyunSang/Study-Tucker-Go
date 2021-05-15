package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8
	B int
	C int8
	D int
	E int8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
