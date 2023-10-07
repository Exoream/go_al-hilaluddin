package repositories

import (
	"clean/api/entity"
	"clean/api/middlewares"
	"clean/api/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.DataInterface {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user entity.Main) (row int, err error) {
	var data = models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	tx := u.db.Create(&data)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (u *userRepository) FindAllUsers() ([]entity.Main, error) {
	var users []models.User

	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	mainData := make([]entity.Main, len(users))

	for i, user := range users {
		mainData[i] = entity.Main{
			ID:        int(user.ID),
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return mainData, nil
}

func (u *userRepository) CheckLogin(email, password string) (entity.Main, string, error) {
	var user models.User

	tx := u.db.Where("email = ? AND password = ?", email, password).First(&user)
	if tx.Error != nil {
		return entity.Main{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(user.ID))
		if errToken != nil {
			return entity.Main{}, "", errToken
		}
	}

	data := entity.Main{
		ID:        int(user.ID),
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return data, token, nil
}


