package services

import (
	"errors"
	"testapi/internal/database"
	"testapi/internal/models"
	"testapi/internal/utils"
	"time"
)

func Authenticate(email, password string) (*models.TokenResponse, error) {
	var user database.User
	result := database.DB.Where("email = ? AND is_valid = 1", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("Invalid credentials")
	}
	var hashedPassword = utils.HashPassword(password)
	// In production, compare hashed passwords!
	if user.Password != hashedPassword {
		return nil, errors.New("Invalid credentials")
	}

	var loginLog = database.UserLoginLog{
		UserID: user.ID,
	}
	database.DB.Create(&loginLog)

	var userDto = models.UserDto{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	var token, err = GenerateAccessToken(userDto)
	if err != nil {
		return nil, err
	}

	var refreshToken, expiresAt = GenerateRefreshToken(user.ID)
	var refreshTokenEntity = database.UserRefreshToken{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}
	database.DB.Create(&refreshTokenEntity)

	var response = &models.TokenResponse{
		Token:            token,
		RefreshToken:     refreshToken,
		ExpiresAt:        expiresAt.Unix(),
		RefreshExpiresAt: expiresAt.Unix(),
	}

	return response, nil
}

func Register(user *models.RegisterRequest) (*models.UserDto, error) {
	var count int64
	countResult := database.DB.Model(&database.User{}).Where("email = ?", user.Email).Count(&count)
	if countResult.Error != nil {
		return nil, countResult.Error
	}
	if count > 0 {
		return nil, errors.New("User already exists")
	}
	hashPassword := utils.HashPassword(user.Password)
	newUser := database.User{
		Email:    user.Email,
		Password: hashPassword,
		IsValid:  true,
	}
	createResult := database.DB.Create(&newUser)
	if createResult.Error != nil {
		return nil, createResult.Error
	}

	userDto := models.UserDto{
		ID:        newUser.ID,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: newUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &userDto, nil
}

func ChangePassword(id uint, req *models.ChangePasswordRequest) (bool, error) {
	isMatch := req.NewPassword == req.NewPasswordRepeat
	if !isMatch {
		return false, errors.New("Passwords do not match")
	}
	var isOldSameWithNew = req.OldPassword == req.NewPassword
	if isOldSameWithNew {
		return false, errors.New("Old password and new password cannot be the same")
	}
	var user database.User
	result := database.DB.Where("ID = ? AND is_valid = 1", id).First(&user)
	if result.Error != nil {
		return false, errors.New("User not found")
	}
	hashedPassword := utils.HashPassword(req.OldPassword)
	if user.Password != hashedPassword {
		return false, errors.New("Incorrect password")
	}
	newPassword := utils.HashPassword(req.NewPassword)
	user.Password = newPassword
	result = database.DB.Save(&user)
	if result.Error != nil {
		return false, result.Error
	}

	passwordChangeLog := database.PasswordChangeLog{
		UserID:      user.ID,
		OldPassword: hashedPassword,
		NewPassword: newPassword,
	}
	database.DB.Create(&passwordChangeLog)
	return true, nil
}

func RefreshToken(req *models.RefreshTokenRequest) (*models.TokenResponse, error) {
	var refreshToken database.UserRefreshToken
	var refreshTokenEntity = database.DB.Where("refresh_token = ? AND expires_at > ?", req.RefreshToken, time.Now()).First(&refreshToken)
	if refreshTokenEntity.RowsAffected == 0 {
		return nil, errors.New("Invalid refresh token")
	}
	var user = database.User{}
	database.DB.Where("ID = ?", refreshToken.UserID).First(&user)
	var userDto = models.UserDto{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	var token, err = GenerateAccessToken(userDto)
	if err != nil {
		return nil, err
	}

	var newRefreshToken, expiresAt = GenerateRefreshToken(user.ID)
	refreshToken.RefreshToken = newRefreshToken
	refreshToken.ExpiresAt = expiresAt
	database.DB.Save(&refreshToken)

	var response = &models.TokenResponse{
		Token:            token,
		RefreshToken:     newRefreshToken,
		ExpiresAt:        expiresAt.Unix(),
		RefreshExpiresAt: expiresAt.Unix(),
	}

	return response, nil
}
