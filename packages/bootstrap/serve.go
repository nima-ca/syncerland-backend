package bootstrap

import (
	"os"
	"syncerland/core/initializers"
	"syncerland/core/migration"
	"syncerland/packages/routing"
	"syncerland/packages/validators"
)

func Serve() {
	// DOC: initialize the app
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migration.MigrateDB()

	validators.RegisterValidators()

	port := os.Getenv("PORT")

	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterMiddlewares()
	routing.RegisterRoutes()

	routing.Serve(":" + port)
}
