package main

import "fmt"

type Stringer interface {
	String() string
}

type Stdunet struct {
	Age int
}

func (s *Stdunet) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Stdunet)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Stdunet{15}

	PrintAge(s)
}
