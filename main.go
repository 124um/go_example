package main

import "fmt"

var msg string

func init() {
	msg = "init msg before main"
}

func main() {
	fmt.Println(msg)
	
	defer deferFufu()

	func() {
		fmt.Println("Anonim func")
	}()

	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	// var number, secondNumber int

	// number = 15
	// secondNumber = 100
	// if number < secondNumber {
	// 	println("первое число меньше второго")
	//   } else if number > secondNumber {
	// 	println("второе число меньше первого")
	//   } else {
	// 	println("числа равны")
	//   }

	// switch secondNumber {
	// case 100:
	//   println("число равно 100")
	// case 200:
	//   println("число равно 200") // будет выведено это
	// }

	// switch {
	// case number < secondNumber:
	//   println("первое число меньше второго")
	// case number > secondNumber:
	//   println("второе число меньше первого")
	// default:
	//   println("числа равны")
	// }

	// for i := 1; i <= 10; i++ {
	// 	println(i)
	// }

	// for number < secondNumber {
	// 	number = number + 1
	// 	println(number)
	// }

	// i := 0
	// for {
	// 	println("Бесконечный цикл")
	// 	if i == 10 {
	// 		break
	// 	}
	// 	i++
	// }

	// var numbers [3]int
	// numbers[0] = 10
	// numbers[1] = 11
	// numbers[2] = 12

	// var numbers_array [3]int = [3]int{10, 11, 12}
	// numbers := [3]int{10, 11, 12} // массив
	numbers := []int{10, 11, 12}      // слайс
	numbers = append(numbers, 13)     // добавление одного элемента
	numbers = append(numbers, 14, 15) // добавление двух элементов
	// println(numbers[0])

	// for i := 0; i < len( numbers); i++ {
	// 	print_a(numbers[i])
	// }

	array_print(numbers)
}

func print_a(a int) int {
	println("=====> ", a)
	return a
}

func array_print(param []int) []int {
	for i := 0; i < len(param); i++ {
		print_a(param[i])
	}
	return param
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = 1
		}
	}
	return min
}

func deferFufu() {
	fmt.Println("Run Defer func!!")
}
