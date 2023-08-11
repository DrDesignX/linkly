package models

import (
	"log"
	"time"

	Initializers "github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" gorm:"unique;not null;size:255"`
	Email     string         `json:"email" gorm:"unique;not null;size:255"`
	Password  string         `json:"password" gorm:"not null;size:255"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// CreateUser creates a new user in the database.
func CreateUser(username, email, password string) error {
	pass, err := utils.HashPassword(password)
	if err != nil {
		log.Println("error generating hash password ", err)
		return err
	}

	newUser := &User{
		Username: username,
		Email:    email,
		Password: pass,
		IsActive: true,
	}

	return Initializers.DB.Create(newUser).Error
}

// getUser gets a user by username
func GetUser(db *gorm.DB, username string) (*User, error) {
	user := &User{}
	err := db.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates information of a user
func UpdateUser(db *gorm.DB, username string, email string, password string) error {
	user, err := GetUser(db, username)
	if err != nil {
		return err
	}
	user.Email = email
	user.Password = password
	return db.Save(user).Error
}

// deleteUser deletes a user
func DeleteUser(db *gorm.DB, username string) error {
	user, err := GetUser(db, username)
	if err != nil {
		return err
	}
	return db.Delete(user).Error
}

// GetLinklies gets all linklies of a user
func GetLinklies(db *gorm.DB, username string) ([]*Linkly, error) {
	user, err := GetUser(db, username)
	if err != nil {
		return nil, err
	}
	links := []*Linkly{}
	err = db.Where("user_id = ?", user.ID).Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}
