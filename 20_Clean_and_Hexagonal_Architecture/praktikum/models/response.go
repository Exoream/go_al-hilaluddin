package models

type UserLoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserResponse struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UsersResponseSuccess struct {
	Message string
	Data    []UserResponse
}

type ResponseMessage struct {
	Message string
}