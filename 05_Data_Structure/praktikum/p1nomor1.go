package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	makeMap := make(map[string]string)
	var mergeArray []string

	for _, list := range arrayA {
		makeMap[list] = " "
		mergeArray = append(mergeArray, list)
	}

	for _, list := range arrayB {
		if _, exist := makeMap[list]; !exist {
			mergeArray = append(mergeArray, list)
		}
	}
	return mergeArray	
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}