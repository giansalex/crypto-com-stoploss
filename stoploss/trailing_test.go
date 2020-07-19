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
	trailing := NewTrailing(exchange, notify, "SELL", "BTC/USDT", 0, 0, 9000)

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
	trailing := NewTrailing(exchange, notify, "BUY", "BTC/USDT", 0, 200, 9000)

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
