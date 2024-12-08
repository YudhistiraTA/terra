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
	v1.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "OK"})
	})

	userService := services.NewUserService(ctx, db)

	controllers.NewUserController(v1, userService)

	return app
}
