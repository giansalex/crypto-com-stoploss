package crypto

// Balance account balance
type Balance struct {
	Normal      float64 `json:"normal"`
	Locked      float64 `json:"locked"`
	BtcValuatin float64 `json:"btcValuatin"`
	Coin        string  `json:"coin"`
}

type balanceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TotalAsset float64   `json:"total_asset"`
		CoinList   []Balance `json:"coin_list"`
	} `json:"data"`
}
