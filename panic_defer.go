package main

import (
	"fmt"
)

func main() {

	defer handlePanic()
	sliceMmm := []string{
		"ewewewe1",
		"ewewewe2",
		"ewewewe3",
		"ewewewe4",
		"ewewewe5",
		"ewewewe6",
		"ewewewe7",
		"ewewewe7",
	}

	fmt.Println(sliceMmm)

	sliceMmm[20] = "panic"

	fmt.Println(sliceMmm)

}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
