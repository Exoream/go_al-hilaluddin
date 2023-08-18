package praktikum

import "fmt"

func E() {
	a := 5

	for i := 1; i <= a; i++ {
		for j := i; j <= a; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

