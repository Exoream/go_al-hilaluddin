package main

// Berikan tujuan dari struct ini dan jika mau diakses diluar package gunakan huruf kapital di awal variabelnya menjadi User
type user struct {
	id       int		// Jika mau diakses diluar package gunakan huruf kapital di awal variabelnya 
	username int		// Username diganti tipe data string, dan jika mau diakses diluar package gunakan huruf kapital di awal variabelnya
	password int		// Password diganti tipe data string dan jika ada password sebaiknya di enkripsi atau tidak secara langsung diberikan tipe data mentah. Selain itu jika mau diakses diluar package gunakan huruf kapital di awal variabelnya
}			

// Berikan tujuan dari struct ini apa ? dan jika mau diakses diluar package gunakan huruf kapital di awal variabelnya menjadi UserService
type userservice struct {		
	t []user		  	// Penamaan nama t itu apa? sebaiknya diganti dengan namaUser atau yang sesuai. Jika mau diakses diluar package gunakan huruf kapital di awal variabelnya
}

// Berikan tujuan dari function ini dan bisa juga variabel u nya di ganti "us" atau yang lebih jelas, agar tidak ada salah paham u == user
func (u userservice) getallusers() []user {   // Sebaiknya gunakan camelCase penamaannya pada getalluser menjadi getAllUser
	return u.t
}

// Berikan tujuan dari function ini dan bisa juga variabel u nya di ganti "us" atau yang lebih jelas, agar tidak ada salah paham u == user
func (u userservice) getuserbyid(id int) user {	 // Gunakan camelCase di getuserbyid menjadi getUserById
	// Berikan komentar perulangan ini untuk apa? Dari sini artinya perulangan ini bertujuan untuk mendapatkan data user berdasarkan id
	for _, r := range u.t {		// Nama variabel r itu diganti dengan yang nama yang sesuai seperti user atau yang sesuai
		if id == r.id {
			return r
		}
	}
	// Sebaiknya tambahkan kondisi jika id user tidak ditemukan
	return user{}
}

