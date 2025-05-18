package route

import (
	user "ahmad.com/src/modules/user"
	handler "ahmad.com/src/modules/user/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Api(router *gin.Engine, db *gorm.DB) {
	// Init repository, service, controller
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userController := handler.NewUserController(userService)

	// Routing
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userController.Index)
		v1.GET("/users/:id", userController.GetByID)
		v1.POST("/users", userController.Create)
		v1.PATCH("/users/:id", userController.Update)
		v1.DELETE("/users/:id", userController.Delete)
	}
}
