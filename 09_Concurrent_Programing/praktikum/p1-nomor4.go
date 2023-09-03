package main

import (
	"fmt"
	"sync"
)

type Data struct {
	data []int
	m sync.Mutex
}

func (d *Data) add(nilai int) {
	d.m.Lock()
	d.data = append(d.data, nilai)
	d.m.Unlock()
}

func main() {
	push := Data {
		data : make([]int, 0),
	}

	push.add(5)
	push.add(10)
	push.add(15)

	fmt.Println("Nilai : ", push.data)
}