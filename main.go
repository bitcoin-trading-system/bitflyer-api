package main

import (
	"fmt"
	"log"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/bitcoin-trading-system/bitflyer-api/router"
)

func main() {
	cfg := config.NewConfig()

	router := router.NewRouter(cfg)

	if err := router.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port)); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
