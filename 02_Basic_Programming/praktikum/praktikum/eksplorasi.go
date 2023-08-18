package praktikum

import (
	"bufio"
	"fmt"
	"os"
)

func check(a string) bool {
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

func G() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Apakah Palindrome ? ")
	fmt.Print("Masukkan Kata : ")
	scanner.Scan()
	fmt.Print("Captured : ")
	text := scanner.Text()
	fmt.Println(text)

	if check(text) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}

}
