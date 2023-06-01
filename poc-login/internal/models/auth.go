package models

type RegisterReq struct {
	AuthID          string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
