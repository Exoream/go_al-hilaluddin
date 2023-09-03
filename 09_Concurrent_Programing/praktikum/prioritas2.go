package main

import (
	"fmt"
	"sync"
)


func main() {
	input := "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	data := make(map[rune]int)

	var wg sync.WaitGroup
	var m sync.Mutex

	for _, value := range input {
		wg.Add(1)
		go func(text rune) {
			m.Lock()
			data[text]++
			m.Unlock()
			wg.Done()
		}(value)
	}
	wg.Wait()

	for value, jumlah := range data {
		fmt.Printf("%c: %d\n", value, jumlah)
	} 
}