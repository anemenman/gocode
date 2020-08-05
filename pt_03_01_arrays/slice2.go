package main

import (
	"fmt"
)

func main() {

	var name []string

	name = append(name, "A")
	fmt.Println(name)
	name = append(name, "A")
	fmt.Println(name)
	name = append(name, "B")
	fmt.Println(name)
}
