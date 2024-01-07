package migration

import (
	"fmt"
	"syncerland/core/initializers"
	"syncerland/models"
)

func MigrateDB() {
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
