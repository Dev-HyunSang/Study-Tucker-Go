package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
