package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password;"`
	IsValid  bool   `gorm:"not null;default:false" json:"is_valid;"`
}

type PasswordChangeLog struct {
	gorm.Model
	UserID      uint   `gorm:"not null" json:"user_id"`
	OldPassword string `gorm:"not null" json:"password"`
	NewPassword string `gorm:"not null" json:"password"`
	User        User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}

type UserLoginLog struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}

type UserRefreshToken struct {
	gorm.Model
	UserID       uint      `gorm:"not null" json:"user_id"`
	RefreshToken string    `gorm:"not null" json:"token"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	User         User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}
