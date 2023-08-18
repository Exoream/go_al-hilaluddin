package praktikum

import "fmt"

func A() {
	var a, b, t float32

	fmt.Print("Masukkan Sisi A : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan Sisi B : ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan Tinggi : ")
	fmt.Scanln(&t)
	l := 0.5 * (a + b) * t

	fmt.Println("Luas Trapesium Adalah : ", l)

}
