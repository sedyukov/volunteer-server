package routes

import (
	"github.com/gofiber/fiber/v2"

	controllers "github.com/sedyukov/volunteer-server/internal/controllers"
)

var (
	getRefusedRoute   = "/api/v3/disseminators/refused"
	getCtDomainsRoute = "/api/v3/ct-domains/"
	getConfigRoute    = "/api/centralized-config"
)

func SetupGatewayRoutes(app *fiber.App) {
	app.Get(getRefusedRoute, controllers.GetRefused)
	app.Get(getCtDomainsRoute, controllers.GetCtDomains)
	app.Get(getConfigRoute, controllers.GetCentralizedConfig)
}

func GetRefusedRoute() string {
	return getRefusedRoute
}

func GetCtDomainsRoute() string {
	return getCtDomainsRoute
}
