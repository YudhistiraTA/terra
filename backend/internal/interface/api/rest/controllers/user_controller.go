package controllers

import (
	"github.com/YudhistiraTA/terra/internal/application/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(r *gin.RouterGroup, userService *services.UserService) {
	userController := &UserController{userService}

	user := r.Group("/user")
	user.POST("/login", userController.Login)
}

func (uc *UserController) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "OK"})
}
