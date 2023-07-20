package model

import (
	"gorm.io/gorm"
	_ "time"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	IsActive bool   `json:"is_active" grom:"default:true"`
	IsAdmin  bool   `json:"is_admin" grom:"default:false"`
}
