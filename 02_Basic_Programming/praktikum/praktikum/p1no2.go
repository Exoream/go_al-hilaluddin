package praktikum

import "fmt"

func B() {
	var a int

	fmt.Print("Masukkan Nilai : ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("Nilai", a, "Merupakan Angka Genap")
	} else {
		fmt.Println("Nilai", a, "Merupakan Angka Ganjil")
	}

}
