package service

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
