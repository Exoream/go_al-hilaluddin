package repositories

import (
	"clean/api/config"
	"clean/api/entity"
	"clean/api/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = config.InitDBTest()

func TestCreateUsers(t *testing.T) {
	t.Run("Succes Create User", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		db.AutoMigrate(&models.User{})

		userRepo := NewUserRepository(db)

		testUser := entity.Main{
			Email:    "hilal@gmail.com",
			Password: "123asd",
		}

		rowsAffected, err := userRepo.Create(testUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, rowsAffected)
	})

	t.Run("Failed Create User", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		userRepo := NewUserRepository(db)

		testUser := entity.Main{
			Email:    "hilal@gmail.com",
			Password: "123asd",
		}

		rowsAffected, err := userRepo.Create(testUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rowsAffected)
	})
}

func TestSelectData(t *testing.T) {
	t.Run("Succes Get Users", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		db.AutoMigrate(&models.User{})
		repo := NewUserRepository(db)

		testUser := models.User{
			Email:    "hilal@gmail.com",
			Password: "123",
		}
		db.Create(&testUser)
		dataResult, err := repo.FindAllUsers()
		assert.Nil(t, err)

		if len(dataResult) > 0 {
			userArray := make([]string, 0)

			for _, user := range dataResult {
				userArray = append(userArray, user.Email)
				userArray = append(userArray, user.Password)
			}

		} else {
			t.Errorf("No data retrieved from the database")
		}
	})

	t.Run("Failed Get Users", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		repo := NewUserRepository(db)

		result, err := repo.FindAllUsers()
		assert.NotNil(t, err)
		assert.Nil(t, result)

	})
}

func TestLogin(t *testing.T) {
	t.Run("Succes Login", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		db.AutoMigrate(&models.User{})
		repo := NewUserRepository(db)

		testUser := models.User{
			Email:    "hilal@gmail.com",
			Password: "123",
		}
		db.Create(&testUser)

		email := "hilal@gmail.com"
		password := "123"
		data, token, err := repo.CheckLogin(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, email, data.Email)
		assert.Equal(t, password, data.Password)
	})

	t.Run("Gagal Login", func(t *testing.T) {
		db.Migrator().DropTable(&models.User{})
		db.AutoMigrate(&models.User{})
		repo := NewUserRepository(db)

		email := "invalid@gmail.com"
		password := "wrongpassword"
		data, token, err := repo.CheckLogin(email, password)

		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.Equal(t, "", data.Email)
		assert.Equal(t, "", data.Password)
	})
}
