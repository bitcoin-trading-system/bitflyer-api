package api

import (
	"net/url"
	"path"
	"strings"
)

func (bu BaseUrl) GetBoardUrl(productCode string) (string, error) {
	return createUrl(string(bu), "/v1/board", url.Values{"product_code": {productCode}})
}

func (bu BaseUrl) GetTickerUrl(productCode string) (string, error) {
	return createUrl(string(bu), "/v1/ticker", url.Values{"product_code": {productCode}})
}

func (bu BaseUrl) GetExecutionsUrl(productCode string, count, before, after int) (string, error) {
	return createUrl(string(bu), "/v1/executions", url.Values{
		"product_code": {productCode},
		"count":        {string(rune(count))},
		"before":       {string(rune(before))},
		"after":        {string(rune(after))},
	})
}

func (bu BaseUrl) GetBoardStateUrl(productCode string) (string, error) {
	return createUrl(string(bu), "/v1/getboardstate", url.Values{"product_code": {productCode}})
}

func (bu BaseUrl) GetHealthUrl(productCode string) (string, error) {
	return createUrl(string(bu), "/v1/gethealth", url.Values{"product_code": {productCode}})
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
