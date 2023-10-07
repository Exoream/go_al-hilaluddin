package repositories

import (
	"db_API/config"
	"db_API/models"
)

func QuerySelectAllBook() ([]models.Book, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func QuerySelectBookById(idBook int) (models.Book, error) {
	var book models.Book
	result := config.DB.First(&book, idBook)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func QueryCreateBook(book *models.Book) (interface{}, error) {
	tx := config.DB.Create(&book)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return book, nil
}


func QueryDeleteBookById(idBook int) error {
	var book models.Book
	result := config.DB.First(&book, idBook)

	if result.Error != nil {
		return result.Error
	}

	config.DB.Delete(&book)
	return nil
}

func QueryUpdateBookById(idBook int, updatedBook *models.Book) (*models.Book, error) {
	var book models.Book
	result := config.DB.First(&book, idBook)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := config.DB.Model(&book).Updates(updatedBook).Error; err != nil {
		return nil, err
	}

	var updatedBookData models.Book
	config.DB.First(&updatedBookData, idBook)

	return &updatedBookData, nil
}