package main

import "fmt"

func MaxSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	currentSum := arr[0]
	resultSum := arr[0]

	for _, value := range arr[1:] {
		// Versi 1
		// currentSum = max(value, (currentSum + value))
		// resultSum = max(resultSum, currentSum) 

		// Versi 2
		hasil := currentSum + value
		if value > hasil {
			currentSum = value
		} else {
			currentSum = hasil
		}

		if currentSum > resultSum {
			resultSum = currentSum 
		} 
	}
	return resultSum
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
 	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}