package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

func (api API) do(method string, reqModel any, resModel any, url string, headers map[string]string, isPrivate bool) error {
	reqJson, err := MarshalJson(reqModel)
	if err != nil {
		return err
	}

	if isPrivate {
		authHeader, err := api.PrivateRequestHeader(StringTimeStamp(), method, url, reqJson)
		if err != nil {
			return err
		}
		headers = MergeStringMap(headers, authHeader)
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

func (api API) request(method, url string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
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
		fmt.Printf("status code: %d, body: %s", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}

	return body, nil
}

// https://lightning.bitflyer.com/docs#%E8%AA%8D%E8%A8%BC:~:text=%E4%BA%86%E6%89%BF%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84%E3%80%82-,%E8%AA%8D%E8%A8%BC,-Private%20API%20%E3%81%AE
func (api *API) PrivateRequestHeader(timeStamp, method, url string, body []byte) (map[string]string, error) {
	path, err := extractPath(url)
	if err != nil {
		return nil, err
	}

	rawMessage := json.RawMessage(body)

	text := timeStamp + method + path + string(rawMessage)

	h := hmac.New(sha256.New, []byte(api.ApiSecret))
	h.Write([]byte(text))
	sign := hex.EncodeToString(h.Sum(nil))

	return map[string]string{
		"ACCESS-KEY":       api.ApiKey,
		"ACCESS-TIMESTAMP": timeStamp,
		"ACCESS-SIGN":      sign,
	}, nil
}

func MergeStringMap(maps ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range maps {
		if m == nil {
			continue
		}
		for k, v := range m {
			if _, ok := res[k]; ok {
				log.Printf("key %s is already exists", k)
			}
			res[k] = v
		}
	}
	return res
}

func MarshalJson(v any) ([]byte, error) {
	if v == nil || v == "" {
		return []byte{}, nil
	}
	return json.Marshal(v)
}
