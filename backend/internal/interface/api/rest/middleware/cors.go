package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func getAllowHeaders() string {
	var h []string = []string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"X-SIGNATURE",
		"X-TIMESTAMP",
	}
	return strings.Join(h, ", ")
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", getAllowHeaders())
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return corsMiddleware
}