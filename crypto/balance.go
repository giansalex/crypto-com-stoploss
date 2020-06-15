package crypto

// Balance account balance
type Balance struct {
	Normal string `json:"normal"`
	Locked string `json:"locked"`
	Coin   string `json:"coin"`
}

type balanceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TotalAsset string    `json:"total_asset"`
		CoinList   []Balance `json:"coin_list"`
	} `json:"data"`
}
