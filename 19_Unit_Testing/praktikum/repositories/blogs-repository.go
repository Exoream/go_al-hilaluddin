package repositories

import (
	"db_API/config"
	"db_API/models"
)

func QuerySelectAllBlog() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func QuerySelectBlogById(idBlog int) (models.Blog, error) {
	var blog models.Blog
	result := config.DB.First(&blog, idBlog)
	if result.Error != nil {
		return blog, result.Error
	}
	return blog, nil
}

func QueryCreateBlog(blog *models.Blog) (interface{}, error) {
	tx := config.DB.Create(&blog)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return blog, nil
}

func QueryDeleteBlogById(idBlog int) error {
	var blog models.Blog
	result := config.DB.First(&blog, idBlog)

	if result.Error != nil {
		return result.Error
	}

	config.DB.Delete(&blog)
	return nil
}

func QueryUpdateBlogById(idBlog int, updatedBlog *models.Blog) (*models.Blog, error) {
	var blog models.Blog
	result := config.DB.First(&blog, idBlog)

	if result.Error != nil {
		return nil, result.Error
	}

	if err := config.DB.Model(&blog).Updates(updatedBlog).Error; err != nil {
		return nil, err
	}

	var updatedBlogData models.Blog
    config.DB.First(&updatedBlogData, idBlog)

    return &updatedBlogData, nil
}
