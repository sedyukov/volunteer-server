package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sedyukov/volunteer-server/internal/remote"
	"github.com/sedyukov/volunteer-server/internal/routes"
	"github.com/sedyukov/volunteer-server/internal/service"
	"github.com/sedyukov/volunteer-server/internal/storage"
	"github.com/spf13/viper"
)

func bootstrapApp() {
	// Start logger service
	logger, err := service.NewLogger("volunteer-server", service.LoggerConfig{
		OutOnly: true,
	})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("Logger sucessfully started")
	app := fiber.New()

	// Initialize JSON storage
	err = storage.Init(logger)
	if err != nil {
		logger.Error().Msg("Storage initialization failed")
		panic(err)
	}
	logger.Info().Msg("Storage initialization finished")

	// Init remote package
	remote.Init(logger)

	// Fetch and save config
	config, err := remote.GetCentralizedConfig()
	if err != nil {
		logger.Error().Msg("Retrieving config failed")
		panic(err)
	}

	storage.SaveCentralizedConfig(config, logger)
	logger.Info().Msg("Config initialization finished")

	// Fetch and save ct-domains
	domains, err := remote.GetCtDomains()
	if err != nil {
		logger.Error().Msg("Retrieving ct-domains failed")
		panic(err)
	}

	storage.SetCtDomains(domains)
	logger.Info().Msg("Ct-domains initialization finished")

	// Fetch and save refused
	refused, err := remote.GetRefused()
	if err != nil {
		logger.Error().Msg("Retrieving refused failed")
		panic(err)
	}

	storage.SetRefused(refused)
	logger.Info().Msg("Refused initialization finished")

	// Identifying your public IP
	err = remote.IdentifyPublicIp()
	if err != nil {
		logger.Error().Msg("Identifying IP failed")
		panic(err)
	}
	logger.Info().Msgf("Your public ip: %s", remote.GetPubleIp())

	// Setup gateway routes
	routes.SetupGatewayRoutes(app)

	// Listening for requests
	var port = viper.GetString("GATEWAY_PORT")
	logger.Info().Msgf("Listening to port %v", port)
	app.Listen(":" + port)
}
