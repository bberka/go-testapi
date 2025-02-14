package database

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey;not null" json:"id"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password;"`
	IsValid   bool       `gorm:"not null;default:false" json:"is_valid;"`
}

type PasswordChangeLog struct {
	ID          uint      `gorm:"primaryKey;not null" json:"id"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	OldPassword string    `gorm:"not null" json:"old_password"`
	NewPassword string    `gorm:"not null" json:"new_password"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}

type UserLoginLog struct {
	ID        uint      `gorm:"primaryKey;not null" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}

type UserRefreshToken struct {
	ID           uint      `gorm:"primaryKey;not null" json:"id"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	RefreshToken string    `gorm:"not null" json:"token"`
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`
	User         User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Association field
}
