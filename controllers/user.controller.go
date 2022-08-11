package controllers

import (
	"net/http"

	"example.com/vote-system-api/models"
	"example.com/vote-system-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uc.UserService.GetUser(&id)

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", uc.CreateUser)
	userroute.GET("/", uc.GetAll)
	userroute.GET("/:id", uc.GetUser)
	userroute.PATCH("/:id", uc.UpdateUser)
	userroute.DELETE("/:id", uc.DeleteUser)
}
