package response

type UserLoginResponse struct {
	SessionToken  string `json:"sessionToken"`
	RefreshToken  string `json:"refreshToken"`
	SessionExpiry string `json:"sessionExpiry"`
	RefreshExpiry string `json:"refreshExpiry"`
	ID            string `json:"id"`
	Email         string `json:"email"`
}
