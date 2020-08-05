package main

import (
	"fmt"
)

func main() {
	var scores [10]int
	scores[0] = 339

	scores2 := [4]int{1, 2, 3, 4}
	fmt.Println(scores)
	fmt.Println(scores2)
	fmt.Println(len(scores))

	for index, value := range scores2 {
		fmt.Println("--->", index, value)
	}
}
