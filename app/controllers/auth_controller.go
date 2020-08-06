package controllers

import (
	"log"
	"net/http"

	"github.com/bracketcore/go-rest-api-template/app/models"
	"github.com/bracketcore/go-rest-api-template/app/services"
	"github.com/bracketcore/go-rest-api-template/app/utils"
	"github.com/bracketcore/go-rest-api-template/app/utils/validation"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *services.UserService
}

func ProvideAuthController(u *services.UserService) *AuthController {
	return &AuthController{userService: u}
}

func (auth *AuthController) RegisterUser(c *gin.Context) {
	var userRequest validation.UserRequest

	if err := userRequest.ValidateSignUp(c); err != nil {
		utils.InputBadRequestResponse(c, err)
		return
	}

	// Hash user password
	hashedPassword, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		log.Println(err)
		utils.InternalServerErrorResponse(c, err)
		return
	}

	// Prepare user to insert into database
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: hashedPassword,
	}

	// Save user to database
	if err := auth.userService.SaveUser(&user); err != nil {
		log.Println(err)
		utils.DuplicateEntryErrorResponse(c, "Email has already been taken")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (auth *AuthController) LoginUser(c *gin.Context) {
	var userRequest validation.UserRequest

	if err := userRequest.ValidateLogin(c); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}
}
