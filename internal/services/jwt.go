package services

import (
	"github.com/golang-jwt/jwt/v4"
	"testapi/internal/constants"
	"testapi/internal/models"
	"testapi/internal/utils"
	"time"
)

func GenerateAccessToken(userDto models.UserDto) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userDto.ID,
		"email":      userDto.Email,
		"created_at": userDto.CreatedAt,
		"updated_at": userDto.UpdatedAt,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(constants.JWT_ACCESS_TOKEN_VALIDITY_TIME * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET))
}

func GenerateRefreshToken(userID uint) (string, time.Time) {
	//generates random string with userid and hashes it for fixed length
	var rand = utils.GenerateRandomString(64)
	var hash = utils.HashMD5(rand + string(userID))
	return hash, time.Now().Add(constants.JWT_REFRESH_TOKEN_VALIDITY_TIME * time.Second)
}

func ParseJwtToken(jwtToken string) (models.UserDto, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET), nil
	})
	if err != nil {
		return models.UserDto{}, err
	}
	if !token.Valid {
		return models.UserDto{}, err
	}
	return models.UserDto{
		ID:        uint(claims["user_id"].(float64)),
		Email:     claims["email"].(string),
		CreatedAt: claims["created_at"].(string),
		UpdatedAt: claims["updated_at"].(string),
	}, nil
}
