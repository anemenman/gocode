package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randArray []int
	for i := 0; i < 100000; i++ {
		randArray = append(randArray, rand.Intn(100))
	}

	fmt.Println(randArray)
}
