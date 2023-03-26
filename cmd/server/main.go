package main

import (
	"github.com/sedyukov/volunteer-server/internal/service"
)

func main() {
	// Load viper config
	err := service.LoadConfig()
	if err != nil {
		panic(err)
	}

	bootstrapApp()
}
