package main

import (
	"fmt"
)

func binerRepresentation (n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = i
	}
	return ans
}


func main() {
	var n int
	fmt.Printf("Masukkan Nilai N : ")
	fmt.Scan(&n)

	result := fmt.Sprintf("%b", binerRepresentation(n))
	fmt.Println("Ouput : ", result)
}