# 17 ORM & Code Structure (MVC)

## Object Relational Mapping (ORM)
ORM adalah singkatan dari "Object-Relational Mapping." Ini adalah sebuah konsep dalam pemrograman yang digunakan untuk menghubungkan dan mengelola data dalam basis data relasional (seperti MySQL, PostgreSQL, atau Oracle) menggunakan objek dalam bahasa pemrograman. GORM adalah satu satu ORM pada bahasa golang.

## Keuntungan ORM

```
~ Mengurangi query berulang
~ Kode yang lebih mudah dipahami
~ perubahan struktur basis data yang mudah dikelola
~ Secara otomatis mengambil data ke objek siap pakai
~ Mudah melakukan penyaringan atau pemfilteran data sebelum menyimpan didatabase
~ Beberapa memiliki cache query
~ Menghindari typo jika menulis query manual

```

## Kekurangan ORM

```
~ Terkadang menambahkan query yang tidak perlu sehingga tidak efisien 
~ Menjalankan database relasional yang tidak perlu
~ Query yang kompleks mengakibatkan penulisan yang panjang
~ Fungsi sql tertentu yang terkait dengan satu vendor mungkin tidak didukung

```
## Database Migration
Database migration adalah proses untuk mengelola perubahan skema basis data seiring waktu. Ini melibatkan mengubah struktur basis data, seperti menambahkan tabel baru, menghapus tabel, atau memodifikasi kolom dalam tabel yang ada.

## Alasan Penggunaan Database Migration

```
~ Mengelola Perubahan Skema
~ Integritas Data
~ Pelacakan Perubahan
~ Manajemen Perubahan
~ Dokumentasi History
~ Selalu Kompatibel dengan Perubahan Versi Aplikasi

```

## DB Relational Di GORM

```
~ Belongs to
~ Has One
~ Has Many
~ Many to Many

```

## DB Transaksi di GORM
Transaksi database menjalankan beberapa perintah SQL dalam satu transaksi yang dapat di-rollback jika terjadi kesalahan. Transaksi database adalah cara yang baik untuk memastikan konsistensi data saat beberapa operasi database harus dilakukan secara atomik.

## Instal GORM

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

```

## Cara Koneksi Ke Database

```
var DB *gorm.DB

// Membuat Koneksi Ke Database
func InitDB() {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "qwerty123",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "db_km5_gorm",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// Membuat Auto Migrate
func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}

```

## Models

```
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

```

## Controller Get All Users

```
// get all users
func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

```

## Controller Get User By Id

```

// get user by id
func GetUserController(c echo.Context) error {
	id := c.Param("id")
	idUsers, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error convert",
			"status":  "failed",
		})
	}

	var users User
	result := DB.First(&users, idUsers)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"status":  "failed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user found",
		"status":  "success",
		"user":    users,
	})
}

```

## Controller Create New User

```
// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

```

### Controller Delete User

```

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	idUsers, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error convert",
			"status":  "failed",
		})
	}

	var users User
	result := DB.First(&users, idUsers)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"status":  "failed",
		})
	}
	DB.Delete(&users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted",
		"status":  "success",
	})
}

```

## Controller Update User

```

// update user by id
func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	idUsers, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error convert",
			"status":  "failed",
		})
	}

	var users User
	result := DB.First(&users, idUsers)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
			"status":  "failed",
		})
	}

	updatedUser := new(User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error binding data",
			"status":  "failed",
		})
	}

	DB.Model(&users).Updates(updatedUser)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user updated",
		"status":  "success",
		"data" : users,
	})
}

```

## Routing

```
func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

```

## Code Structure MVC
Structure code dalam pola arsitektur MVC (Model-View-Controller) adalah cara untuk mengatur kode dalam suatu aplikasi sehingga terbagi menjadi tiga komponen utama: Model, View, dan Controller. Setiap komponen memiliki tanggung jawab yang berbeda dalam aplikasi. Hal ini membantu dalam pengorganisasian, pemahaman, dan pemeliharaan kode yang lebih baik, mengurangi redundansi, memungkinkan kolaborasi tim yang lebih efisien, dan mempermudah pengujian serta skalabilitas proyek. Dengan kata lain, struktur membantu menjadikan pengembangan perangkat lunak lebih teratur dan terkelola.

## Structure Folder MVC Secara Umum

```
~ config
Tempat menyimpan berkas konfigurasi aplikasi seperti pengaturan database, konfigurasi server, atau konfigurasi lainnya.

~ controllers
Direktori ini berisi file-controller yang bertanggung jawab untuk menangani permintaan dari pengguna, menjalankan logika aplikasi, dan mengirimkan data ke tampilan. Biasanya, setiap rute atau endpoint URL memiliki controller yang sesuai.

~ models
Ini adalah tempat untuk mendefinisikan model-data yang berhubungan dengan basis data. Model digunakan untuk mengakses, mengubah, dan memanipulasi data dalam basis data.

~ routes
Direktori ini berisi definisi rute atau endpoint URL untuk aplikasi Anda. Rute menghubungkan permintaan pengguna dengan controller yang sesuai.

```



