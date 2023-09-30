package repositories

import (
    "db_API/models"
    "errors"
)

type UserRepositoryMock struct{}

func NewUserRepositoryMock() *UserRepositoryMock {
    return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) QuerySelectAllUser() ([]models.User, error) {
    // Mengembalikan error sebagai simulasi kegagalan query
    return nil, errors.New("Database error: unable to fetch users")
}
