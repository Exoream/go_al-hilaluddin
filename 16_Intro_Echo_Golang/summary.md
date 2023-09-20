# 16 Intro Echo Golang


## Third Party Library Golang
Dalam bahasa pemrograman Go, "third-party library" merujuk kepada perpustakaan (library) yang dikembangkan oleh pihak ketiga atau bukan oleh pengembang bahasa Go sendiri. Penggunaan library ini membantu menghemat waktu dan upaya pengembangan, mempercepat proses pengembangan, dan memungkinkan penggunaan fungsionalitas yang matang yang telah dikembangkan oleh komunitas yang lebih luas. Beberapa contoh third party library yaitu Echo, Go Kit, Logrus, gorm.io, cobra dan masih banyak lainnya.


## Framework Echo
Echo merupakan salah satu framework web yang digunakan pada bahasa golang. Framework Echo dikenal karena kesederhanaannya, performa tinggi, dan minimalisme. Ini menyediakan fitur untuk routing, penanganan middleware, rendering data, skalabilitas, dan data binding sehingga cocok untuk mengembangkan RESTful API dan aplikasi web.


## Cara instal Echo

```
cd ~/project_name
go mod init myapp
go get github.com/labstack/echo/v4

```

## Basic Routing dan Controller Echo
Routing dalam konteks pengembangan web adalah proses menentukan bagaimana aplikasi web menangani permintaan HTTP yang masuk. Kontroler adalah bagian dari aplikasi web yang bertanggung jawab untuk menangani logika bisnis dan memproses permintaan HTTP yang masuk. Contoh dari Routing dan Controller

```
package main
// import Echo dan net/http
import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main(){
    // Membuat Instance Echo
    e := echo.New()

    // Membuat route dari controller UserController
    e.GET("/", UserController)

    // Memulai Server di Port 8080
    e.Start(":8080")
}

// User Controller
func UserController(c echo.Context) error {
    return c.String(http.StatusOk, "Budi")
}

```

## Data Rendering Format JSON
Data rendering adalah proses mengambil data dari aplikasi dan mengubahnya menjadi bentuk yang sesuai untuk ditampilkan kepada pengguna akhir melalui antarmuka pengguna. Dalam kasus ini kita mengubah ke format JSON. Contohnya

```
// import Echo dan net/http
import (
    "net/http"

    "github.com/labstack/echo/v4"
)

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main() {
    // Membuat instance Echo
    e := echo.New()

    // Menentukan route dari controller getDataHandler
    e.GET("/api/data", getDataHandler)

    // Memulai server pada port 8080
    e.Start(":8080")
}

// Controller untuk route "/api/data"
func getDataHandler(c echo.Context) error {
    data := map[string]interface{}{
        "message": "Ini adalah data dalam format JSON",
        "status":  200,
    }
    return c.JSON(http.StatusOK, data)
}

```

## Data Retrive Url Params
Data retrieval from URL params (pengambilan data dari parameter URL) adalah proses mengambil data atau nilai yang terdapat dalam URL sebuah permintaan HTTP. Contohnya

```
package main

import (
    "net/http"
    "strconv"
    "github.com/labstack/echo/v4"
)

// Struktur untuk data pengguna
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    e := echo.New()

    // Menentukan rute dengan parameter URL ":id" dan controller
    e.GET("/user/:id", func(c echo.Context) error {
        // Mengambil data dari parameter URL ":id"
        userIDStr := c.Param("id")

        // Mengonversi userIDStr ke tipe data integer
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            return c.String(http.StatusBadRequest, "ID pengguna tidak valid")
        }

        // Membuat instansi User
        user := User{
            ID:   userID,
            Name: "John Doe", 
        }

        // Mengembalikan nilai user dalam format JSON
        return c.JSON(http.StatusOK, user)
    })

    // Membuka port server
    e.Start(":8080")
}

```

## Data Retrive Query Param
Data retrieval from query parameters (pengambilan data dari parameter query) adalah proses mengambil data atau nilai yang terdapat dalam URL sebagai bagian dari query string dalam sebuah permintaan HTTP. Contohnya

```

package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    // Menentukan rute dengan query parameter "search"
    e.GET("/search", func(c echo.Context) error {
        // Mengambil data dari query parameter "search"
        searchTerm := c.QueryParam("search")

        // Lakukan logika berdasarkan data yang diambil
        // Misalnya, tampilkan hasil pencarian dalam respons
        return c.String(http.StatusOK, "Hasil pencarian untuk: "+searchTerm)
    })

    e.Start(":8080")
}

```

## Data Retrive Form Value
Data retrieval from form values (pengambilan data dari nilai formulir) adalah proses mengambil data yang dikirimkan melalui formulir HTML dalam permintaan HTTP. Form values adalah cara umum untuk mengirimkan data dari sisi klien (peminta) ke sisi server (penyedia) dalam aplikasi web. Contohnya

```

package main

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
)

type User struct {
    ID   int    `json:"id" form:"id"`
    Name string `json:"name" form:"name"`
    Age  int    `json:"age" form:"age"`
}

func createUser(c echo.Context) error {
    // Membaca data pengguna dari form value
    name := c.FormValue("name")
    ageStr := c.FormValue("age")

    // Mengonversi ageStr ke tipe data integer
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return c.String(http.StatusBadRequest, "Usia tidak valid")
    }

    // Membuat ID unik untuk pengguna baru
    newUserID := len(users) + 1

    // Membuat pengguna baru
    newUser := User{
        ID:   newUserID,
        Name: name,
        Age:  age,
    }

    // Menambahkan pengguna baru ke daftar pengguna
    users = append(users, newUser)

    // Mengembalikan data pengguna yang baru dibuat
    return c.JSON(http.StatusCreated, newUser)
}

func main() {
    e := echo.New()

    // Menangani data pengguna yang dikirimkan melalui form value
    e.POST("/createuser", createUser)

    e.Start(":8080")
}

```

## Binding Data
Binding data dalam konteks pengembangan web adalah proses menghubungkan atau mengikat data yang diterima dalam permintaan HTTP ke variabel atau struktur data yang dapat digunakan dalam kode aplikasi. Contohnya

```
package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type User struct {
    Name  string `json:"name" form:"name"`
    Email string `json:"email" form:"email"`
}

func main() {
    e := echo.New()

    e.POST("/createuser", func(c echo.Context) error {
        // Membaca data pengguna dari form value dan mengikatnya ke dalam struktur User
        var newUser User
        if err := c.Bind(&newUser); err != nil {
            return c.String(http.StatusBadRequest, "Data tidak valid")
        }

        // Sekarang, newUser berisi data dari form value yang dikirimkan oleh pengguna

        // Contoh: Mengembalikan data pengguna dalam format JSON
        return c.JSON(http.StatusCreated, newUser)
    })

    e.Start(":8080")
}

```
 

