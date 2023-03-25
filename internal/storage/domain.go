package storage

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/schollz/jsonstore"
)

type Domain struct {
	CooperationRefused bool   `json:"cooperationRefused"`
	Url                string `json:"url"`
}

var (
	domainStorageName = "domains.json.gz"
	domainsKey        = "domains"
)

func initDomainsStorage(logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	var domains [1]Domain
	exampleDomain := Domain{
		CooperationRefused: false,
		Url:                "www.example.com",
	}
	domains[0] = exampleDomain

	err := ks.Set(domainsKey, domains)
	if err != nil {
		logger.Error().Msg("Setting example domains to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, domainStorageName)
		if err != nil {
			logger.Error().Msg("Saving to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetDomains() []Domain {
	var domains []Domain

	ks, err := jsonstore.Open(domainStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(domainsKey, &domains)
	if err != nil {
		panic(err)
	}

	return domains
}

func SetDomains(domains []Domain) {
	ks, err := jsonstore.Open(domainStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Set(domainsKey, &domains)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := jsonstore.Save(ks, domainStorageName)
		if err != nil {
			panic(err)
		}

	}()
	wg.Wait()
}
