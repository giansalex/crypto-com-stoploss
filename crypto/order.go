package crypto

// Order order parameters
type Order struct {
	Side   string  `json:"side"`
	Symbol string  `json:"symbol"`
	Type   int     `json:"type"`
	Volume float64 `json:"volume"`
}

type orderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID int `json:"order_id"`
	} `json:"data"`
}
