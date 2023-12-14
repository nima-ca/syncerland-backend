package main

import (
	applicationModels "syncerland/app/application/models"
	userModels "syncerland/app/user/models"
	"syncerland/core/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&userModels.User{})
	initializers.DB.AutoMigrate(&applicationModels.Application{})
}
