package api

import (
	"net/http"

	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

type PublicAPI interface {
	GetBoard(productCode string) (models.Board, error)
	GetTicker(productCode string) (models.Ticket, error)
	GetExecutions(productCode string, count, before, after int) ([]models.Execution, error)
	GetBoardState(productCode string) (models.BoardStatus, error)
	GetHealth(productCode string) (models.Health, error)
}

func NewPublicAPI(cfg config.Config) PublicAPI {
	return API{
		BaseUrl:   BaseUrl(cfg.BitflyerConfig.BaseEndPoint),
		ApiKey:    cfg.BitflyerConfig.ApiKey,
		ApiSecret: cfg.BitflyerConfig.ApiSecret,
	}
}

func (api API) GetBoard(productCode string) (models.Board, error) {
	url, err := api.BaseUrl.GetBoardUrl(productCode)
	if err != nil {
		return models.Board{}, err
	}
	resModel := models.Board{}
	if err := api.do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Board{}, err

	}
	return resModel, nil
}

func (api API) GetTicker(productCode string) (models.Ticket, error) {
	url, err := api.BaseUrl.GetTickerUrl(productCode)
	if err != nil {
		return models.Ticket{}, err
	}
	resModel := models.Ticket{}
	if err := api.do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Ticket{}, err
	}
	return resModel, nil
}

func (api API) GetExecutions(productCode string, count, before, after int) ([]models.Execution, error) {
	url, err := api.BaseUrl.GetExecutionsUrl(productCode, count, before, after)
	if err != nil {
		return []models.Execution{}, err
	}
	var resModel []models.Execution
	if err := api.do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return []models.Execution{}, err
	}
	return resModel, nil
}

func (api API) GetBoardState(productCode string) (models.BoardStatus, error) {
	url, err := api.BaseUrl.GetBoardStateUrl(productCode)
	if err != nil {
		return models.BoardStatus{}, err
	}
	resModel := models.BoardStatus{}
	if err := api.do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.BoardStatus{}, err
	}
	return resModel, nil
}

func (api API) GetHealth(productCode string) (models.Health, error) {
	url, err := api.BaseUrl.GetHealthUrl(productCode)
	if err != nil {
		return models.Health{}, err
	}
	resModel := models.Health{}
	if err := api.do(http.MethodGet, nil, &resModel, url, nil); err != nil {
		return models.Health{}, err
	}
	return resModel, nil
}
