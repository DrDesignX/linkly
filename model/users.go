package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	IsActive  bool      `json:"is_active" grom:"default:true"`
	IsAdmin   bool      `json:"is_admin" grom:"default:false"`
}
