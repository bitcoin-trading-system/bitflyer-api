package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bitcoin-trading-system/bitflyer-api/api/models"
	"github.com/bitcoin-trading-system/bitflyer-api/config"
)

type PrivateAPI interface {
	GetBalance() ([]models.Balance, error)
	GetCollateral() (models.Collateral, error)

	PostSendChildOrder(req models.SendChildOrderRequest, isDry bool) (models.ChildOrder, error)
	PostCancelChildOrder(req models.CancelChildOrderRequest, isDry bool) error
	GetChildOrders() ([]models.ChildOrder, error)
}

func NewPrivateAPI(cfg config.Config) PrivateAPI {
	return &API{
		BaseUrl:   BaseUrl(cfg.BitflyerConfig.BaseEndPoint),
		ApiKey:    cfg.BitflyerConfig.ApiKey,
		ApiSecret: cfg.BitflyerConfig.ApiSecret,
	}
}

func (api *API) GetBalance() ([]models.Balance, error) {
	url, err := api.BaseUrl.GetBalanceUrl()
	if err != nil {
		return []models.Balance{}, err
	}

	var resModel []models.Balance
	if err := api.do(http.MethodGet, nil, &resModel, url, nil, true); err != nil {
		return []models.Balance{}, err
	}
	return resModel, nil
}

func (api *API) GetCollateral() (models.Collateral, error) {
	url, err := api.BaseUrl.GetCollateralUrl()
	if err != nil {
		return models.Collateral{}, err
	}
	resModel := models.Collateral{}
	if err := api.do(http.MethodGet, nil, &resModel, url, nil, true); err != nil {
		return models.Collateral{}, err
	}
	return resModel, nil
}

func (api *API) PostSendChildOrder(req models.SendChildOrderRequest, isDry bool) (models.ChildOrder, error) {
	url, err := api.BaseUrl.PostSendChildOrderUrl()
	if err != nil {
		return models.ChildOrder{}, err
	}
	resModel := models.ChildOrder{}

	if !isDry {
		if err := api.do(http.MethodPost, req, &resModel, url, nil, true); err != nil {
			return models.ChildOrder{}, err
		}
	}

	return resModel, nil
}

func (api *API) PostCancelChildOrder(req models.CancelChildOrderRequest, isDry bool) error {
	url, err := api.BaseUrl.PostCancelChildOrderUrl()
	if err != nil {
		return err
	}

	if !isDry {
		if err := api.do(http.MethodPost, req, nil, url, nil, true); err != nil {
			return err
		}
	}

	return nil
}

func (api *API) GetChildOrders() ([]models.ChildOrder, error) {
	url, err := api.BaseUrl.GetChildOrdersUrl()
	if err != nil {
		return []models.ChildOrder{}, err
	}
	var resModel []models.ChildOrder
	if err := api.do(http.MethodGet, nil, &resModel, url, nil, true); err != nil {
		return []models.ChildOrder{}, err
	}
	return resModel, nil
}

func StringTimeStamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
