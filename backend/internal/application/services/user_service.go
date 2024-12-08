package services

import (
	"context"
	"os"
	"time"

	"github.com/YudhistiraTA/terra/internal/application/command"
	"github.com/YudhistiraTA/terra/internal/application/common"
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx context.Context
	db  *sqlc.Queries
}

func NewUserService(ctx context.Context, db *sqlc.Queries) *UserService {
	return &UserService{ctx, db}
}

func (us *UserService) Login(user command.UserLoginCommand) (*command.UserLoginCommandResult, error) {
	dbUser, err := us.db.GetUserByEmail(us.ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return nil, err
	}
	claimStem := common.UserClaim{
		ID:    dbUser.ID.String(),
		Email: dbUser.Email,
	}
	sessionClaim := claimStem
	sessionClaim.IssuedAt = jwt.NewNumericDate(time.Now())
	sessionClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	sessionToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaim)
	sessionId, err := sessionToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	refreshClaim := claimStem
	refreshClaim.IssuedAt = jwt.NewNumericDate(time.Now())
	refreshClaim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	refreshId, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &command.UserLoginCommandResult{
		RefreshToken:       refreshId,
		SessionToken:       sessionId,
		SessionTokenExpiry: sessionClaim.ExpiresAt.Time,
		RefreshTokenExpiry: refreshClaim.ExpiresAt.Time,
		Claim:              claimStem,
	}, nil
}
