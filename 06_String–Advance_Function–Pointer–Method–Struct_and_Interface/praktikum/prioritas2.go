package main

import "fmt"

func caesar(offset int, input string) string {
	resString := ""
	for _, value := range input {
		resAscii := int(value)
		resAscii += offset
		if value >= 'a' && value <= 'z' {
			resAscii = ((resAscii - int('a')) % 26) + int('a')
		}
		resString += string(resAscii)
	}
	return resString
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cncv
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}