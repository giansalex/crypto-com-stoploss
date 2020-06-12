package crypto

// Order order parameters
type Order struct {
	APIKey string  `json:"api_key"`
	Side   string  `json:"side"`
	Symbol string  `json:"symbol"`
	Time   int     `json:"time"`
	Type   int     `json:"type"`
	Volume float64 `json:"volume"`
	Sign   string  `json:"sign"`
}

type orderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID int `json:"order_id"`
	} `json:"data"`
}
