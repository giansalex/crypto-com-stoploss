package crypto

// Order order parameters
type Order struct {
	Price    float64 `json:"price"`
	Side     string  `json:"side"`
	Symbol   string  `json:"symbol"`
	Type     string  `json:"type"`
	Quantity float64 `json:"quantity"`
}

type orderResponse struct {
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
	Result  struct {
		OrderID string `json:"order_id"`
	} `json:"result"`
}
