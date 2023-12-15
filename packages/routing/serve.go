package routing

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

// DOC: it bootstraps the Fiber on given port
func Serve(port string) {
	router := GetRouter()

	// Handle Panics
	router.Use(recover.New())

	log.Fatal(router.Listen(port))
}
