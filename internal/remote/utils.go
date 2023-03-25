package remote

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sedyukov/volunteer-server/internal/common"
)

type IP struct {
	Query string
}

var client = &http.Client{Timeout: 10 * time.Second}

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

func MakeOwnConfig() {
	// TODO: implement logic to build config based on own ip
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

	PublicIp = ip.Query

	return nil
}
