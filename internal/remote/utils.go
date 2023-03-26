package remote

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sedyukov/volunteer-server/internal/common"
	"github.com/sedyukov/volunteer-server/internal/routes"
)

type IP struct {
	Query string
}

func getJson(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetLocalesFromConfig(config common.ConfigResponse) []string {
	dataArray := config.Data
	var locales []string

	for i, value := range dataArray {
		locales[i] = value.CountryName
	}

	return locales
}

func getCtDomainUrl(config common.ConfigResponse) (string, error) {
	dataArray := config.Data

	for _, value := range dataArray {
		if value.CountryCode == russiaConrtyCode {
			return value.RegistryURL, nil
		}
	}

	return "", errors.New("CtDomains url not found")
}

func getRefusedDomainUrl(config common.ConfigResponse) (string, error) {
	dataArray := config.Data

	for _, value := range dataArray {
		if value.CountryCode == russiaConrtyCode {
			return value.Specifics.CooperationRefusedORIURL, nil
		}
	}

	return "", errors.New("CtDomains url not found")
}

func GetOwnConfig() common.ConfigResponse {
	convertedConfig := ConvertConfig(publicIp, externalConfig)
	return convertedConfig
}

func ConvertConfig(host string, config common.ConfigResponse) common.ConfigResponse {
	var convertedConfig common.ConfigResponse

	convertedConfig.Meta = config.Meta

	dataArray := config.Data

	for _, value := range dataArray {
		var tmp common.Data

		tmp = value

		if value.CountryCode == russiaConrtyCode {
			tmp.RegistryURL = getOwnCtDomainsEndpoint()
		}

		tmp.Specifics.CooperationRefusedORIURL = getOwnRefusedEndpoint()

		convertedConfig.Data = append(convertedConfig.Data, tmp)
	}

	return convertedConfig
}

func IdentifyPublicIp() error {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	var ip IP
	json.Unmarshal(body, &ip)

	publicIp = ip.Query

	return nil
}

func getOwnRefusedEndpoint() string {
	return getRefusedEndpoint(GetOwnIpWithPort())
}

func getRefusedEndpoint(url string) string {
	return url + routes.GetRefusedRoute()
}

func getOwnCtDomainsEndpoint() string {
	return getCtDomainsEndpoint(GetOwnIpWithPort())
}

func getCtDomainsEndpoint(url string) string {
	return url + routes.GetCtDomainsRoute()
}

func getConfigPeerEndpoint(url string) string {
	return url + routes.GetExternalConfigRoute()
}
func GetOwnIpWithPort() string {
	hostWithProtocol := "http://" + publicIp

	if common.Port == "80" {
		return hostWithProtocol
	}

	return hostWithProtocol + ":" + common.Port
}
