package remote

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/sedyukov/volunteer-server/internal/common"
)

var (
	logger            zerolog.Logger
	centralizedConfig common.ConfigResponse
	publicIp          string
	client            = &http.Client{Timeout: 20 * time.Second}
	russiaConrtyCode  = "RU"
	OwnConfig         common.ConfigResponse
)

func GetPubleIp() string {
	return publicIp
}
