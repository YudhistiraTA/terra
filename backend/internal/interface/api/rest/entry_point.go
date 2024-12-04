package rest

import (
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewEntryPoint() *gin.Engine {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	v1 := app.Group("/v1")
	v1.GET("/health", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "OK"}) })
	return app
}
