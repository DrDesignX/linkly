package migrate

import (
	"log"

	"github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/models"
)

func initializer() {
	if err := initializers.LoadEnvVariables(); err != nil {
		log.Fatal(err)
	}
	if err := initializers.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
}

func RunMigrate() error {
	if err := initializers.DB.AutoMigrate(&models.Linkly{}, &models.User{}); err != nil {
		return err
	}

	return nil
}
