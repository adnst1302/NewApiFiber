package boot

import (
	"NewApiFiber/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartApiService() {
	app := fiber.New()
	router.RouteList(app)
	log.Fatal(app.Listen(":3000"))
}
