package rest

import (
	"context"

	"github.com/YudhistiraTA/terra/internal/application/services"
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/controllers"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewRestServer(ctx context.Context, db *sqlc.Queries) *gin.Engine {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())

	v1 := app.Group("/v1")

	userService := services.NewUserService(ctx, db)
	controllers.NewUserController(v1, userService)

	v1.Use(middleware.Authentication(ctx, db))
	v1.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})

	return app
}
