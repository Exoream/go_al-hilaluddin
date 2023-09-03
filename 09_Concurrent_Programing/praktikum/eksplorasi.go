package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Category string `json:"category"`
}

func product(id int, c chan<- Product, wg sync.WaitGroup) {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)
	
	respon, err := http.Get(url)
	if err != nil {
		fmt.Println("Eror Product ID", id, err)
		return
	}

	var product Product
	if err := json.NewDecoder(respon.Body).Decode(&product); err != nil {
		fmt.Printf("Error decoding product %d: %v\n", id, err)
		return
	}
	c <- product
	respon.Body.Close()
	wg.Done()
}

func main() {
	id := []int{1, 2, 3, 4}

	var wg sync.WaitGroup
	c := make(chan Product, len(id))

	for _, id := range id {
		wg.Add(1)
		go func (id int) {
			product(id, c, wg)
		}(id)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for product := range c {
		fmt.Println("Product Data : ")
		fmt.Println("Title : ", product.Title)
		fmt.Println("Price : ", product.Price)
		fmt.Println("Category : ", product.Category)
		fmt.Println("=========")
		fmt.Println("=========")
	}
}
