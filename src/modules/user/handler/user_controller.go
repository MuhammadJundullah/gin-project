package user

import (
	"net/http"
	"strconv"

	user "ahmad.com/src/modules/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService user.UserService
}


func NewUserController(userService user.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Index(ctx *gin.Context) {
	data := uc.UserService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Error",
			"data":   "Invalid ID",
		})
		return
	}

	data := uc.UserService.GetByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Create(ctx *gin.Context) {
	data, err := uc.UserService.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	data, err := uc.UserService.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (uc *UserController) Update(ctx *gin.Context) {
	data, err := uc.UserService.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}
