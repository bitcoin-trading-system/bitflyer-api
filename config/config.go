package config

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"os"
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

	return config
}
