package main

import "fmt"

// Disini saya menggunakan kompleksitas O(√n)
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(13)) // true

	fmt.Println(primeNumber(17)) // true

	fmt.Println(primeNumber(20)) // false

	fmt.Println(primeNumber(35)) // false

	fmt.Println(primeNumber(1500450271)) // true
}