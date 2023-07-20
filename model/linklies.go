package model

import (
	"time"
	"gorm.io/gorm"
)

var(
	ExpirationDate = time.Now().Add(30 * 24 * time.Hour)
)
type Linkly struct {
	gorm.Model
	Redirect       string    `json:"redirect" gorm:"not null"`
	Linkly         string    `json:"linkly" gorm:"not null;unique"`
	Clicked        int       `json:"clicked" gorm:"default:0"`
	IsAvailable    bool      `json:"is_available" gorm:"default:true"`
	User_id        int       `json:"user_id" gorm:"not null"`
	ExpirationDate time.Time `json:"expiration_date" gorm:""`
	User           User      `gorm:"foreignKey:User_id"`
}
