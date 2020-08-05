package main

import (
	"errors"
	"fmt"
)

func process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}

func main() {
	fmt.Println(process(1))
	fmt.Println(process(2))
	fmt.Println(process(0))
}
