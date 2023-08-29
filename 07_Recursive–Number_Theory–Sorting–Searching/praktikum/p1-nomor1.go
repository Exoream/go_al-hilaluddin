package main

import "fmt"

func fibonacci(number int) int {
	number1 := 0
	number2 := 1
	number3 := 0
	for i := 2; i <= number; i++ {
		number3 = number1 + number2
		number1 = number2
		number2 = number3
	}
	return number3
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}