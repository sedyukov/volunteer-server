package storage

import "github.com/rs/zerolog"

func Init(logger zerolog.Logger) error {
	err := initDomainsStorage(logger)
	if err != nil {
		return err
	}

	return nil
}
