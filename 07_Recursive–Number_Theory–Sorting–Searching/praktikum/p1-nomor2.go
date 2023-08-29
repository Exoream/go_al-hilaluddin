package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	count := make(map[string]int)
	for _, value := range items {
		count[value]++
	}
	
	var pairs []pair
	for name, count := range count {
		pairs = append(pairs, pair{name, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})
	
	return pairs
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}