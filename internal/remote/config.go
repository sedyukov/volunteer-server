package remote

import "github.com/sedyukov/volunteer-server/internal/common"

var (
	// urls to fetch config
	configUrlAmazonS3         = "https://censortracker.s3.eu-central-1.amazonaws.com/config.json"
	configUrlAmazonCloudfront = "https://d204gfm9dw21wi.cloudfront.net/"
	configUrlAmazonGoogleApis = "https://storage.googleapis.com/censortracker/config.json"
)

func GetExternalConfig() (common.ConfigResponse, error) {
	resp, err := fetchExternalConfig()

	if err != nil {
		logger.Error().Msg("Faild while fetching config from centralized server")
		return resp, err
	}

	externalConfig = resp

	return resp, nil
}

func fetchExternalConfig() (common.ConfigResponse, error) {
	resp := new(common.ConfigResponse)

	centralizedFetching = true

	err := getJson(configUrlAmazonS3, resp)

	if err == nil {
		return *resp, nil
	}

	logger.Info().Msg("Failed configUrlAmazonS3")

	err = getJson(configUrlAmazonCloudfront, resp)

	if err == nil {
		return *resp, nil
	}

	logger.Info().Msg("Failed configUrlAmazonGoogleApis")

	err = getJson(configUrlAmazonGoogleApis, resp)

	if err == nil {
		return *resp, nil
	}

	logger.Info().Msg("Failed configUrlAmazonGoogleApis")

	// p2p config fetching
	centralizedFetching = false

	url := getConfigPeerEndpoint(common.PeerUrl)
	err = getJson(url, resp)

	if err == nil {
		return *resp, nil
	}

	logger.Info().Msg("Failed fetch peer config")

	return *resp, err
}
