# 18 Middleware

## Middleware
Middleware adalah perangkat lunak perantara yang berfungsi sebagai perantara antara perangkat lunak atau komponen yang berbeda dalam sistem perangkat lunak. Middleware digunakan untuk mengelola, memproses, atau memodifikasi permintaan dan respons dalam arsitektur perangkat lunak yang terdistribusi atau berbasis jaringan.

## Instal JWT Echo 

```
go get github.com/labstack/echo-jwt

```

## Instal JWT Golang

```
go get github.com/golang-jwt/jwt

```
## Implementasi Middleware

```
~ Log Middleware
Log middleware digunakan untuk mencatat informasi tentang permintaan dan respons dalam aplikasi.

~ Session Middleware
Middleware untuk mengelola dan mempertahankan data sesi pengguna dalam aplikasi web.

~ Auth Middleware
Middleware untuk mengontrol akses ke sumber daya berdasarkan izin dan peran pengguna.

~ Custom Middleware
Middleware yang Anda buat khusus untuk tugas-tugas kustom dalam aplikasi Anda yang tidak tercakup oleh middleware bawaan atau yang disediakan oleh framework.

```

## Contoh Third Party Middleware

```
~ Negroni
~ Echo
~ Interpose
~ Alice

```

## Setup Middleware Echo
## Echo#Pre()
Metode echo.Pre() digunakan untuk mendaftarkan middleware pada tingkat root (sebelum rute-rute aplikasi) dan akan dieksekusi sebelum middleware yang mendaftar menggunakan echo.Use(). Middleware yang didaftarkan dengan echo.Pre() biasanya digunakan untuk tugas yang memerlukan eksekusi sebelum proses rute, seperti penanganan otentikasi global, validasi permintaan, atau inisialisasi data yang dibutuhkan. Contohnya : 

```
~ HTTPSRedirect
~ HTTPSWWWRedirect
~ WWWRedirect
~ AddTrailingSlash
~ RemoveTrailingSlash
~ MethodOverride
~ Rewrite

```

## Echo#Use()
Metode echo.Use() digunakan untuk mendaftarkan middleware yang akan dieksekusi pada tingkat grup rute atau rute individu .Middleware yang didaftarkan dengan echo.Use() akan dieksekusi sesuai dengan urutan pendaftaran dalam grup rute atau rute tertentu. Ini memungkinkan Anda untuk mengatur middleware yang lebih spesifik untuk rute tertentu dalam aplikasi Anda. Contohnya :

```
~ BodyLimit
~ Logger
~ Gzip
~ Recover
~ BasicAuth
~ JWTAuth
~ Secure
~ CORS
~ Static

```

## CORS Middleware
CORS (Cross-Origin Resource Sharing) Middleware adalah sebuah komponen perantara dalam framework web yang mengelola permintaan lintas domain dalam aplikasi web. CORS Middleware memungkinkan atau membatasi permintaan HTTP dari domain yang berbeda ke sumber daya di server Anda.

## Rate Limiter Middleware
Rate limiter middleware adalah sebuah komponen perantara yang digunakan untuk mengontrol seberapa sering klien dapat mengakses sumber daya atau layanan dalam suatu aplikasi. Middleware ini membatasi jumlah permintaan atau operasi yang dapat dilakukan oleh klien dalam interval waktu tertentu. Tujuan utama dari rate limiter adalah untuk melindungi server dari penyalahgunaan, serangan brute force, atau kelebihan permintaan yang dapat mengakibatkan beban server yang berlebihan.

## Log Middleware
Middleware logging adalah komponen perantara yang digunakan untuk mencatat informasi atau kejadian tertentu selama eksekusi perangkat lunak. Logging digunakan untuk memantau, menganalisis, atau memecahkan masalah dalam aplikasi Anda. Middleware logging akan mencatat informasi penting seperti permintaan HTTP, respons, kesalahan, atau aktivitas lainnya yang membantu dalam pemecahan masalah dan pemantauan aplikasi.

## Auth Middleware
Middleware otentikasi (auth middleware) adalah komponen perantara dalam sebuah aplikasi web atau API yang digunakan untuk memeriksa dan mengautentikasi pengguna sebelum memberi akses ke sumber daya yang dilindungi. Middleware ini berfungsi untuk memastikan bahwa pengguna yang mengakses suatu rute atau sumber daya telah melakukan otentikasi dengan benar dan memiliki izin yang sesuai.

## Basic Authentication
Basic authentication adalah metode otentikasi yang menggunakan header HTTP "Authorization" untuk mengirimkan nama pengguna (username) dan kata sandi (password) dalam bentuk yang dienkripsi base64.

```
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
    // Be careful to use constant time comparison to prevent timing attacks
    if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
        subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
        return true, nil
    }
    return false, nil
}))

```

## JWT Middleware
Otentikasi Berbasis Token JWT (JSON Web Token) adalah metode otentikasi yang menggunakan token berformat JSON untuk mengidentifikasi pengguna dan memberikan izin akses. Token JWT berisi informasi terkait pengguna, seperti ID, peran, atau klaim lainnya, yang digenerate oleh server setelah pengguna berhasil melakukan otentikasi. Token tersebut dikirimkan dalam header HTTP "Authorization" dalam setiap permintaan yang memerlukan otentikasi. Server memverifikasi tanda tangan token untuk memastikan keasliannya dan mengizinkan atau menolak akses berdasarkan klaim yang terdapat dalam token.

## Implementasi JWT Middleware

```
// Create constants

const SECRET_JWT = "legal"

```

```
// Create UserResponse

type UserResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

```

```
// Create JWT Middleware

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}

```

```
// Create Repo User Check Login
func CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User
	// select * from users where email = `email` and password = `password`
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return models.User{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(data.ID))
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

```

```
// Create Login Controller
func LoginController(c echo.Context) error {
	var loginReq = requests.LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind "+errBind.Error()))
	}

	data, token, err := repositories.CheckLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("error login"))
	}
	response := map[string]any{
		"token":   token,
		"user_id": data.ID,
		"name":    data.Name,
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success recieve user data", response))
}

```

```
// Routes

e.POST("/login", controllers.LoginController)
e.GET("/users", controllers.GetUserController, middlewares.JWTMiddleware())

```