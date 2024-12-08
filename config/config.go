package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	BaseConfig     BaseConfig     `toml:"baseConfig"`
	BitflyerConfig BitFlyerConfig `toml:"bitFlyerConfig"`
}

type BaseConfig struct {
	Port string `toml:"port"`
}

type BitFlyerConfig struct {
	ApiKey       string `toml:"apiKey"`
	ApiSecret    string `toml:"apiSecret"`
	BaseEndPoint string `toml:"baseEndPoint"`
}

func NewConfig(tomlFilePath string) Config {
	var config Config

	if _, err := toml.DecodeFile(tomlFilePath, &config); err != nil {
		panic(err)
	}

	return config
}
