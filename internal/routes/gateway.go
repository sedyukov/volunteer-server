package routes

import (
	"github.com/gofiber/fiber/v2"

	controllers "github.com/sedyukov/volunteer-server/internal/controllers"
)

var (
	getRefusedRoute        = "/api/v3/disseminators/refused"
	getCtDomainsRoute      = "/api/v3/ct-domains/"
	getExternalConfigRoute = "/api/centralized-config"
	getOwnConfigRoute      = "/api/config"
)

func SetupGatewayRoutes(app *fiber.App) {
	app.Get(getRefusedRoute, controllers.GetRefused)
	app.Get(getCtDomainsRoute, controllers.GetCtDomains)
	app.Get(getExternalConfigRoute, controllers.GetexternalConfig)
	app.Get(getOwnConfigRoute, controllers.GetOwnConfig)
}

func GetRefusedRoute() string {
	return getRefusedRoute
}

func GetCtDomainsRoute() string {
	return getCtDomainsRoute
}

func GetExternalConfigRoute() string {
	return getExternalConfigRoute
}
