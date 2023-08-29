package main

import "fmt"


func playingDomino(cards [][]int, deck []int) interface{} {
	for _, value := range cards {
		if value[0] == deck [0] || value[1] == deck [0] || value[0] == deck[1] || value[1] == deck[1] {
			return value
		}
	}
	return "Tutup Kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6, 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))              // "Tutup kartu"
}