package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "api",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

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
