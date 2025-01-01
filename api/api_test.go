package api

import (
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)


var APITestConfig config.Config

func init() {
	APITestConfig = config.NewConfig("../conf/local.toml", "../env/.env.local")
}
