package crypto

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// API Crypto.com Api
type API struct {
	apiKey string
	client *http.Client

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
	params.Add("api_key", "")
	params.Add("time", "")
	params.Add("sign", "")

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
	params.Add("time", order.Time)
	params.Add("type", order.Type)
	params.Add("volume", order.Volume)
	params.Add("sign", order.Sign)
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
