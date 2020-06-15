package crypto

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
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
func (api *API) GetBalance(coin string) ([]Balance, error) {
	method := "private/get-account-summary"

	request := make(map[string]interface{})
	request["id"] = api.createID()
	request["method"] = method
	request["params"] = map[string]interface{}{
		"currency": coin,
	}
	request["api_key"] = api.apiKey
	request["nonce"] = api.unixTime()

	api.sign(request)

	payload, _ := json.Marshal(request)
	resp, err := api.client.Post(api.BasePath+method, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response balanceResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Result.Accounts, nil
}

// CreateOrder create order
func (api *API) CreateOrder(order Order) (string, error) {
	method := "private/create-order"

	params := map[string]interface{}{
		"instrument_name": order.Symbol,
		"side":            order.Side,
		"type":            order.Type,
	}

	if order.Type == "LIMIT" {
		params["price"] = order.Price
	}

	if order.Type == "MARKET" && order.Side == "BUY" {
		params["notional"] = order.Quantity
	} else {
		params["quantity"] = order.Quantity
	}

	request := make(map[string]interface{})
	request["id"] = api.createID()
	request["method"] = method
	request["params"] = params
	request["api_key"] = api.apiKey
	request["nonce"] = api.unixTime()

	api.sign(request)

	payload, _ := json.Marshal(request)
	resp, err := api.client.Post(api.BasePath+method, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var response orderResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Code != 0 {
		return "", errors.New(response.Message)
	}

	return response.Result.OrderID, nil
}

func (api *API) createID() int64 {
	return time.Now().UTC().Unix()
}

func (api *API) unixTime() int64 {
	return time.Now().UTC().UnixNano() / 1e6
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

func (api *API) sign(request map[string]interface{}) {
	params := request["params"].(map[string]interface{})
	paramString := ""
	for _, keySort := range api.getSortKeys(params) {
		paramString += keySort + fmt.Sprintf("%v", params[keySort])
	}
	sigPayload := fmt.Sprintf("%v%v%s%s%v", request["method"], request["id"], api.apiKey, paramString, request["nonce"])

	key := []byte(api.apiSecret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(sigPayload))

	request["sig"] = hex.EncodeToString(mac.Sum(nil))
}

func (api *API) getSortKeys(params map[string]interface{}) []string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
