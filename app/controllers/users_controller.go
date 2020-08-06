package controllers

import (
	"github.com/bracketcore/go-rest-api-template/app/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService *services.UserService
}

func ProvideUsersController(u *services.UserService) *UsersController {
	return &UsersController{userService: u}
}

func (usersController *UsersController) Profile(c *gin.Context) {

}
