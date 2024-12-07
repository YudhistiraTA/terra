package mapper

import (
	"time"

	"github.com/YudhistiraTA/terra/internal/application/command"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/response"
)

func ToUserLoginResponse(cmd *command.UserLoginCommandResult) *response.UserLoginResponse {
	return &response.UserLoginResponse{
		SessionToken:  cmd.SessionToken,
		RefreshToken:  cmd.RefreshToken,
		SessionExpiry: cmd.SessionTokenExpiry.Format(time.RFC3339),
		RefreshExpiry: cmd.RefreshTokenExpiry.Format(time.RFC3339),
		ID:            cmd.Claim.ID,
		Email:         cmd.Claim.Email,
	}
}
