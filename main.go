package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/bitcoin-trading-system/bitflyer-api/router"
)

func main() {
	tomlFilePath := flag.String("conf", "conf/local.toml", "tomlファイルの名前")
	envFilePath := flag.String("env", "env/.env.local", "envファイルのパス")
	flag.Parse()

	cfg := config.NewConfig(*tomlFilePath, *envFilePath)

	r := router.NewRouter(cfg)

	if err := r.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port)); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
