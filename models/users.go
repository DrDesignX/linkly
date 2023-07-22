package models

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

// CreateUser creates a new user
func CreateUser(db *gorm.DB, username string, email string, password string) error {
	newuser := &User{
		Username: username,
		Email:    email,
		Password: password,
	}
	return db.Create(newuser).Error
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