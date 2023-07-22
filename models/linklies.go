package models

import (
	"time"
	"gorm.io/gorm"
)

var(
	expirationDate = time.Now().Add(30 * 24 * time.Hour)
)
type Linkly struct {
	gorm.Model
	Redirect       string    `json:"redirect" gorm:"not null"`
	Linkly         string    `json:"linkly" gorm:"not null;unique"`
	Clicked        int       `json:"clicked" gorm:"default:0"`
	IsAvailable    bool      `json:"is_available" gorm:"default:true"`
	User_id        int       `json:"user_id"`
	ExpirationDate time.Time `json:"expiration_date" gorm:""`
	User           User      `gorm:"foreignKey:User_id"`
}


// CreateLink creates a new link
func CreateLink(db *gorm.DB, redirectURL string, random string, user_id int) error {
	newlink := &Linkly{
		Redirect:       redirectURL,
		Linkly:         random,
		ExpirationDate: ExpirationDate,
		IsAvailable:    true,
		User_id:        user_id,
	}

	return db.Create(newlink).Error
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
func GetLink(db *gorm.DB, linkly string) (*Linkly, error) {
	link := &Linkly{}
	err := db.Where("linkly = ?", linkly).First(link).Error
	if err != nil {
		return nil, err
	}
	return link, nil
}

// UpdateLink updates a link
func UpdateLink(db *gorm.DB, linkly string, redirectURL string) error {
	link, err := GetLink(db, linkly)
	if err != nil {
		return err
	}
	link.Redirect = redirectURL
	return db.Save(link).Error
}

// DeleteLink deletes a link
func DeleteLink(db *gorm.DB, linkly string) error {
	link, err := GetLink(db, linkly)
	if err != nil {
		return err
	}
	return db.Delete(link).Error
}

// UpdateClick updates a link (plus one)
func UpdateClick(db *gorm.DB, linkly string) error {
	link, err := GetLink(db, linkly)
	if err != nil {
		return err
	}
	link.Clicked += 1
	return db.Save(link).Error
}