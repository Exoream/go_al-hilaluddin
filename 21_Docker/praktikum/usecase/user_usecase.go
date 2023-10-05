package usecase

import (
	"clean/api/entity"
	"errors"
	"fmt"
	"regexp"
)

type userUseCase struct {
	userRepo entity.DataInterface
}


func NewUserUsecase(userRepo entity.DataInterface) entity.UseCaseInterface {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(data entity.Main) (row int, err error) {

	// Validasi email dan format email
	if data.Email == "" {
		return 0, errors.New("error. email harus diisi")
	}

	// RegEx untuk memeriksa format email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, data.Email)
	if !match {
		return 0, errors.New("error. format email tidak valid")
	}

	if data.Password == "" {
		return 0, errors.New("error. password harus diisi")
	}

	row , err = uc.userRepo.Create(data)
	if err != nil {
		return 0, err
	}
	return row, nil
}

// CheckLogin implements entity.UseCaseInterface.
func (uc *userUseCase) CheckLogin(email string, password string) (entity.Main, string, error) {
	if email == "" || password == "" {
		return entity.Main{}, "", errors.New("error. email dan password harus diisi")
	}
	
	// RegEx untuk memeriksa format email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return entity.Main{}, "", errors.New("error. format email tidak valid")
	}

	userData, token, err := uc.userRepo.CheckLogin(email, password)
	if err != nil {
		if err.Error() == "record not found" {
			return entity.Main{}, "", fmt.Errorf("user not found")
		}
		return entity.Main{}, "", err
	}

	return userData, token, nil
}

// Find implements entity.UseCaseInterface.
func (uc *userUseCase) FindAllUsers() ([]entity.Main, error) {
	users, err := uc.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}