# Crypto.com Stop-Loss Bot [![Go Report Card](https://goreportcard.com/badge/github.com/giansalex/crypto-com-trailing-stop-loss)](https://goreportcard.com/report/github.com/giansalex/crypto-com-trailing-stop-loss)

![Crypto.com Exchange](https://raw.githubusercontent.com/giansalex/crypto-com-trailing-stop-loss/master/doc/crypto-com.png)

[Crypto.com](https://crypto.com/exchange) Trailing Stop-Loss Bot and optional Telegram notifications. 

> A trailing stop order sets the stop price at a fixed amount below the market price with an attached "trailing" amount. As the market price rises, the stop price rises by the trail amount, but if the stock price falls, the stop loss price doesn't change, and a market order is submitted when the stop price is hit.

## Build
Build executable `go build -ldflags "-s -w" -o crypto`.

## Run
> You can use [docker image](https://hub.docker.com/r/giansalex/crypto-com-stoploss).

First create [API Keys](https://crypto.com/exchange/personal/api-management). 

Simple command to run bot stoploss -
Require environment variables: `CRYPTO_APIKEY`, `CRYPTO_SECRET`.
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60
```

For buy orders (example: Buy 100 USDT when `BTC` up 0.5%)

```sh
./crypto -type=BUY -pair=BTC/USDT -percent=0.5 -amount=100
```

Use telegram for notifications - 
Require additional environment variables: `TELEGRAM_TOKEN`
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60 -telegram.chat=<chat-id>
```

![Crypto bot Telegram](https://raw.githubusercontent.com/giansalex/crypto-com-trailing-stop-loss/master/doc/telegram-cryptobot.png)

List available parameters 
```sh
  -pair string
        market pair, example: MCO/USDT
  -interval int
        interval in seconds to update price, example: 30 (30 sec.) (default 30)
  -percent float
        stop loss percent, example: 3.0 (3%)
  -amount float
        (optional) amount to order (sell or buy) on stoploss, default all balance
  -telegram.chat int
        (optional) telegram Chat ID for notify
```
