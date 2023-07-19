package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Setup() {
	dsn := "root@tcp(127.0.0.1:3306)/linkly"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
		return
	}
	fmt.Println("Connecton DB is OK")
	err = db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&Linkly{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("AutoMigrate is OK")
}
