package remote

import "github.com/sedyukov/volunteer-server/internal/common"

// for now it's only for Russia
func GetCtDomainsFromPeer() ([]string, error) {
	url := getCtDomainsEndpoint(common.PeerUrl)

	resp := new([]string)

	err := getJson(url, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil
}

// currently is the same for all countries (for now will take for Russia)
func GetRefusedFromPeer() ([]common.Refused, error) {
	url := getRefusedEndpoint(common.PeerUrl)

	resp := new([]common.Refused)

	err := getJson(url, resp)

	if err != nil {
		return nil, err
	}

	return *resp, nil
}
