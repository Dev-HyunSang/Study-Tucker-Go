package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Println(str)
	fmt.Println("%s", str)
	fmt.Println("%q", str)
}
