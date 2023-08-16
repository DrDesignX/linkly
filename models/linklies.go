package models

import (
	"log"
	"time"

	"github.com/drdesignx/linkly/initializers"
	"gorm.io/gorm"
)

var (
	expirationDate = time.Now().Add(30 * 24 * time.Hour)
)

type Linkly struct {
	ID             int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Redirect       string         `json:"redirect" gorm:"not null"`
	Linkly         string         `json:"linkly" gorm:"not null;unique"`
	Clicked        int            `json:"clicked" gorm:"default:0"`
	IsAvailable    bool           `json:"is_available" gorm:"default:true"`
	User_id        int            `json:"user_id"`
	ExpirationDate time.Time      `json:"expiration_date" gorm:""`
	User           User           `gorm:"foreignKey:User_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// CreateLink creates a new link
func CreateLink(redirectURL string, random string, user_id int) error {
	newlink := &Linkly{
		Redirect:       redirectURL,
		Linkly:         random,
		ExpirationDate: expirationDate,
		IsAvailable:    true,
		User_id:        user_id,
	}

	return initializers.DB.Create(newlink).Error
}

// GetLink gets a link by user id
func GetLinkByUser(db *gorm.DB, user_id int) ([]*Linkly, error) {
	links := []*Linkly{}
	err := db.Where("user_id = ?", user_id).Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

// GetLink gets a link by linkly
func GetLink(linkly string) (*Linkly, error) {
	link := &Linkly{}
	err := initializers.DB.Where("linkly= ?", linkly).First(link).Error
	if err != nil {
		return nil, err
	}
	log.Println(link.Redirect)
	return link, nil
}

// UpdateLink updates a link
func UpdateLink(linkly string, redirectURL string) error {
	link, err := GetLink(linkly)
	if err != nil {
		return err
	}
	link.Redirect = redirectURL
	return initializers.DB.Save(link).Error
}

// DeleteLink deletes a link
func DeleteLink(linkly string) error {
	link, err := GetLink(linkly)
	if err != nil {
		return err
	}
	return initializers.DB.Delete(link).Error
}

// UpdateClick updates a link (plus one)
func UpdateClick(linkly string) error {
	link, err := GetLink(linkly)
	if err != nil {
		return err
	}
	link.Clicked += 1
	return initializers.DB.Save(link).Error
}
