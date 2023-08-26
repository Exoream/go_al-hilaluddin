package main

import "fmt"

// Disini saya menggunakan kompleksitas O(log n)
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		result := pow(x, n/2)
		return result * result
	} else {
		result := pow(x, (n-1)/2)
		return result * result * x
	} 
}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(5, 3)) // 125

	fmt.Println(pow(10, 2)) // 100

	fmt.Println(pow(2, 5)) // 32

	fmt.Println(pow(7, 3)) // 343

}