package request

import "github.com/YudhistiraTA/terra/internal/application/command"

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (req *UserLoginRequest) ToUserLoginCommand() command.UserLoginCommand {
	return command.UserLoginCommand{
		Email:    req.Email,
		Password: req.Password,
	}
}
