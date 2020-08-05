package main

import "fmt"

func main() {
	stra := "the spice must flow"
	fmt.Println(stra)
	byts := []byte(stra)
	fmt.Println(byts)
	strb := string(byts)
	fmt.Println(strb)
}
