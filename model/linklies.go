package model

import (
	"time"
)

type Linkly struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Redirect       string    `json:"redirect" gorm:"not null"`
	Linkly         string    `json:"linkly" gorm:"not null;unique"`
	Clicked        int       `json:"clicked" gorm:"default:0"`
	IsAvailable    bool      `json:"is_available" gorm:"default:true"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;autoCreateTime:nano"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	User_id        int       `json:"user_id" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:""`
	User           User      `gorm:"foreignKey:User_id"`
}