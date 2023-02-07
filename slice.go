package main

import (
	"fmt"
)

func mainSlice() {

	arrayMess := make([]string, 2)
	arrayMess = append(arrayMess, "66")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))
	arrayMess = append(arrayMess, "11")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))
	arrayMess = append(arrayMess, "1111")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))
	arrayMess = append(arrayMess, "222")
	arrayMess = append(arrayMess, "222")
	arrayMess = append(arrayMess, "222")
	arrayMess = append(arrayMess, "222")
	arrayMess = append(arrayMess, "222")
	arrayMess = append(arrayMess, "222")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))
	arrayMess = append(arrayMess, "33")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))
	arrayMess = append(arrayMess, "11")
	fmt.Println(arrayMess)
	fmt.Println(len(arrayMess))
	fmt.Println(cap(arrayMess))

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

	for index, value := range sliceMmm {
		fmt.Println(index, value)
	}
	for _, value := range sliceMmm {
		fmt.Println(value)
	}
	for index, _ := range sliceMmm {
		fmt.Println(index)
	}
	for index, _ := range sliceMmm {
		if index == 5 {
			break //or return
		}
		fmt.Println(index)
	}
}
