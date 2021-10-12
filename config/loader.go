package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func Load(file string, config interface{}) error {

	vp := viper.New()
	vp.SetConfigFile(file)

	if err := vp.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	if err := vp.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshall the config: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %v", err)
	}

	return nil
}
