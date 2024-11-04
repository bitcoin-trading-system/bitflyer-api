package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	BaseConfig BaseConfig `toml:"baseConfig"`
}

type BaseConfig struct {
	Port string `toml:"port"`
}

func NewConfig() Config {
	var config Config

	if _, err := toml.DecodeFile("conf/local.toml", &config); err != nil {
		panic(err)
	}

	return config
}
