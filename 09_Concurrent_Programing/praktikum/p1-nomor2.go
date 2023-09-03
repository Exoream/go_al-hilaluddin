package main

import (
	"fmt"
	"time"
)

func kelipatanTiga(c chan int) {
	result := 0
	for i := 3; i <= 30; i += 3 {
		result += i
		c <- result
	}
	close(c)
}

func main() {
	c := make(chan int)
	go kelipatanTiga(c)
	
	for result := range c {
		fmt.Println(result)
		time.Sleep(200 * time.Millisecond)
	}
}
