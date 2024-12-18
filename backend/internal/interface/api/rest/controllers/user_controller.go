package controllers

import (
	"net/http"

	"github.com/YudhistiraTA/terra/internal/application/services"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/request"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/response"
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
	var userLoginRequest request.UserLoginRequest
	if err := ctx.BindJSON(&userLoginRequest); err != nil {
		response.ValidationErrorResponse(ctx, err, userLoginRequest)
		return
	}

	userLoginCommand := userLoginRequest.ToUserLoginCommand()
	result, err := uc.userService.Login(userLoginCommand)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Invalid email or password"})
		return
	}

	ctx.SetCookie("sessionToken", result.SessionToken, 900, "/", "", false, false)
	ctx.SetCookie("refreshToken", result.RefreshToken, 86400, "/", "", false, false)
	ctx.JSON(200, response.NewSuccessResponse(nil))
}
