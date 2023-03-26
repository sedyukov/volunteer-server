package storage

import "github.com/rs/zerolog"

func Init(logger zerolog.Logger) error {
	err := initRefusedStorage(logger)
	if err != nil {
		return err
	}

	err = initCtDomainsStorage(logger)
	if err != nil {
		return err
	}

	return nil
}
