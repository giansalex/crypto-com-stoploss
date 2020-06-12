package crypto

// Order order parameters
type Order struct {
	Price  string `json:"price"`
	Side   string `json:"side"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Volume string `json:"volume"`
}

type orderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID int `json:"order_id"`
	} `json:"data"`
}
