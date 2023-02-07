package main

import (
	"fmt"
)

func main() {

	// user := map[string]int{
	// 	"Das":  2,
	// 	"Bos":  4,
	// 	"Dudu": 5,
	// }

	//initial map
	user := make(map[string]int)

	user["Das"] = 5999
	user["Das"] = 1
	user["Das3"] = 2
	user["Bos"] = 3

	age, ex := user["Das"]
	fmt.Println(age, ex)
	if ex {
		fmt.Println("Das", ex)
	}
	fmt.Println(age, ex)

	fmt.Println("non user -----", len(user))

	age, ex = user["Huega"]
	if ex {
		fmt.Println("Huega", ex)
	}
	fmt.Println(age, ex)

	for key, value := range user {
		fmt.Println(key, value)
	}

	fmt.Println("Add -----", len(user))

	user["Huigarda"] = 1000
	user["Nobody"] = 1

	for key, value := range user {
		fmt.Println(key, value)
	}

	fmt.Println("Delete -----", len(user))

	delete(user, "Das")
	delete(user, "Bos")

	for key, value := range user {
		fmt.Println(key, value)
	}
}
