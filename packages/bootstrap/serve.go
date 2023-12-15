package bootstrap

import (
	"os"
	"syncerland/core/initializers"
	"syncerland/packages/routing"
)

func Serve() {
	// DOC: initialize the app
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	port := os.Getenv("PORT")

	// DOC: initialize the router (Fiber)
	routing.Init()
	routing.RegisterRoutes()
	routing.Serve(":" + port)
}
