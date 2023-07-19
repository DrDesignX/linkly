package model

import(
	"fmt"
	"log"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Linkly struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Redirect       string    `json:"redirect" gorm:"not null"`
	Linkly         string    `json:"linkly" gorm:"not null;unique"`
	Clicked        int       `json:"clicked"`
	IsAvailable    bool      `json:"is_available"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserID         int       `json:"user_id"`
	ExpirationDate time.Time `json:"expiration_date"`

	User User `gorm:"foreignKey:UserID"` 
}


type User struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"not null;unique"`
	Email       string    `json:"email" gorm:"not null;unique"`
	Password    string    `json:"password" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
	IsAdmin     bool      `json:"is_admin"`
}


func Setup(){
	dsn := "root@tcp(127.0.0.1:3306)/linkly"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
