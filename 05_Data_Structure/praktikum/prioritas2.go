package main

import "fmt"

func PairSum(arr []int, target int) []int {
	number := make(map[int]int)

	for index, value := range arr {
		result := target - value
		if index2, value2 := number[result]; value2 {
			return []int{index2, index}
		}
		number[value] = index
	}
	return nil
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}