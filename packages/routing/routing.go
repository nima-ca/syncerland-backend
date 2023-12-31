package routing

import (
	userRoutes "syncerland/app/user/routes"

	"github.com/gofiber/fiber/v2"
)

// DOC: it creates a router and assign it to GlobalRouter variable
func Init() {
	GlobalRouter = fiber.New()
}

// DOC: it returns the Global Router
func GetRouter() *fiber.App {
	return GlobalRouter
}

// DOC: it registers all the routes in different modules
func RegisterRoutes() {
	router := GetRouter()

	userRoutes.Register(router)
}
