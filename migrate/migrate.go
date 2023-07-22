package migrate

import (
	"log"

	"github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/models"
)


func initializer(){
	if err := initializers.LoadEnvVariable(); err != nil {
		log.Fatal(err)
	}
	if err := initializers.ConnectoToDB(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initializers.DB.AutoMigrate(&models.Linkly{}, &models.User{})
}