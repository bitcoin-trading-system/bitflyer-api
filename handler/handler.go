package handler

import (
	"net/http"
	"net/url"

	"github.com/bitcoin-trading-system/bitflyer-api/config"
	"github.com/bitcoin-trading-system/bitflyer-api/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase usecase.IBitflyerUseCase
	Config  config.Config
}

type IHandler interface {
	GetBoard(ctx *gin.Context)
	GetTicker(ctx *gin.Context)
	GetExecutions(ctx *gin.Context)
	GetBoardState(ctx *gin.Context)
	GetHealth(ctx *gin.Context)
}

func NewHandler(cfg config.Config) IHandler {
	u := usecase.NewBitflyerUseCase(cfg)
	return &Handler{UseCase: u, Config: cfg}
}

func (h *Handler) GetBoard(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	board, err := h.UseCase.GetBoard(productCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, board)
}

func (h *Handler) GetTicker(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	ticker, err := h.UseCase.GetTicker(productCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ticker)
}

type GetExecutionsQueryParams struct {
	ProductCode string `form:"product_code"`
	Count       string `form:"count"`
	Before      string `form:"before"`
	After       string `form:"after"`
}

func NewGetExecutionsQueryParams(qp url.Values) *GetExecutionsQueryParams {
	return &GetExecutionsQueryParams{
		ProductCode: qp.Get("product_code"),
		Count:       qp.Get("count"),
		Before:      qp.Get("before"),
		After:       qp.Get("after"),
	}
}

func (h *Handler) GetExecutions(ctx *gin.Context) {
	qp := NewGetExecutionsQueryParams(ctx.Request.URL.Query())

	executions, err := h.UseCase.GetExecutions(qp.ProductCode, qp.Count, qp.Before, qp.After)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, executions)
}

func (h *Handler) GetBoardState(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	boardState, err := h.UseCase.GetBoardState(productCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, boardState)
}

func (h *Handler) GetHealth(ctx *gin.Context) {
	productCode := ctx.Request.URL.Query().Get("product_code")
	health, err := h.UseCase.GetHealth(productCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, health)
}
