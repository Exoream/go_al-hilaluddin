package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 1000; x++ {
		for y := 1; y <= 1000; y++ {
		z := int(math.Sqrt(float64(c - x*x - y*y)))
		
		if (x+y+z) == a && (x * y * z) == b && (x*x + y*y + z*z) == c {
			fmt.Println(x, y, z)
			return
		}
	}
}
	fmt.Println("No Solution")	
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}