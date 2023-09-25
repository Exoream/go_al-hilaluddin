package repositories

import (
	"db_API/config"
	"db_API/helpers"
	"db_API/middlewares"
	"db_API/models"
	"errors"
)

func QueryCheckLoginUser(email, password string) (models.User, string, error) {
	var user models.User
	tx := config.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return models.User{}, "", tx.Error
	}

	if helpers.CheckPasswordHash(user.Password, password) {
		token, errToken := middlewares.CreateToken(int(user.ID))
		if errToken != nil {
			return models.User{}, "", errToken
		}
		return user, token, nil
	}

	// if !helpers.CheckPasswordHash(user.Password, password) {
	// 	// Kata sandi tidak cocok
	// 	return models.User{}, "", errors.New("Incorrect password")
	// }

	return models.User{}, "", errors.New("Login failed")
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
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

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
