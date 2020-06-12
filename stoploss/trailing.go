package stoploss

import "math"

// Trailing stop-loss runner
type Trailing struct {
	exchange   *Exchange
	market     string
	baseCoin   string
	lastStop   float64
	stopFactor float64
}

// NewTrailing new trailing instance
func NewTrailing(exchange *Exchange, baseCoin string, countCoin string, factor float64) *Trailing {
	return &Trailing{
		exchange:   exchange,
		market:     baseCoin + countCoin,
		baseCoin:   baseCoin,
		stopFactor: factor,
	}
}

// RunStop check stop loss apply
func (tlg *Trailing) RunStop() bool {
	marketPrice, _ := tlg.exchange.GetMarketPrice(tlg.market)

	tlg.lastStop = tlg.refreshStop(tlg.lastStop, marketPrice)

	if marketPrice > tlg.lastStop {
		return false
	}

	quantity, _ := tlg.exchange.GetBalance(tlg.baseCoin)
	tlg.exchange.Sell(tlg.market, quantity)

	return true
}

func (tlg *Trailing) refreshStop(stop float64, price float64) float64 {
	return math.Max(stop, price*(1-tlg.stopFactor))
}
