package main

import (
	"fmt"
	"math"
)


func Frog(jumps []int) int {
	minCost := make([]int, len(jumps))
	minCost[0] = 0
	minCost[1] = minCost[0] + int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < len(jumps); i++ {
		oneStep := minCost[i-1] + int(math.Abs(float64(jumps[i] - jumps[i-1])))
		twoStep := minCost[i-2] + int(math.Abs(float64(jumps[i] - jumps[i-2])))
		minCost[i] = min(oneStep, twoStep)
	}
	return minCost[len(jumps) - 1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}