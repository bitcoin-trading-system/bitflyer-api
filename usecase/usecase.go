package usecase

import (
	"github.com/bitcoin-trading-system/bitflyer-api/api"
	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

const (
	ProductCodeBTCJPY = "BTC_JPY"

	GetExecutionsDefaultCount = "100"
)

type IBitflyerUseCase interface {
	GetBoard(productCode string) (models.Board, error)
	GetTicker(productCode string) (models.Ticket, error)
	GetExecutions(productCode, count, before, after string) ([]models.Execution, error)
	GetBoardState(productCode string) (models.BoardStatus, error)
	GetHealth(productCode string) (models.Health, error)
}

type BitflyerUseCase struct {
	PublicAPI api.PublicAPI
	Config    config.Config
}

func NewBitflyerUseCase(cfg config.Config) IBitflyerUseCase {
	publicAPI := api.NewPublicAPI(cfg)
	return &BitflyerUseCase{PublicAPI: publicAPI, Config: cfg}
}

func (bu *BitflyerUseCase) GetBoard(productCode string) (models.Board, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	return bu.PublicAPI.GetBoard(productCode)
}

func (bu *BitflyerUseCase) GetTicker(productCode string) (models.Ticket, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	return bu.PublicAPI.GetTicker(productCode)
}

func (bu *BitflyerUseCase) GetExecutions(productCode string, count, before, after string) ([]models.Execution, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if count == "" {
		count = GetExecutionsDefaultCount
	}

	return bu.PublicAPI.GetExecutions(productCode, count, before, after)
}

func (bu *BitflyerUseCase) GetBoardState(productCode string) (models.BoardStatus, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	return bu.PublicAPI.GetBoardState(productCode)
}

func (bu *BitflyerUseCase) GetHealth(productCode string) (models.Health, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	return bu.PublicAPI.GetHealth(productCode)
}
