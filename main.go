package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthcheck ok!",
		})
	})

	err := router.Run(fmt.Sprintf(":%s", cfg.BaseConfig.Port))
	if err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
