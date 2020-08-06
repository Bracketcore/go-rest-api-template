package main

import (
	"github.com/bracketcore/go-rest-api-template/app"
	"github.com/bracketcore/go-rest-api-template/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := app.InitDB(config)
	defer db.Close()

	router := gin.Default()

	userRouter := app.InitUserRouter(db)
	userRouter.LoadGuestRoutes(router)

	server := &http.Server{Addr: "localhost:" + config.Port, Handler: router}
	log.Fatal(server.ListenAndServe())
}
