package services

import (
	"github.com/bracketcore/go-rest-api-template/app/models"
	"github.com/bracketcore/go-rest-api-template/app/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func ProvideUserService(u *repositories.UserRepository) *UserService {
	return &UserService{UserRepository: u}
}

func (u *UserService) SaveUser(user *models.User) error {
	return u.UserRepository.Save(user)
}
