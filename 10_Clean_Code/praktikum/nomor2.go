package main

import "fmt"

// Saya mengganti nama struct menjadi mobil dan jika mau diakses diluar package gunakan huruf kapital di awal variabelnya menjadi Mobil
// Disini code hanya digunakan di package main saja. Maka saya tidak usah membuat Pascal Case
type mobil struct {
	// Menghapus totalRoda karena tidak perlu
	kecepatanPerJam int		// Saya merubah nama variabel menjadi camelCase 
}

// Saya menghapus struct mobil karena tidak perlu

// Function ini bertujuan untuk mendapatkan nilai kecepatan dari mobil
func (m *mobil) berjalan(kecepatanBaru int) {	// Saya menambahkan kecepatanBaru int di method berjalan	
	m.tambahKecepatan(kecepatanBaru)			// Maka dari itu nilai 10 tadi diganti dengan kecepatan
}

// Function ini bertujuan untuk menambah kecepatan mobil
func (m *mobil) tambahKecepatan(kecepatanBaru int) {		// Saya membuat nama method menjadi camelCase
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	mobilCepat := mobil{}		// Mengganti variabel mobilcepat menjadi camelCase
	mobilCepat.berjalan(15)		// Memberikan nilai kecepatan di mobil berjalan sebesar 15
	mobilCepat.berjalan(10)		// Memberikan nilai kecepatan di mobil berjalan sebesar 10 
	mobilCepat.berjalan(5)		// Memberikan nilai kecepatan di mobil berjalan sebesar 5

	mobilLamban := mobil{}		// Mengganti variabel mobillamban menjadi camelCase
	mobilLamban.berjalan(5)		// Memberikan nilai kecepatan di mobil berjalan sebesar 5
	
	fmt.Println(mobilCepat)
	fmt.Println(mobilLamban)
}