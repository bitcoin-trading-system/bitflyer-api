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
	// Public API
	GetBoard(ctx *gin.Context)
	GetTicker(ctx *gin.Context)
	GetExecutions(ctx *gin.Context)
	GetBoardState(ctx *gin.Context)
	GetHealth(ctx *gin.Context)

	// Private API
	GetBalance(ctx *gin.Context)
	GetCollateral(ctx *gin.Context)
	PostSendChildOrder(ctx *gin.Context)
	PostCancelChildOrder(ctx *gin.Context)
	GetChildOrders(ctx *gin.Context)
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

func (h *Handler) GetBalance(ctx *gin.Context) {
	balance, err := h.UseCase.GetBalance()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, balance)
}

func (h *Handler) GetCollateral(ctx *gin.Context) {
	collateral, err := h.UseCase.GetCollateral()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, collateral)
}

func (h *Handler) PostSendChildOrder(ctx *gin.Context) {
	req, err := NewPostSendChildOrderRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isDry := ctx.Request.URL.Query().Get("dry_run") == "1"
	childOrder, err := h.UseCase.PostSendChildOrder(req.ProductCode, req.ChildOrderType, req.Side, req.Price, req.Size, req.MinuteToExpire, req.TimeInForce, isDry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, childOrder)
}

type PostSendChildOrderRequestBody struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          int     `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int     `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

func NewPostSendChildOrderRequestBody(ctx *gin.Context) (PostSendChildOrderRequestBody, error) {
	var req PostSendChildOrderRequestBody
	err := ctx.BindJSON(&req)
	return req, err
}

func (h *Handler) PostCancelChildOrder(ctx *gin.Context) {
	req, err := NewCancelChildOrderRequestBody(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isDry := ctx.Request.URL.Query().Get("dry_run") == "1"
	err = h.UseCase.PostCancelChildOrder(req.ProductCode, req.ChildOrderID, isDry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

type CancelChildOrderRequestBody struct {
	ProductCode  string `json:"product_code"`
	ChildOrderID string `json:"child_order_id"`
}

func NewCancelChildOrderRequestBody(ctx *gin.Context) (CancelChildOrderRequestBody, error) {
	var req CancelChildOrderRequestBody
	err := ctx.BindJSON(&req)
	return req, err
}

func (h *Handler) GetChildOrders(ctx *gin.Context) {
	childOrders, err := h.UseCase.GetChildOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, childOrders)
}
