package repositories

import (
	"db_API/prioritas2/config"
	"db_API/prioritas2/models"
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

func QueryCreateBook(book *models.Book) error {
	return config.DB.Save(book).Error
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
