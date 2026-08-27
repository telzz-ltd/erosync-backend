package dto

import "erosync/internal/model"

type AuthResponse struct {
	User         model.User `json:"user"`
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
}
