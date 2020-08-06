//+build wireinject

package app

import (
	"github.com/bracketcore/go-rest-api-template/app/controllers"
	"github.com/bracketcore/go-rest-api-template/app/repositories"
	"github.com/bracketcore/go-rest-api-template/app/services"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitUserRouter(db *gorm.DB) *UserRouter {
	wire.Build(
		repositories.ProvideUserRepository,
		services.ProvideUserService,
		controllers.ProvideUsersController,
		controllers.ProvideAuthController,
		ProvideUserRouter,
	)
	return &UserRouter{}
}
