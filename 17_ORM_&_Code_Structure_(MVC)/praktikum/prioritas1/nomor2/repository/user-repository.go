package repository

import (
	"db_API/prioritas1/nomor2/config"
	"db_API/prioritas1/nomor2/model"
)

func QuerySelectAllUser() ([]model.User, error) {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
        return nil, err
    }
	return users, nil
}

func QuerySelectUserById(idUser int) (model.User, error) {
	var user model.User
	result := config.DB.First(&user, idUser)
	if result.Error != nil {
        return user, result.Error
    }
    return user, nil
}

func QueryCreateUser(user *model.User) error {
	return config.DB.Save(user).Error
}

func QueryDeleteUserById(idUser int) error {
	var user model.User
	result := config.DB.First(&user, idUser)

	if result.Error != nil {
        return result.Error
    }

	config.DB.Delete(&user)
    return nil
}

func QueryUpdateUserById(idUser int, updatedUser *model.User) error {
	var user model.User
	result := config.DB.First(&user, idUser)

	if result.Error != nil {
        return result.Error
    }

	if err := config.DB.Model(&user).Updates(updatedUser).Error; 
	err != nil {
        return err
    }

	return nil
}


