package routes

import (
	"github.com/gofiber/fiber/v2"

	controllers "github.com/sedyukov/volunteer-server/internal/controllers"
)

func SetupGatewayRoutes(app *fiber.App) {
	app.Get("/api/blocked-domains", controllers.GetBlockedDomains)
	app.Get("/api/centralized-config", controllers.GetCentralizedConfig)
}
