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
	orderType  string
	market     string
	baseCoin   string
	countCoin  string
	lastStop   float64
	quantity   float64
	stopFactor float64
}

// NewTrailing new trailing instance
func NewTrailing(exchange *Exchange, notify *Notify, orderType string, market string, factor float64, quantity float64) *Trailing {
	pair := strings.Split(strings.ToUpper(market), "/")

	tlg := &Trailing{
		exchange:   exchange,
		notify:     notify,
		orderType:  strings.ToUpper(orderType),
		market:     pair[0] + "_" + pair[1],
		baseCoin:   pair[0],
		countCoin:  pair[1],
		quantity:   quantity,
		stopFactor: factor,
	}

	if tlg.orderType == "BUY" {
		tlg.lastStop = math.MaxFloat64
	}

	return tlg
}

// RunStop check stop loss apply
func (tlg *Trailing) RunStop() bool {
	if tlg.orderType == "BUY" {
		return tlg.runBuy()
	}

	return tlg.runSell()
}

func (tlg *Trailing) runSell() bool {
	marketPrice, err := tlg.exchange.GetMarketPrice(tlg.market)
	if err != nil {
		tlg.notify.Send("Cannot get market price, error:" + err.Error())
		return true
	}

	stop := tlg.refreshStop(tlg.lastStop, marketPrice)

	if marketPrice > stop {
		tlg.notifyStopLossChange(tlg.lastStop, stop, marketPrice)

		tlg.lastStop = stop
		return false
	}

	quantity := tlg.quantity
	if quantity == 0 {
		quantity, err = tlg.exchange.GetBalance(tlg.baseCoin)
		if err != nil {
			tlg.notify.Send("Cannot get balance, error:" + err.Error())
			return true
		}
	}

	order, err := tlg.exchange.Sell(tlg.market, quantity)
	if err != nil {
		tlg.notify.Send("Cannot create sell order, error:" + err.Error())
	} else {
		tlg.notify.Send(fmt.Sprintf("Sell: %.4f %s - Market Price: %.6f - Order ID: %s", quantity, strings.ToUpper(tlg.baseCoin), marketPrice, order))
	}

	return true
}

func (tlg *Trailing) runBuy() bool {
	marketPrice, err := tlg.exchange.GetMarketPrice(tlg.market)
	if err != nil {
		tlg.notify.Send("Cannot get market price, error:" + err.Error())
		return true
	}

	stop := tlg.refreshBuyStop(tlg.lastStop, marketPrice)

	if stop > marketPrice {
		tlg.notifyStopLossChange(tlg.lastStop, stop, marketPrice)

		tlg.lastStop = stop
		return false
	}

	quantity := tlg.quantity
	if quantity == 0 {
		quantity, err = tlg.exchange.GetBalance(tlg.countCoin)
		if err != nil {
			tlg.notify.Send("Cannot get balance, error:" + err.Error())
			return true
		}
	}

	order, err := tlg.exchange.Buy(tlg.market, quantity)
	if err != nil {
		tlg.notify.Send("Cannot create buy order, error:" + err.Error())
	} else {
		tlg.notify.Send(fmt.Sprintf("Buy: %.4f %s - Market Price (%s): %.6f - Order ID: %s", quantity, tlg.countCoin, tlg.baseCoin, marketPrice, order))
	}

	return true
}

func (tlg *Trailing) refreshBuyStop(stop float64, price float64) float64 {
	return math.Min(stop, price*(1+tlg.stopFactor))
}

func (tlg *Trailing) refreshStop(stop float64, price float64) float64 {
	return math.Max(stop, price*(1-tlg.stopFactor))
}

func (tlg *Trailing) notifyStopLossChange(prev float64, next float64, price float64) {
	result := big.NewFloat(prev).Cmp(big.NewFloat(next))

	if result == 0 {
		return
	}

	tlg.notify.Send(fmt.Sprintf("Stop-loss %s (%s): %.6f - Market Price: %.6f", tlg.market, tlg.orderType, next, price))
}
