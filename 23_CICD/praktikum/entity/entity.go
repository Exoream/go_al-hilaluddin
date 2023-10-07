package entity

import (
	"time"

)

type Main struct {
	ID        int   
	Email     string 
	Password  string 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Create(data Main) (row int, err error)
	FindAllUsers() ([]Main, error)
	CheckLogin(email, password string) (Main, string, error)
}

type UseCaseInterface interface {
	Create(data Main) (row int, err error)
	CheckLogin(email, password string) (Main, string, error)
	FindAllUsers() ([]Main, error)
}