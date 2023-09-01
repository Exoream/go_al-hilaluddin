package main

import "fmt"

func fibonacci(number int) int {
	memoization := make(map[int]int)			
	if value, exist := memoization[number]; exist { 	// Menggabungkan dengan teknik memoization
		return value
	}												
	// Jika sudah ada nilai fibonacci dalam map, kita tidak perlu melakukan perulangan lagi

	if number == 0 {
		memoization[number] = 0
        return 0
    } else if number == 1 {
		memoization[number] = 1
        return 1
    }
	memoization[number] = fibonacci(number-1) + fibonacci(number-2)
    return memoization[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}