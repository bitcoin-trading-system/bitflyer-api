package api

import (
	"net/url"
	"path"
	"strings"
)

func (bu BaseUrl) GetBoardUrl(productCode string) (string, error) {
	qVal := url.Values{}
	if productCode != "" {
		qVal.Add("product_code", productCode)
	}
	return createUrl(string(bu), "/v1/board", qVal)
}

func (bu BaseUrl) GetTickerUrl(productCode string) (string, error) {
	qVal := url.Values{}
	if productCode != "" {
		qVal.Add("product_code", productCode)
	}

	return createUrl(string(bu), "/v1/ticker", qVal)
}

func (bu BaseUrl) GetExecutionsUrl(productCode, count, before, after string) (string, error) {
	qVal := url.Values{}
	if productCode != "" {
		qVal.Add("product_code", productCode)
	}
	if count != "" {
		qVal.Add("count", count)
	}
	if before != "" {
		qVal.Add("before", before)
	}
	if after != "" {
		qVal.Add("after", after)
	}

	return createUrl(string(bu), "/v1/executions", qVal)
}

func (bu BaseUrl) GetBoardStateUrl(productCode string) (string, error) {
	qVal := url.Values{}
	if productCode != "" {
		qVal.Add("product_code", productCode)
	}
	return createUrl(string(bu), "/v1/getboardstate", qVal)
}

func (bu BaseUrl) GetHealthUrl(productCode string) (string, error) {
	qVal := url.Values{}
	if productCode != "" {
		qVal.Add("product_code", productCode)
	}
	return createUrl(string(bu), "/v1/gethealth", qVal)
}

func (bu BaseUrl) GetBalanceUrl() (string, error) {
	return createUrl(string(bu), "/v1/me/getbalance", url.Values{})
}

func (bu BaseUrl) GetCollateralUrl() (string, error) {
	return createUrl(string(bu), "/v1/me/getcollateral", url.Values{})
}

func (bu BaseUrl) PostSendChildOrderUrl() (string, error) {
	return createUrl(string(bu), "/v1/me/sendchildorder", url.Values{})
}

func (bu BaseUrl) PostCancelChildOrderUrl() (string, error) {
	return createUrl(string(bu), "/v1/me/cancelchildorder", url.Values{})
}

func (bu BaseUrl) GetChildOrdersUrl() (string, error) {
	return createUrl(string(bu), "/v1/me/getchildorders", url.Values{})
}

func createUrl(baseUrl, p string, qVal url.Values, el ...string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	pEls := append([]string{p}, el...)
	u.Path = withSuffixSlash(path.Join(pEls...))

	u.RawQuery = qVal.Encode()

	return u.String(), nil
}

func withSuffixSlash(s string) string {
	if strings.HasSuffix(s, "/") {
		return s
	}
	return s + "/"
}

func extractPath(u string) (string, error) {
	uObj, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return uObj.Path, nil
}
