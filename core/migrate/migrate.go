package main

import (
	"fmt"
	"syncerland/core/initializers"
	"syncerland/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Application{},
		&models.Country{},
		&models.Interview{},
		&models.Interviewer{},
		&models.Job{},
		&models.Note{},
		&models.Offer{},
	)

	fmt.Println("Migration was successful")
}
