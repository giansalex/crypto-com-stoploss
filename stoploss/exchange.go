package stoploss

import (
	"strconv"

	cryptoCom "github.com/giansalex/crypto-com-trailing-stop-loss/crypto"
)

// Exchange wrapper to connect to crypto.com API
type Exchange struct {
	api *cryptoCom.API
}

// GetBalance get balance for coin
func (exchange *Exchange) GetBalance(coin string) (string, error) {
	balances, err := exchange.api.GetBalance()
	if err != nil {
		return "0", err
	}

	for _, balance := range balances {
		if balance.Coin == coin {
			return balance.Normal, nil
		}
	}

	return "0", nil
}

// GetMarketPrice get last price for market pair
func (exchange *Exchange) GetMarketPrice(market string) (float64, error) {
	price, err := exchange.api.GetPrice(market)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(price.Last, 64)
}

// Sell create a sell order to market price
func (exchange *Exchange) Sell(market string, quantity string) (int, error) {
	order := cryptoCom.Order{
		Side:   "SELL",
		Symbol: market,
		Type:   "2", // type=2: Market Price
		Volume: quantity,
	}

	return exchange.api.CreateOrder(order)
}
