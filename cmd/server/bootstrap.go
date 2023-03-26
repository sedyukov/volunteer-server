package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"github.com/sedyukov/volunteer-server/internal/common"
	"github.com/sedyukov/volunteer-server/internal/remote"
	"github.com/sedyukov/volunteer-server/internal/routes"
	"github.com/sedyukov/volunteer-server/internal/service"
	"github.com/sedyukov/volunteer-server/internal/storage"
)

func bootstrapApp() {
	// Get port from config
	var port = viper.GetString("SERVER_PORT")
	common.Port = port

	var peerUrl = viper.GetString("PEER_URL")
	common.PeerUrl = peerUrl

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

	// Fetch and save centralized config
	config, err := remote.GetExternalConfig()
	if err != nil {
		logger.Error().Msg("Retrieving config failed")
		panic(err)
	}

	storage.SaveExternalConfig(config, logger)
	logger.Info().Msg("Config initialization finished")

	// Fetch and save ct-domains
	domains, err := remote.GetCtDomains()
	if err != nil {
		logger.Error().Msg("Retrieving ct-domains failed")

		if remote.IsCentralizedFetching() {
			domains, err = remote.GetCtDomainsFromPeer()

			if err != nil {
				logger.Error().Msg("Retrieving ct-domains from peer failed")
				panic(err)
			}
		} else {
			panic(err)
		}

	}

	storage.SetCtDomains(domains)
	logger.Info().Msg("Ct-domains initialization finished")

	// Fetch and save refused
	refused, err := remote.GetRefused()
	if err != nil {
		logger.Error().Msg("Retrieving refused failed")

		if remote.IsCentralizedFetching() {
			refused, err = remote.GetRefusedFromPeer()

			if err != nil {
				logger.Error().Msg("Retrieving refused from peer failed")
				panic(err)
			}
		} else {
			panic(err)
		}
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

	// Convert and save own config
	ownConfig := remote.GetOwnConfig()
	storage.SaveOwnConfig(ownConfig, logger)
	logger.Info().Msg("Own config initialization finished")

	// Setup gateway routes
	routes.SetupGatewayRoutes(app)

	logger.Info().Msgf("Listening to port %v", common.Port)
	app.Listen(":" + port)
}
