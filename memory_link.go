package main

import (
	"fmt"
)


func mainMmm() {
	message := "i will ninjs" 
	printMessage(&message)
	fmt.Println(message)
	printMessage(&message)
	fmt.Println(message)
	printMessage(&message)
	fmt.Println(message)


	// links
	numInt := 10
	var p *int

	fmt.Println(p)
	fmt.Println(numInt)

	p = &numInt

	*p =  999
	fmt.Println(numInt)
}

func printMessage(message *string) {
	*message += "add in func printMessage ()" 
}