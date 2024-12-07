package services

import (
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	db *sqlc.Queries
}

func NewUserService(db *sqlc.Queries) *UserService {
	return &UserService{db}
}

func (us *UserService) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "OK"})
}
