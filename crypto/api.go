package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

// NewAPI create new API
func NewAPI(apiKey string, apiSecret string) *API {

	return &API{apiKey, apiSecret, http.DefaultClient, "https://api.crypto.com/v2/"}
}

// GetPrice get current ticket price
func (api *API) GetPrice(ticket string) (*Price, error) {
	params := url.Values{}
	params.Add("instrument_name", ticket)

	resp, err := api.client.Get(api.BasePath + "public/get-ticker?" + params.Encode())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response priceResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return nil, errors.New(response.Message)
	}

	return &response.Result.Data, nil
}

// GetBalance account balance
func (api *API) GetBalance() ([]Balance, error) {
	params := url.Values{}
	params.Add("api_key", api.apiKey)
	params.Add("time", fmt.Sprintf("%d", api.unixTime()))
	params.Add("sign", api.createSign(params))

	resp, err := api.client.PostForm(api.BasePath+"account", params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response balanceResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != "0" {
		return nil, errors.New(response.Msg)
	}

	return response.Data.CoinList, nil
}

// CreateOrder create order
func (api *API) CreateOrder(order Order) (int, error) {
	params := url.Values{}
	params.Add("api_key", api.apiKey)

	if order.Type == "1" {
		params.Add("price", order.Price)
	}

	params.Add("side", order.Side)
	params.Add("symbol", order.Symbol)
	params.Add("time", fmt.Sprintf("%d", api.unixTime()))
	params.Add("type", order.Type)
	params.Add("volume", order.Volume)
	params.Add("sign", api.createSign(params))

	resp, err := api.client.PostForm(api.BasePath+"order", params)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var response orderResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != "0" {
		return 0, errors.New(response.Msg)
	}

	return response.Data.OrderID, nil
}

func (api *API) unixTime() int64 {
	return time.Now().UTC().Unix() * 1000
}

func (api *API) createSign(data url.Values) string {
	rawData := ""
	for key, values := range data {
		rawData += key + strings.Join(values, "")
	}
	rawData += api.apiSecret
	hash := sha256.Sum256([]byte(rawData))

	return hex.EncodeToString(hash[:])
}
