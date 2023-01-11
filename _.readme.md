# Crypto.com Stop-Loss Bot [![Go Report Card](https://goreportcard.com/badge/github.com/giansalex/crypto-com-trailing-stop-loss)](https://goreportcard.com/report/github.com/giansalex/crypto-com-trailing-stop-loss)

![Crypto.com Exchange](https://raw.githubusercontent.com/giansalex/crypto-com-trailing-stop-loss/master/doc/crypto-com.png)

[Crypto.com](https://crypto.com/exchange) Trailing Stop-Loss Bot and optional Telegram notifications. 

> A trailing stop order sets the stop price at a fixed amount below the market price with an attached "trailing" amount. As the market price rises, the stop price rises by the trail amount, but if the stock price falls, the stop loss price doesn't change, and a market order is submitted when the stop price is hit.

## Installation
Follow this [guide](https://github.com/giansalex/crypto-com-trailing-stop-loss/wiki/Installation) or use [Docker image](https://hub.docker.com/r/giansalex/crypto-com-stoploss)

## Run

Require [API Keys](https://crypto.com/exchange/personal/api-management).    
Set Environment variables:
- `CRYPTO_APIKEY`
- `CRYPTO_SECRET`
- `TELEGRAM_TOKEN` (optional to notify)

Simple command to run bot stoploss
> Sell all BTC balance to market price when down 3%.
```sh
./crypto -pair=BTC/USDT -percent=3
```

For buy orders
> Buy 100 USDT when `BTC` up 0.5%
```sh
./crypto -type=BUY -pair=BTC/USDT -percent=0.5 -amount=100
```

For sell orders with static stoploss
> SELL 0.1 BTC when `BTC` down to 9400 USDT
```sh
./crypto -pair=BTC/USDT -price=9400 -amount=0.1
```


## Notifications

- Telegram.
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60 -telegram.chat=<user-id>
```
![Crypto bot Telegram](https://raw.githubusercontent.com/giansalex/crypto-com-trailing-stop-loss/master/doc/crypto-com-telegram-bot.png)

> For get user id, talk to the [userinfobot](https://t.me/userinfobot)


- Mailing.
```sh
./crypto -pair=BTC/USDT -percent=3 \
      -mail.host="smtp.example.com" \
      -mail.port=587 \
      -mail.user="user@example.com" \
      -mail.pass="xxxx" \
      -mail.from="user@example.com" \
      -mail.to="bob@gmail.com"
```

> You can notify both: telegram, mail.


List available parameters 
```sh
  -amount float
        (optional) amount to order (sell or buy) on stoploss
  -interval int
        interval in seconds to update price, example: 30 (30 sec.) (default 30)
  -mail.from string
        (optional) email sender
  -mail.host string
        (optional) SMTP Host
  -mail.pass string
        (optional) SMTP Password
  -mail.port int
        (optional) SMTP Port (default 587)
  -mail.to string
        (optional) email receptor
  -mail.user string
        (optional) SMTP User
  -pair string
        market pair, example: MCO/USDT
  -percent float
        percent (for trailing stop loss), example: 3.0 (3%)
  -price float
        price (for static stop loss), example: 9200.00 (BTC price)
  -stop-change
        Notify on stoploss change (default: false)
  -telegram.chat int
        (optional) telegram User ID for notify
  -type string
        order type: SELL or BUY (default "SELL")
```

## Signup
Are you new user?, [signup](https://platinum.crypto.com/r/chr2wsfs6g) and use this referral code: `chr2wsfs6g` for earn bonus coin.
