package controllers

import (
	"log"
	"net/http"

	"github.com/YudhistiraTA/terra/internal/application/command"
	"github.com/YudhistiraTA/terra/internal/application/services"
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/mapper"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/request"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController struct {
	postService *services.PostService
}

func NewPostController(r *gin.RouterGroup, postService *services.PostService) {
	postController := &PostController{postService}
	post := r.Group("/posts")
	post.GET("/list", postController.GetPostList)
	post.POST("/create", postController.CreatePost)
}

func (pc *PostController) GetPostList(ctx *gin.Context) {
	var search *string
	if ctx.Query("search") != "" {
		q := ctx.Query("search")
		search = &q
	}
	var cursor *uuid.UUID
	if ctx.Query("cursor") != "" {
		q, err := uuid.Parse(ctx.Query("cursor"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid cursor"})
			return
		}
		cursor = &q
	}
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		return
	}
	cmd := command.PostListCommand{
		Search: search,
		Cursor: cursor,
		UserId: user.(sqlc.User).ID,
	}
	result, err := pc.postService.GetPostList(cmd)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, response.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	data := mapper.ToPostListResponse(result)
	ctx.JSON(200, response.NewSuccessResponse(data))
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var req request.CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationErrorResponse(ctx, err, req)
		return
	}
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
		return
	}
	cmd := command.CreatePostCommand{
		UserId:  user.(sqlc.User).ID,
		Title:   req.Title,
		Content: req.Content,
	}
	err := pc.postService.CreatePost(cmd)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}
	ctx.JSON(200, response.NewSuccessResponse(nil))
}
