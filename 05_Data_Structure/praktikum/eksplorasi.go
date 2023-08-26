package main

import (
	"fmt"
	"math"
)

func main() {
	matriks := [3][3]int {
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	// var result int
	// var result2 int
	// for i := 0; i < len(matriks); i++ {
	// 	result += matriks[i][i]
	// }

	// for i := 0; i < len(matriks); i++ {
	// 	result2 += matriks[i][len(matriks)-1-i]
	// }

	// absulote := math.Abs(float64(result-result2))
	// fmt.Println("Nilai Akhir Penjumlahan Adalah : ", absulote)
	
	var result int
	var result2 int
	for index, _ := range matriks {
		result += matriks[index][index]
		result2 += matriks[index][len(matriks)-1-index]
	}

	fmt.Println(result)
	fmt.Println(result2)
	absulote := math.Abs(float64(result-result2))
	fmt.Println("Nilai Akhir Adalah :", absulote)
}