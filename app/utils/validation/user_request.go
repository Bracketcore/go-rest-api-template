package validation

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// UserRequest input
type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *UserRequest) bindRequest(c *gin.Context) error {
	// var newRequest UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return err
	}

	// *request = newRequest
	return nil
}

// ValidateLogin request
func (request *UserRequest) ValidateLogin(c *gin.Context) error {
	if err := request.bindRequest(c); err != nil {
		return err
	}

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Email, validation.Required, is.EmailFormat),
		validation.Field(&request.Password, validation.Required),
	); err != nil {
		return err
	}

	return nil
}

// ValidateSignUp request
func (request *UserRequest) ValidateSignUp(c *gin.Context) error {

	if err := request.bindRequest(c); err != nil {
		return err
	}

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required, validation.Length(2, 30)),
		validation.Field(&request.Email, validation.Required, is.EmailFormat),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 50)),
	); err != nil {
		return err
	}

	return nil
}
