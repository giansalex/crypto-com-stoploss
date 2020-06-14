# Crypto.com Stop-Loss Bot

[Crypto.com](https://crypto.com/) Trailing Stop-Loss Bot. 

## Run

Simple comman to run trailing stoploss.
Require environment variables: `CRYPTO_APIKEY`, `CRYPTO_SECRET`.
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60
```

Use telegram for notifications.
Require additional environment variables: `TELEGRAM_TOKEN`
```sh
./crypto -pair=BTC/USDT -percent=3 -interval=60 -telegram.channel=XXXXXXXX
```