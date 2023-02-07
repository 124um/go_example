package main

import (
	"fmt"

)

func mainAnonim() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}