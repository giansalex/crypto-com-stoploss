package stoploss

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// Trailing stop-loss runner
type Trailing struct {
	exchange   *Exchange
	notify     *Notify
	market     string
	baseCoin   string
	lastStop   float64
	quantity   string
	stopFactor float64
}

// NewTrailing new trailing instance
func NewTrailing(exchange *Exchange, notify *Notify, baseCoin string, countCoin string, factor float64, quantity string) *Trailing {
	return &Trailing{
		exchange:   exchange,
		notify:     notify,
		market:     baseCoin + countCoin,
		baseCoin:   baseCoin,
		quantity:   quantity,
		stopFactor: factor,
	}
}

// RunStop check stop loss apply
func (tlg *Trailing) RunStop() bool {
	marketPrice, _ := tlg.exchange.GetMarketPrice(tlg.market)

	stop := tlg.refreshStop(tlg.lastStop, marketPrice)

	if marketPrice > stop {
		tlg.notifyStopLossChange(tlg.lastStop, stop, marketPrice)

		tlg.lastStop = stop
		return false
	}

	quantity := tlg.quantity
	if quantity == "" {
		quantity, _ = tlg.exchange.GetBalance(tlg.baseCoin)
	}

	tlg.exchange.Sell(tlg.market, quantity)
	tlg.notify.Send(fmt.Sprintf("Sell: %s %s - Market Price: %.6f", quantity, strings.ToUpper(tlg.baseCoin), marketPrice))

	return true
}

func (tlg *Trailing) refreshStop(stop float64, price float64) float64 {
	return math.Max(stop, price*(1-tlg.stopFactor))
}

func (tlg *Trailing) notifyStopLossChange(prev float64, next float64, price float64) {
	result := big.NewFloat(prev).Cmp(big.NewFloat(next))

	if result == 0 {
		return
	}

	tlg.notify.Send(fmt.Sprintf("Stop-loss %s: %.6f - Market Price: %.6f", strings.ToUpper(tlg.market), next, price))
}
