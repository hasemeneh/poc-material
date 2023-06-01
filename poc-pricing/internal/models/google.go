package models

import "time"

type GoogleAuthenticatedResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
type AuthenticatorResponse struct {
	*Token
	AuthID string
}
type Token struct {
	AccessToken  string    `json:"access_token"`
	Expiry       time.Time `json:"expiry"`
	RefreshToken string    `json:"refresh_token"`
}
