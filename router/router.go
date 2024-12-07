package router

import (
	"net/http"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/bitcoin-trading-system/bitflyer-api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
	h := handler.NewHandler(cfg)

	r := gin.Default()

	return setUpHandler(r, h)
}

func setUpHandler(r *gin.Engine, h handler.IHandler) *gin.Engine {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthcheck ok!",
		})
	})

	r.GET("/board/:product_code", h.GetBoard)
	r.GET("/ticker/:product_code", h.GetTicker)
	r.GET("/executions/:product_code", h.GetExecutions)
	r.GET("/board_state/:product_code", h.GetBoardState)
	r.GET("/health/:product_code", h.GetHealth)

	return r
}
