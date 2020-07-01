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

For sell orders with static stoploss (example: SELL 0.1 BTC when `BTC` down to 9400 USDT)

```sh
./crypto -pair=BTC/USDT -price=9400 -amount=0.1
```

Use telegram for notifications - 
Require additional environment variables: `TELEGRAM_TOKEN`
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60 -telegram.chat=<user-id>
```
> For get user id, talk o the [userinfobot](https://t.me/userinfobot)

![Crypto bot Telegram](https://raw.githubusercontent.com/giansalex/crypto-com-trailing-stop-loss/master/doc/telegram-cryptobot.png)

List available parameters 
```sh
  -type string
        order type: SELL or BUY (default: SELL)
  -pair string
        market pair, example: MCO/USDT
  -interval int
        interval in seconds to update price, example: 30 (30 sec.) (default 30)
  -price float
        price (for static stoploss)
  -percent float
        percent (for trailing stoploss), example: 3.0 (3%)
  -amount float
        (optional) amount to order (sell or buy) on stoploss, default all balance
  -telegram.chat int
        (optional) telegram User ID for notify
```
