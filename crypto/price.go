package crypto

// Price price of crypto.com ticket
type Price struct {
	High   float64 `json:"h"`
	Vol    float64 `json:"v"`
	Last   float64 `json:"a"`
	Low    float64 `json:"l"`
	Buy    float64 `json:"b"`
	Sell   float64 `json:"k"`
	Change float64 `json:"c"`
	Time   int64   `json:"t"`
}

type priceResponse struct {
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
	Result  struct {
		Data Price `json:"data"`
	} `json:"result"`
}
