package main

import (
  "net/http"
  "strconv"
  "github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUserControllerById(c echo.Context) error {
  id := c.Param("id")
  idUsers, errConv := strconv.Atoi(id)
  if errConv != nil {
    return c.JSON(http.StatusBadRequest, map[string]any{
      "message" : "error convert",
      "status" : "failed",
    })
  }
  
  for _, user := range users {
    if user.Id == idUsers {
      return c.JSON(http.StatusOK, map[string]interface{}{
        "message" : "user found",
        "status" : "success",
        "user" : user,
      })
    }
  }

  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
    "status":  "failed",
  })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  id := c.Param("id")
  idUsers, errConv := strconv.Atoi(id)
  if errConv != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "error convert",
      "status":  "failed",
    })
  }

  var newUsers []User
  userFound := false

  for _, user := range users {
    if user.Id == idUsers {
      userFound = true
    } else {
      newUsers = append(newUsers, user)
    }
  }

  if userFound {
    users = newUsers
    return c.JSON(http.StatusOK, map[string]interface{}{
      "message": "user deleted",
      "status":  "success",
    })
  }

  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
    "status":  "failed",
  })

}


// update user by id
func UpdateUserController(c echo.Context) error {
  id := c.Param("id")
  idUsers, errConv := strconv.Atoi(id)
  if errConv != nil {
    return c.JSON(http.StatusBadRequest, map[string]any{
      "message" : "error convert",
      "status" : "failed",
    })
  }

  updatedUser := User{}
  err := c.Bind(&updatedUser)
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]any{
      "message" : "error bind data",
      "status" : "failed",
    })
  }

  for index, user := range users {
    if user.Id == idUsers {
      users[index] = updatedUser
      return c.JSON(http.StatusOK, map[string]any{
        "message" : "user update",
        "status" : "succes",
        "updated_user" : updatedUser,
      })
    }
  }

  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "message": "user not found",
    "status":  "failed",
  })
}


// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}
// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.GET("/users/:id", GetUserControllerById)
  e.POST("/users", CreateUserController)
  e.PUT("/users/:id", UpdateUserController)
  e.DELETE("/users/:id", DeleteUserController)
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}