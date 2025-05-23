package main

import (
	"ahmad.com/src/config"
	// "ahmad.com/src/modules/user"
	route "ahmad.com/src/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	db := config.InitDB()
	// userRepo := user.NewUserRepository(db)

	route.Api(r, db)

	r.Run()
}