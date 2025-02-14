package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	AgreeTerms     bool   `json:"agree_terms"`
}

type ChangePasswordRequest struct {
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
