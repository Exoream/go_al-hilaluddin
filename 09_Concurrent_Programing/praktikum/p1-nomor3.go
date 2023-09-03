package main

import (
	"fmt"
	"time"
)

func kelipatanTiga2(c chan int) {
	result := 0
	for i := 3; ; i += 3 {
		result += i
		c <- result
	}
}

func main() {
	c := make(chan int, 30)
	go kelipatanTiga2(c)
	
	for i := 0; i < 30; i++ {
		result := <-c
		fmt.Println(result)
		time.Sleep(100 * time.Millisecond)
	}
}
