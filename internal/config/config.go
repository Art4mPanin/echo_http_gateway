package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	GRPC GRPCConfig `yaml:"grpc"`
	Host string     `yaml:"host"`
}
type GRPCConfig struct {
	InfoPort int           `yaml:"infoport"`
	AuthPort int           `yaml:"authport"`
	Timeout  time.Duration `yaml:"timeout"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
