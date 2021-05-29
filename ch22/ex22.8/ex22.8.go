package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	m := [M]int{}

	m[hash(23)] = 10
	m[hash(259)] = 50
	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])
}
