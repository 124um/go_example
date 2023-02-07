package main

import (
	"fmt"

	// "github.com/golang/protobuf/ptypes/any"
)

type Number interface{
	int | float32
}

type User struct{
	email string
	name string
	age int
}


func mainGen() {

	a:= []int{2,3,4,5}
	b:= []float32{12.22,2.55,45.2334}
	str := []string{"1221aww", "dfdf44", "f", "gg"}
	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(searchElement(str, "f"))


	us := []User{
		{
			email: "fdkfdkfjdkgf.com",
			name: "Fuka",
			age: 33,
		},
		{
			email: "111.com",
			name: "Buka",
			age: 55,
		},
		{
			email: "2222.com",
			name: "Geka",
			age: 44,
		},
		{
			email: "333.com",
			name: "Chuka",
			age: 20,
		},
	}


	fmt.Println("-----", searchElement(us, User{email: "333.com",
	name: "Chuka",
	age: 20,}))

	printAny(User{email: "333.com",
	name: "Chuka",
	age: 20,})

	printAny(us)

}

func sum[V Number](array []V) V{
	var result V
	for _, num := range array {
		result += num
	}

	return result
}

func searchElement[C comparable](elements []C, searchEle C) bool  {
	for _, ele := range elements {
		if ele == searchEle {
			return true
		}
	}
	return false
}

func printAny(arr any){
	fmt.Println(arr)
}