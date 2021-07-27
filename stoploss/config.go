package stoploss

// Config stop-loss configuration
type Config struct {
	OrderType        string
	Market           string
	Price            float64
	Quantity         float64
	StopFactor       float64
	NotifyStopChange bool
}
