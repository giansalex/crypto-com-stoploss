package crypto

// Balance account balance
type Balance struct {
	Balance   float64 `json:"balance"`
	Available float64 `json:"available"`
	Order     float64 `json:"order"`
	Stake     float64 `json:"stake"`
	Currency  string  `json:"currency"`
}

type balanceResponse struct {
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
	Result  struct {
		Accounts []Balance `json:"accounts"`
	} `json:"result"`
}
