package command

import (
	"time"

	"github.com/YudhistiraTA/terra/internal/application/common"
)

type UserLoginCommand struct {
	Email    string
	Password string
}

type UserLoginCommandResult struct {
	SessionToken       string
	RefreshToken       string
	SessionTokenExpiry time.Time
	RefreshTokenExpiry time.Time
	Claim              common.UserClaim
}
