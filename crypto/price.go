package crypto

// Price price of crypto.com ticket
type Price struct {
	High int     `json:"high"`
	Vol  float64 `json:"vol"`
	Last float64 `json:"last"`
	Low  float64 `json:"low"`
	Buy  string  `json:"buy"`
	Sell string  `json:"sell"`
	Rose float64 `json:"rose"`
	Time int64   `json:"time"`
}

type priceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Price  `json:"data"`
}
