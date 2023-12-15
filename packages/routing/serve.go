package routing

// DOC: it bootstraps the Fiber on given port
func Serve(port string) {
	router := GetRouter()

	router.Listen(port)
}
