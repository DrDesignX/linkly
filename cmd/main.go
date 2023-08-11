package main

import (
	"log"
	"os"

	"github.com/drdesignx/linkly/initializers"
	"github.com/drdesignx/linkly/migrate"
	"github.com/drdesignx/linkly/routes"
)

func Initialize() error {
	if err := initializers.LoadEnvVariables(); err != nil {
		log.Println("Initialization Failed " + err.Error())
		return err
	}
	if err := initializers.ConnectToDB(); err != nil {
		log.Println("Initialization Failed " + err.Error())
		return err
	}
	if err := migrate.RunMigrate(); err != nil {
		log.Println("Initialization Failed " + err.Error())
		return err
	}
	return nil
}

func main() {
	err := Initialize()
	if err != nil {
		log.Println("Initialization Failed " + err.Error())
		os.Exit(1)
	}
	if err := routes.Routes("3000"); err != nil {
		log.Println("Initialization Failed " + err.Error())
	}
}
