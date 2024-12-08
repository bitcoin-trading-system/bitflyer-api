package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/bitcoin-trading-system/bitflyer-api/router"
)

func main() {
	comlFilePath := flag.String("conf", "local.toml", "設定ファイルの名前")
	flag.Parse()

	cfg := config.NewConfig(*comlFilePath)

	r := router.NewRouter(cfg)

	if err := r.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port)); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
