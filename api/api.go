package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type BaseUrl string

type API struct {
	BaseUrl   BaseUrl
	ApiKey    string
	ApiSecret string
}

func (api API) do(method string, reqModel any, resModel any, url string, headers map[string]string) error {
	reqJson, err := json.Marshal(reqModel)
	if err != nil {
		return err
	}

	res, err := api.request(method, url, reqJson, headers)
	if err != nil {
		return err
	}

	resJson, err := api.readResponse(res)
	if err != nil {
		return err
	}

	return json.Unmarshal(resJson, resModel)
}

func (api API) request(method, url string, body any, headers map[string]string) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return http.DefaultClient.Do(req)
}

func (api API) readResponse(resp *http.Response) ([]byte, error) {
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return body, nil
}
