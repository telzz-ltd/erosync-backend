package users

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type VerifyEmailRequest struct {
	OtpCode string `json:"code" validate:"required"`
}

type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
