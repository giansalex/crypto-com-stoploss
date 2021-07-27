package stoploss

import (
	"testing"

	"github.com/matryer/is"
)

type MockExchange struct {
	balance float64
	price   float64
}

func (exchange *MockExchange) GetBalance(coin string) (float64, error) {
	return exchange.balance, nil
}

func (exchange *MockExchange) GetMarketPrice(market string) (float64, error) {

	return exchange.price, nil
}

func (exchange *MockExchange) Sell(market string, quantity float64) (string, error) {
	return "00001", nil
}

func (exchange *MockExchange) Buy(market string, quantity float64) (string, error) {
	return "00001", nil
}

func TestSell(t *testing.T) {

	notify := NewNotify("", 0)
	exchange := &MockExchange{0.01, 0}
	config := &Config{
		OrderType:  "SELL",
		Market:     "BTC/USDT",
		StopFactor: 0,
		Quantity:   0,
		Price:      9000,
	}
	trailing := NewTrailing(exchange, notify, config)

	is := is.New(t)

	exchange.price = 9200
	sold := trailing.RunStop()
	is.True(!sold)

	exchange.price = 9100
	sold = trailing.RunStop()
	is.True(!sold)

	exchange.price = 8900
	sold = trailing.RunStop()
	is.True(sold)
}

func TestBuy(t *testing.T) {

	notify := NewNotify("", 0)
	exchange := &MockExchange{}
	config := &Config{
		OrderType:  "BUY",
		Market:     "BTC/USDT",
		StopFactor: 0,
		Quantity:   200,
		Price:      9000,
	}
	trailing := NewTrailing(exchange, notify, config)

	is := is.New(t)

	exchange.price = 8700
	bought := trailing.RunStop()
	is.True(!bought)

	exchange.price = 8900
	bought = trailing.RunStop()
	is.True(!bought)

	exchange.price = 9000
	bought = trailing.RunStop()
	is.True(bought)
}
