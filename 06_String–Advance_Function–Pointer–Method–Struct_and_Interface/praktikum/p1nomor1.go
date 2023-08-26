package main

import "fmt"

type car struct {
	carType string
	fuelin float64
}

func (c car) fuelCar()  {
	const fuel = 1.5
	estDistance := c.fuelin * fuel
	fmt.Println("Nama Mobil : ", c.carType)
	fmt.Println("Estimasi Bensin : ", estDistance)
}

func main() {
	c := car {
		carType : "SUV",
		fuelin: 10.5,
	}
	c.fuelCar()
}