package crypto

// Price price of crypto.com ticket
type Price struct {
	High int    `json:"high"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Low  string `json:"low"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Rose string `json:"rose"`
	Time int64  `json:"time"`
}

type priceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Price  `json:"data"`
}
