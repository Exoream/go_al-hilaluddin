package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	average := 0
	for _, value := range s.score {
		average += value
	}
	average /= len(s.score)
	return float64(average)
}

func (s Student) Min() (min int, name string) {
	name = s.name[0]
	min = s.score[0]
	for index, value := range s.score {
		if value < min {
			min = value
			name = s.name[index]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	name = s.name[0]
	max = s.score[0]
	for index, value := range s.score {
		if value > max {
			max = value
			name = s.name[index]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i+1) + " Student’s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + strconv.Itoa(i+1) + " Student’s Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")
}