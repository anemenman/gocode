package main

import (
	"fmt"
)

type A struct {
	Name string
}

func (a *A) a() {
	fmt.Println(a.Name)
}

type St struct {
	*A
	Power int
}

func main() {
	g := &St{
		A:     &A{Name: "Goku"},
		Power: 0,
	}

	g.a()
}
