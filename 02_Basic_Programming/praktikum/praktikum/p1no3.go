package praktikum

import "fmt"

func C() {
	var a int
	fmt.Print("Masukkan Nilai : ")
	fmt.Scan(&a)

	if a >= 80 && a <= 100 {
		fmt.Println("Anda Mendapatkan Nilai A")
	}else if a >= 65 && a <= 79 {
		fmt.Println("Anda Mendapatkan Nilai B")
	}else if a >= 50 && a <= 64 {
		fmt.Println("Anda Mendapatkan Nilai C")
	}else if a >= 35 && a <= 49 {
		fmt.Println("Anda Mendapatkan Nilai D")
	}else if a >= 0 && a <= 34 {
		fmt.Println("Anda Mendapatkan Nilai E")
	}else{
		fmt.Println("Nilai Invalid")
	}
}
