package remote

import (
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	// urls to fetch config
	configUrlAmazonS3         = "https://censortracker.s3.eu-central-1.amazonaws.com/config.json"
	configUrlAmazonCloudfront = "https://d204gfm9dw21wi.cloudfront.net/"
	configUrlAmazonGoogleApis = "https://storage.googleapis.com/censortracker/config.json"
)

func GetCentralizedConfig() (common.ConfigResponse, error) {
	resp, err := fetchCentralizedConfig()

	if err != nil {
		logger.Error().Msg("Faild while fetching config from centralized server")
		return resp, err
	}

	centralizedConfig = resp

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

// for now it's only for Russia
func GetCtDomains() ([]string, error) {
	url, err := getCtDomainUrl(centralizedConfig)

	if err != nil {
		logger.Error().Msg("Faild while getting ct domain url")
		return nil, err
	}

	resp := new([]string)

	err = getJson(url, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil
}

// currently is the same for all countries (so now will take for Russia)
func GetRefused() ([]common.Refused, error) {
	url, err := getRefusedDomainUrl(centralizedConfig)

	if err != nil {
		logger.Error().Msg("Faild while getting refused domain url")
		return nil, err
	}

	resp := new([]common.Refused)

	err = getJson(url, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil
}
