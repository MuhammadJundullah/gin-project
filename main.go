package main

import (
	"ahmad.com/src/config"
	route "ahmad.com/src/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	db := config.DB()

	route.Api(r, db)

	r.Run()
}