package main

import (
	"fmt"
)

func main() {

	lookup := make(map[string]int)
	lookup["Goku"] = 11
	lookup["AA"] = 12
	lookup["kkk"] = 13
	v, exists := lookup["Goku"]

	fmt.Println(v, exists)

	fmt.Println(len(lookup))

	delete(lookup, "Goku")
	delete(lookup, "bla...")

	fmt.Println(len(lookup))

	fmt.Println("---------")
	for k, v := range lookup {
		fmt.Println(k, v)
	}
}
