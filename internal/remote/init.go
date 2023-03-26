package remote

import "github.com/rs/zerolog"

func Init(loggerInstance zerolog.Logger) {
	logger = loggerInstance
}
