package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Setup() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/linkly"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(&User{}, &Linkly{})
	if err != nil {
		log.Panic(err)
	}
	return db
}

func CreateLink(db *gorm.DB, redirectURL string, random string, user_id int) error {
	newlink := &Linkly{
		Redirect:       redirectURL,
		Linkly:         "fuc",
		ExpirationDate: ExpirationDate,
		IsAvailable:    true,
		User_id:        user_id,
	}

	return db.Create(newlink).Error
}
