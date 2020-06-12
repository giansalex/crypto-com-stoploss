package stoploss

import "math"

// Trailing stop-loss runner
type Trailing struct {
	lastStop float64
	exchange *Exchange

	Market     string
	Type       string
	StopFactor float64
}

// RunStop check stop loss apply
func (tlg *Trailing) RunStop() bool {
	marketPrice, _ := tlg.exchange.GetMarketPrice(tlg.Market)

	tlg.lastStop = tlg.refreshStop(tlg.lastStop, marketPrice)

	if marketPrice > tlg.lastStop {
		return false
	}

	quantity, _ := tlg.exchange.GetBalance(tlg.Market)
	tlg.exchange.Sell(tlg.Market, quantity)

	return true
}

func (tlg *Trailing) refreshStop(stop float64, price float64) float64 {
	return math.Max(stop, price*tlg.StopFactor)
}
