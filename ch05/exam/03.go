package main

import "fmt"

func main()  {
	var a = 123
	var b int = 4567
	f := 3.14159269

	fmt.Println("%6d\n", a)
	fmt.Println("%06d\n", b)
	fmt.Println("%6.2f\n", f)
}
