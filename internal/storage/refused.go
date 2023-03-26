package storage

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/schollz/jsonstore"
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	refusedStorageName = "refused.json.gz"
	refusedKey         = "refused"
)

func initRefusedStorage(logger zerolog.Logger) error {
	ks := new(jsonstore.JSONStore)

	var refused [1]common.Refused
	exampleRefused := common.Refused{
		CooperationRefused: false,
		Url:                "www.example.com",
	}
	refused[0] = exampleRefused

	err := ks.Set(refusedKey, refused)
	if err != nil {
		logger.Error().Msg("Setting example refused to storage failed")
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() error {
		defer wg.Done()

		err := jsonstore.Save(ks, refusedStorageName)
		if err != nil {
			logger.Error().Msg("Saving to storage failed")
			return err
		}

		return nil
	}()
	wg.Wait()

	return nil
}

func GetRefused() []common.Refused {
	var refused []common.Refused

	ks, err := jsonstore.Open(refusedStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Get(refusedKey, &refused)
	if err != nil {
		panic(err)
	}

	return refused
}

func SetRefused(refused []common.Refused) {
	ks, err := jsonstore.Open(refusedStorageName)

	if err != nil {
		panic(err)
	}

	err = ks.Set(refusedKey, &refused)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := jsonstore.Save(ks, refusedStorageName)
		if err != nil {
			panic(err)
		}

	}()
	wg.Wait()
}
