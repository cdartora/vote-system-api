package controllers

import (
	"net/http"

	"example.com/vote-system-api/src/middleware"
	"example.com/vote-system-api/src/models"
	"example.com/vote-system-api/src/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userservice services.UserService) UserController {
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
	idParam := ctx.Param("id")
	id, _ := uuid.FromString(idParam)
	user, err := uc.UserService.GetUser(&id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := uuid.FromString(idParam)
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := uc.UserService.UpdateUser(&user, &id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := uuid.FromString(idParam)
	err := uc.UserService.DeleteUser(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.RequestUser

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	token, ok := uc.UserService.Login(user)
	if ok != true {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid fields"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/", uc.CreateUser)
	userroute.POST("/login", uc.Login)
	userroute.Use(middleware.AuthorizeJWT)
	userroute.GET("/:id", uc.GetUser)
	userroute.GET("/", uc.GetAll)
	userroute.PATCH("/:id", uc.UpdateUser)
	userroute.DELETE("/:id", uc.DeleteUser)
}
