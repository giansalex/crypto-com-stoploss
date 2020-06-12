package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// API Crypto.com Api
type API struct {
	apiKey    string
	apiSecret string
	client    *http.Client

	BasePath string
}

// GetPrice get current ticket price
func (api *API) GetPrice(ticket string) (*Price, error) {
	params := url.Values{}
	params.Add("symbol", ticket)

	resp, err := api.client.Get(api.BasePath + "/v1/ticker?" + params.Encode())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response priceResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return nil, errors.New(response.Msg)
	}

	return &response.Data, nil
}

// GetBalance account balance
func (api *API) GetBalance() ([]Balance, error) {
	params := url.Values{}
	params.Add("api_key", api.apiKey)
	params.Add("time", api.unixTime())
	params.Add("sign", api.createSign(params))

	resp, err := api.client.PostForm(api.BasePath+"/v1/account", params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response balanceResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return nil, errors.New(response.Msg)
	}

	return response.CoinList, nil
}

// Sell create sell order
func (api *API) Sell(order Order) (int, error) {
	params := url.Values{}
	params.Add("api_key", order.APIKey)
	params.Add("side", order.Side)
	params.Add("symbol", order.Symbol)
	params.Add("time", api.unixTime())
	params.Add("type", order.Type)
	params.Add("volume", order.Volume)
	params.Add("sign", api.createSign(params))

	resp, err := api.client.PostForm(api.BasePath+"/v1/order", params)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var response orderResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return 0, errors.New(response.Msg)
	}

	return response.Data.OrderID, nil
}

func (api *API) unixTime() int64 {
	return time.Now().UTC().Unix()
}

func (api *API) createSign(data url.Values) string {
	rawData := ""
	for key, values := range data {
		rawData += key + strings.Join(values, "")
	}
	bytesToSign = []byte(rawData + api.apiSecret)

	return hex.EncodeToString(sha256.Sum256(bytesToSign))
}
