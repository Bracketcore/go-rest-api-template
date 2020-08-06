package repositories

import (
	"github.com/bracketcore/go-rest-api-template/app/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Save(user *models.User) error {
	return u.DB.Save(user).Error
}
