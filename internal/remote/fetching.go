package remote

import (
	"github.com/sedyukov/volunteer-server/internal/common"
)

// for now it's only for Russia
func GetCtDomains() ([]string, error) {
	url, err := getCtDomainUrl(externalConfig)

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

// currently is the same for all countries (for now will take for Russia)
func GetRefused() ([]common.Refused, error) {
	url, err := getRefusedDomainUrl(externalConfig)

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
