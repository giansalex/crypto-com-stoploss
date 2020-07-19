package stoploss

import (
	"strings"

	cryptoCom "github.com/giansalex/crypto-com-trailing-stop-loss/crypto"
)

// CryptoExchange wrapper to connect to crypto.com API
type CryptoExchange struct {
	api *cryptoCom.API
}

// NewExchange create Exchange instance
func NewExchange(api *cryptoCom.API) *CryptoExchange {
	return &CryptoExchange{api}
}

// GetBalance get balance for coin
func (exchange *CryptoExchange) GetBalance(coin string) (float64, error) {
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
func (exchange *CryptoExchange) GetMarketPrice(market string) (float64, error) {
	price, err := exchange.api.GetPrice(market)
	if err != nil {
		return 0, err
	}

	return price.Last, nil
}

// Sell create a sell order to market price
func (exchange *CryptoExchange) Sell(market string, quantity float64) (string, error) {
	order := cryptoCom.Order{
		Side:     "SELL",
		Symbol:   market,
		Type:     "MARKET",
		Quantity: quantity,
	}

	return exchange.api.CreateOrder(order)
}

// Buy create a buy order to market price
func (exchange *CryptoExchange) Buy(market string, quantity float64) (string, error) {
	order := cryptoCom.Order{
		Side:     "BUY",
		Symbol:   market,
		Type:     "MARKET",
		Quantity: quantity,
	}

	return exchange.api.CreateOrder(order)
}
