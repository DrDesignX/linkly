package initializers

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectoToDB() error{
	var err error
	dsn := os.Getenv("DB_URL")
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
		return err
	}
	fmt.Println("Database connected")
	return nil
}