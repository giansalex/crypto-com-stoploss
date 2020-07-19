package stoploss

// Exchange wrapper to connect to any exchange
type Exchange interface {
	GetBalance(coin string) (float64, error)
	GetMarketPrice(market string) (float64, error)
	Sell(market string, quantity float64) (string, error)
	Buy(market string, quantity float64) (string, error)
}