package storage

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/schollz/jsonstore"
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	externalConfigStorageName = "config-external.json.gz"
	externalConfigKey         = "config-external"
	ownConfigStorageName      = "config.json.gz"
	ownConfigKey              = "config"
)

func SaveExternalConfig(config common.ConfigResponse, logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	err := ks.Set(externalConfigKey, config)
	if err != nil {
		logger.Error().Msg("Setting external config to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, externalConfigStorageName)
		if err != nil {
			logger.Error().Msg("Saving external config to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetExternalConfig() common.ConfigResponse {
	var config common.ConfigResponse

	ks, err := jsonstore.Open(externalConfigStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(externalConfigKey, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func SaveOwnConfig(config common.ConfigResponse, logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	err := ks.Set(ownConfigKey, config)
	if err != nil {
		logger.Error().Msg("Setting centralized config to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, ownConfigStorageName)
		if err != nil {
			logger.Error().Msg("Saving centralized config to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetOwnConfig() common.ConfigResponse {
	var config common.ConfigResponse

	ks, err := jsonstore.Open(ownConfigStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(ownConfigKey, &config)
	if err != nil {
		panic(err)
	}

	return config
}
