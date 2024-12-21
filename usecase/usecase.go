package usecase

import (
	"fmt"
	"slices"

	"github.com/bitcoin-trading-system/bitflyer-api/api"
	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

const (
	ProductCodeBTCJPY = "BTC_JPY"
	ProductCodeXRPJPY = "XRP_JPY"
	ProductCodeETHJPY = "ETH_JPY"
	ProductCodeXLMJPY = "XLM_JPY"
	ProductCOdeMONAJPY = "MONA_JPY"

	ProductCodeETHBTC = "ETH_BTC"
	ProductCodeBCHBTC = "BCH_BTC"
	ProductCodeFXBTCJPY = "FX_BTC_JPY"

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
	if !validateProductCode(productCode) {
		return models.Board{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetBoard(productCode)
}

func (bu *BitflyerUseCase) GetTicker(productCode string) (models.Ticket, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Ticket{}, fmt.Errorf("invalid product code: %s", productCode)
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

	if !validateProductCode(productCode) {
		return []models.Execution{}, fmt.Errorf("invalid product code: %s", productCode)
	}

	return bu.PublicAPI.GetExecutions(productCode, count, before, after)
}

func (bu *BitflyerUseCase) GetBoardState(productCode string) (models.BoardStatus, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.BoardStatus{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetBoardState(productCode)
}

func (bu *BitflyerUseCase) GetHealth(productCode string) (models.Health, error) {
	if productCode == "" {
		productCode = ProductCodeBTCJPY
	}
	if !validateProductCode(productCode) {
		return models.Health{}, fmt.Errorf("invalid product code: %s", productCode)
	}
	return bu.PublicAPI.GetHealth(productCode)
}

func validateProductCode(productCode string) bool {
	allowProductCodes := []string{
		ProductCodeBTCJPY,
		ProductCodeXRPJPY,
		ProductCodeETHJPY,
		ProductCodeXLMJPY,
		ProductCOdeMONAJPY,
		ProductCodeETHBTC,
		ProductCodeBCHBTC,
		ProductCodeFXBTCJPY,
	}

	return slices.Contains(allowProductCodes, productCode)
}
