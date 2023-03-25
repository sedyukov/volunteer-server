package storage

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/schollz/jsonstore"
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	centralizedConfigStorageName = "config-centralized.json.gz"
	centralizeConfigKey          = "config-centralized"
)

func SaveCentralizedConfig(config common.ConfigResponse, logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	err := ks.Set(centralizeConfigKey, config)
	if err != nil {
		logger.Error().Msg("Setting centralized config to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, centralizedConfigStorageName)
		if err != nil {
			logger.Error().Msg("Saving centralized config to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetCentralizedConfig() common.ConfigResponse {
	var config common.ConfigResponse

	ks, err := jsonstore.Open(centralizedConfigStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(centralizeConfigKey, &config)
	if err != nil {
		panic(err)
	}

	return config
}
