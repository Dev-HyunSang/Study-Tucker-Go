package main

type Stringer interface {
	String() string
}

type Student struct{}

func main() {
	var stringer Stringer
	stringer.(*Student)
}
