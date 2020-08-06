package app

import (
	"github.com/bracketcore/go-rest-api-template/app/controllers"
	"github.com/bracketcore/go-rest-api-template/app/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	AuthController  *controllers.AuthController
	UserControllers *controllers.UsersController
}

// ProvideUserRouter for initialization
func ProvideUserRouter(auth *controllers.AuthController,
	usersController *controllers.UsersController) *UserRouter {
	return &UserRouter{
		AuthController:  auth,
		UserControllers: usersController,
	}
}

// LoadGuestRoutes for access
func (u *UserRouter) LoadGuestRoutes(router *gin.Engine) {
	guestRouter := router.Group("/api/v1")

	guestRouter.POST("/login", u.AuthController.LoginUser)
	guestRouter.POST("/signup", u.AuthController.RegisterUser)
}

// LoadAuthRoutes for access
func (u *UserRouter) LoadAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/api/v1")

	authRouter.Use(middlewares.Auth())
}
