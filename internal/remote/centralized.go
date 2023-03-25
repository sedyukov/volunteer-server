package remote

import (
	"github.com/rs/zerolog"
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	// urls to fetch config
	configUrlAmazonS3         = "https://censortracker.s3.eu-central-1.amazonaws.com/config.json"
	configUrlAmazonCloudfront = "https://d204gfm9dw21wi.cloudfront.net/"
	configUrlAmazonGoogleApis = "https://storage.googleapis.com/censortracker/config.json"
)

func GetCentralizedConfig(logger zerolog.Logger) (common.ConfigResponse, error) {
	resp, err := fetchCentralizedConfig()

	if err != nil {
		logger.Error().Msg("Faild while fetching config from centralized server")
		return resp, err
	}

	return resp, nil
}

func fetchCentralizedConfig() (common.ConfigResponse, error) {
	resp := new(common.ConfigResponse)

	err := getJson(configUrlAmazonS3, resp)

	if err == nil {
		return *resp, nil
	}

	err = getJson(configUrlAmazonCloudfront, resp)

	if err == nil {
		return *resp, nil
	}

	err = getJson(configUrlAmazonGoogleApis, resp)

	if err == nil {
		return *resp, nil
	}

	return *resp, err
}
