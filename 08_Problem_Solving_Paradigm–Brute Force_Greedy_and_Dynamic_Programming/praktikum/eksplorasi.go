package main

import (
	"fmt"
	"strings"
)

func numberToRomawi(number int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	
	var result strings.Builder
	for i := range value {
		for number >= value[i] {
			result.WriteString(romawi[i])
			number -= value[i]
		}
	}
	return result.String()
}

func main() {
	fmt.Println(numberToRomawi(4))
	fmt.Println(numberToRomawi(9))
	fmt.Println(numberToRomawi(23))
	fmt.Println(numberToRomawi(2021))
	fmt.Println(numberToRomawi(1646))
}