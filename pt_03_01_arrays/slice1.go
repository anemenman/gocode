package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 3, 4}
	fmt.Println(scores)

	scores2 := make([]int, 10)
	fmt.Println(scores2)

	scores3 := make([]int, 0, 10)
	fmt.Println(scores3)
	//scores3[5] = 999
	//fmt.Println(scores3)

	scores3 = append(scores3, 5)
	fmt.Println(scores3)

	c := cap(scores3)
	fmt.Println("c--->", c)

	scores4 := make([]int, 0, 20)
	scores4 = scores4[0:6]
	fmt.Println(scores4)
}
