package storage

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/schollz/jsonstore"
)

var (
	ctDomainStorageName = "ct-domains.json.gz"
	ctDomainsKey        = "ct-domains"
)

func initCtDomainsStorage(logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	var domains [1]string

	domains[0] = "www.example.com"

	err := ks.Set(ctDomainsKey, domains)
	if err != nil {
		logger.Error().Msg("Setting example domains to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, ctDomainStorageName)
		if err != nil {
			logger.Error().Msg("Saving to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetCtDomains() []string {
	var domains []string

	ks, err := jsonstore.Open(ctDomainStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(ctDomainsKey, &domains)
	if err != nil {
		panic(err)
	}

	return domains
}

func SetCtDomains(domains []string) {
	ks, err := jsonstore.Open(ctDomainStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Set(ctDomainsKey, &domains)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := jsonstore.Save(ks, ctDomainStorageName)
		if err != nil {
			panic(err)
		}

	}()
	wg.Wait()
}
