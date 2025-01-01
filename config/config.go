package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	BaseConfig     BaseConfig     `toml:"baseConfig"`
	BitflyerConfig BitFlyerConfig `toml:"bitFlyerConfig"`
}

type BaseConfig struct {
	Port string `toml:"port"`
}

type BitFlyerConfig struct {
	ApiKey       string
	ApiSecret    string
	BaseEndPoint string `toml:"baseEndPoint"`
}

func NewConfig(tomlFilePath, envFilePath string) Config {
	var config Config

	if _, err := toml.DecodeFile(tomlFilePath, &config); err != nil {
		panic(err)
	}

	if err := godotenv.Load(envFilePath); err != nil {
		panic(err)
	}

	config.BitflyerConfig.ApiKey = os.Getenv("API_KEY")
	config.BitflyerConfig.ApiSecret = os.Getenv("API_SECRET")

	if err := config.mustCheck(); err != nil {
		panic(err)
	}

	return config
}

func (cfg Config) mustCheck() error {
	if cfg.BaseConfig.Port == "" {
		return errors.New("port is empty")
	}
	if cfg.BitflyerConfig.ApiKey == "" {
		return errors.New("apiKey is empty")
	}
	if cfg.BitflyerConfig.ApiSecret == "" {
		return errors.New("apiSecret is empty")
	}
	if cfg.BitflyerConfig.BaseEndPoint == "" {
		return errors.New("baseEndPoint is empty")
	}
	return nil

}
