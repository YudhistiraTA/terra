package common

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
