package praktikum

import "fmt"

func F() {
	var a int

	fmt.Print("Masukkan Angka : ")
	fmt.Scanln(&a)

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Println(i)
		}
	}
}
