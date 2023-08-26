package main

import "fmt"

func Mapping(slice []string) map[string]int {
	countSlice := make(map[string]int)

	for _, value := range slice {
		countSlice[value]++
	}
	return countSlice
}

func main() {
	// Test cases
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}