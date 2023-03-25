package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"github.com/sedyukov/volunteer-server/internal/remote"
	"github.com/sedyukov/volunteer-server/internal/routes"
	"github.com/sedyukov/volunteer-server/internal/service"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func main() {
	// Load viper config
	err := service.LoadConfig()
	if err != nil {
		panic(err)
	}

	bootstrapApp()
}

func bootstrapApp() {
	// Start logger service
	logger, err := service.NewLogger("volunteer-server", service.LoggerConfig{
		OutOnly: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Logger sucessfully started for gateway")
	app := fiber.New()

	// Initialize JSON storage
	config, err := remote.GetCentralizedConfig(logger)
	if err != nil {
		logger.Error().Msg("Retrieving config failed")
		panic(err)
	}

	storage.SaveCentralizedConfig(config, logger)
	logger.Info().Msg("Config initialization finished")

	// Identifying your public IP
	err = remote.IdentifyPublicIp()
	if err != nil {
		logger.Error().Msg("Identifying IP failed")
		panic(err)
	}
	logger.Info().Msgf("Your public ip is: %s", remote.PublicIp)

	// Initialize JSON storage
	err = storage.Init(logger)
	if err != nil {
		logger.Error().Msg("Storage initialization failed")
		panic(err)
	}
	logger.Info().Msg("Storage initialization finished")

	// Setup gateway routes
	routes.SetupGatewayRoutes(app)

	// Listening for requests
	var port = viper.GetString("GATEWAY_PORT")
	logger.Info().Msgf("Listening to port %v", port)
	app.Listen(":" + port)
}
