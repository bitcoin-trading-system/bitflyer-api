package usecase

import (
	"github.com/bitcoin-trading-system/bitflyer-api/api"
	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

type IBitflyerUseCase interface {
	GetBoard(productCode string) (models.Board, error)
	GetTicker(productCode string) (models.Ticket, error)
	GetExecutions(productCode string, count, before, after int) ([]models.Execution, error)
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
	return bu.PublicAPI.GetBoard(productCode)
}

func (bu *BitflyerUseCase) GetTicker(productCode string) (models.Ticket, error) {
	return bu.PublicAPI.GetTicker(productCode)
}

func (bu *BitflyerUseCase) GetExecutions(productCode string, count, before, after int) ([]models.Execution, error) {
	return bu.PublicAPI.GetExecutions(productCode, count, before, after)
}

func (bu *BitflyerUseCase) GetBoardState(productCode string) (models.BoardStatus, error) {
	return bu.PublicAPI.GetBoardState(productCode)
}

func (bu *BitflyerUseCase) GetHealth(productCode string) (models.Health, error) {
	return bu.PublicAPI.GetHealth(productCode)
}
