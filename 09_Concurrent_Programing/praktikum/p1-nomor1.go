package main

import (
	"fmt"
	"time"
)

func kelipatanNumber(panjang, x int) {
	for i := 0; i <= panjang; i += x {
		fmt.Println(i)
	}
}

func main() {
	go kelipatanNumber(30, 2)
	time.Sleep(3 * time.Second)
}
