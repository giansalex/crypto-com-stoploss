package stoploss

import (
	"strings"

	cryptoCom "github.com/giansalex/crypto-com-trailing-stop-loss/crypto"
)

// Exchange wrapper to connect to crypto.com API
type Exchange struct {
	api *cryptoCom.API
}

// NewExchange create Exchange instance
func NewExchange(api *cryptoCom.API) *Exchange {
	return &Exchange{api}
}

// GetBalance get balance for coin
func (exchange *Exchange) GetBalance(coin string) (float64, error) {
	coin = strings.ToUpper(coin)
	balances, err := exchange.api.GetBalance(coin)
	if err != nil {
		return 0, err
	}

	for _, balance := range balances {
		if strings.ToUpper(balance.Currency) == coin {
			return balance.Available, nil
		}
	}

	return 0, nil
}

// GetMarketPrice get last price for market pair
func (exchange *Exchange) GetMarketPrice(market string) (float64, error) {
	price, err := exchange.api.GetPrice(market)
	if err != nil {
		return 0, err
	}

	return price.Last, nil
}

// Sell create a sell order to market price
func (exchange *Exchange) Sell(market string, quantity float64) (string, error) {
	order := cryptoCom.Order{
		Side:     "SELL",
		Symbol:   market,
		Type:     "MARKET",
		Quantity: quantity,
	}

	return exchange.api.CreateOrder(order)
}

// Buy create a buy order to market price
func (exchange *Exchange) Buy(market string, quantity float64) (string, error) {
	order := cryptoCom.Order{
		Side:     "BUY",
		Symbol:   market,
		Type:     "MARKET",
		Quantity: quantity,
	}

	return exchange.api.CreateOrder(order)
}
