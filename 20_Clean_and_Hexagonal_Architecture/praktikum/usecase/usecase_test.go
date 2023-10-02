package usecase

import (
	"clean/api/entity"
	mocks "clean/api/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testUserData := entity.Main{
		Email:    "hilal@gmail.com",
		Password: "123",
	}

	mockDataRepo.On("Create", testUserData).Return(1, nil)
	row, err := uc.Create(testUserData)

	assert.Nil(t, err)
	assert.Equal(t, 1, row)
	mockDataRepo.AssertCalled(t, "Create", testUserData)
}

func TestCreateWithError(t *testing.T) {
    mockDataRepo := new(mocks.DataInterface)
    uc := NewUserUsecase(mockDataRepo)

    testUserData := entity.Main{
        Email:    "hilal@gmail.com",
        Password: "123",
    }

    mockDataRepo.On("Create", testUserData).Return(0, errors.New("kesalahan penggunaan"))
    row, err := uc.Create(testUserData)

    assert.NotNil(t, err) 
    assert.Equal(t, 0, row) 
    mockDataRepo.AssertCalled(t, "Create", testUserData)
}

func TestCreateInvalidEmailFormat(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testUserData := entity.Main{
		Email:    "invalid-email",
		Password: "123",
	}

	_, err := uc.Create(testUserData)
	assert.EqualError(t, err, "error. format email tidak valid")
}

func TestCreateEmptyEmail(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testUserData := entity.Main{
		Email:    "",
		Password: "123",
	}

	_, err := uc.Create(testUserData)
	assert.EqualError(t, err, "error. email harus diisi")
}

func TestUCreateEmptyPassword(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testUserData := entity.Main{
		Email:    "hilal@gmail.com",
		Password: "",
	}

	_, err := uc.Create(testUserData)
	assert.EqualError(t, err, "error. password harus diisi")
}

func TestCheckLoginEmptyEmailAndPassword(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	userData, token, err := uc.CheckLogin("", "")
	assert.EqualError(t, err, "error. email dan password harus diisi")

	assert.Equal(t, entity.Main{}, userData)
	assert.Equal(t, "", token)
}

func TestCheckLoginInvalidEmailFormat(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testEmail := "invalid-email"

	userData, token, err := uc.CheckLogin(testEmail, "123")
	assert.EqualError(t, err, "error. format email tidak valid")

	assert.Equal(t, entity.Main{}, userData)
	assert.Equal(t, "", token)
}

func TestCheckLoginUserNotFound(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testEmail := "hilal@gmail.com"
	testPassword := "123"

	mockDataRepo.On("CheckLogin", testEmail, testPassword).Return(entity.Main{}, "", errors.New("record not found"))
	userData, token, err := uc.CheckLogin(testEmail, testPassword)

	assert.EqualError(t, err, "user not found")

	assert.Equal(t, entity.Main{}, userData)
	assert.Equal(t, "", token)
}

func TestFindAllUsersError(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	mockDataRepo.On("FindAllUsers").Return(nil, errors.New("error fetching users"))
	users, err := uc.FindAllUsers()

	assert.EqualError(t, err, "error fetching users")
	assert.Nil(t, users)
}

func TestCheckLoginOtherError(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testEmail := "hilal@gmail.com"
	testPassword := "123"

	unexpectedErr := errors.New("unexpected error")
	mockDataRepo.On("CheckLogin", testEmail, testPassword).Return(entity.Main{}, "", unexpectedErr)

	userData, token, err := uc.CheckLogin(testEmail, testPassword)
	assert.EqualError(t, err, "unexpected error")

	assert.Equal(t, entity.Main{}, userData)
	assert.Equal(t, "", token)
}

func TestCheckLoginSuccess(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	testEmail := "hilal@gmail.com"
	testPassword := "password123"

	expectedUserData := entity.Main{
		Email: "hilal@gmail.com",
	}
	expectedToken := "token123"

	mockDataRepo.On("CheckLogin", testEmail, testPassword).Return(expectedUserData, expectedToken, nil)
	userData, token, err := uc.CheckLogin(testEmail, testPassword)

	assert.Nil(t, err)
	assert.Equal(t, expectedUserData, userData)
	assert.Equal(t, expectedToken, token)
}

func TestFindAllUsersSuccess(t *testing.T) {
	mockDataRepo := new(mocks.DataInterface)
	uc := NewUserUsecase(mockDataRepo)

	expectedUsers := []entity.Main{
		{Email: "hilal@gmail.com"},
		{Email: "hilal2@gmail.com"},
	}

	mockDataRepo.On("FindAllUsers").Return(expectedUsers, nil)
	users, err := uc.FindAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, users)
}
