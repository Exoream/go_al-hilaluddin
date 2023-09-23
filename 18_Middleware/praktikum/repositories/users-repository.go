package repositories

import (
	"db_API/config"
	"db_API/middlewares"
	"db_API/models"
)

func QueryCheckLoginUser(email string, password string) (models.User, string, error) {
	var user models.User
	result := config.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return models.User{}, "", result.Error
	}

	var token string
	if result.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(user.ID))
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return user, token, nil
}

func QuerySelectAllUser() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func QuerySelectUserById(idUser int) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, idUser)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func QueryCreateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func QueryDeleteUserById(idUser int) error {
	var user models.User
	result := config.DB.First(&user, idUser)

	if result.Error != nil {
		return result.Error
	}

	config.DB.Delete(&user)
	return nil
}

func QueryUpdateUserById(idUser int, updatedUser *models.User) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, idUser)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := config.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	var updatedUserData models.User
	config.DB.First(&updatedUserData, idUser)

	return &updatedUserData, nil
}
