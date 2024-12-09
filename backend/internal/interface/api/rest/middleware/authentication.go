package middleware

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/YudhistiraTA/terra/internal/application/common"
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrUnauthorized = response.ErrorResponse{Message: "Unauthorized"}

func Authentication(serverCtx context.Context, db *sqlc.Queries) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			panic("JWT_SECRET is not set")
		}
		sessionToken := ctx.Request.Header.Get("Authorization")
		if sessionToken == "" {
			ctx.JSON(401, ErrUnauthorized)
			ctx.Abort()
			return
		}
		token := strings.TrimPrefix(sessionToken, "Bearer ")
		if token == "" {
			ctx.JSON(401, ErrUnauthorized)
			ctx.Abort()
			return
		}
		jwtToken, err := jwt.ParseWithClaims(token, &common.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			if err.Error() == "token has invalid claims: token is expired" {
				refreshToken, err := ctx.Cookie("refreshToken")
				if err != nil {
					ctx.JSON(401, ErrUnauthorized)
					ctx.Abort()
					return
				}
				refreshTokenClaims, err := jwt.ParseWithClaims(refreshToken, &common.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtSecret), nil
				})
				if err != nil {
					ctx.JSON(401, ErrUnauthorized)
					ctx.Abort()
					return
				}
				claims, _ := refreshTokenClaims.Claims.(*common.UserClaim)
				uuid, err := uuid.Parse(claims.ID)
				if err != nil {
					ctx.JSON(401, ErrUnauthorized)
					ctx.Abort()
					return
				}
				user, err := db.GetUserById(serverCtx, uuid)
				if err != nil {
					ctx.JSON(401, ErrUnauthorized)
					ctx.Abort()
					return
				}
				claimStem := common.UserClaim{
					ID:    user.ID.String(),
					Email: user.Email,
				}
				sessionClaim := claimStem
				sessionClaim.IssuedAt = jwt.NewNumericDate(time.Now())
				sessionClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
				sessionToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaim)
				sessionId, err := sessionToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
				if err != nil {
					ctx.JSON(401, ErrUnauthorized)
					ctx.Abort()
					return
				}
				ctx.SetCookie("sessionToken", sessionId, 900, "/", "", false, false)
			} else {
				ctx.JSON(401, ErrUnauthorized)
				ctx.Abort()
				return
			}
		}
		claims, _ := jwtToken.Claims.(*common.UserClaim)

		uuid, err := uuid.Parse(claims.ID)
		if err != nil {
			ctx.JSON(401, ErrUnauthorized)
			ctx.Abort()
			return
		}
		user, err := db.GetUserById(serverCtx, uuid)
		if err != nil {
			ctx.JSON(401, ErrUnauthorized)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
