package cmd

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/giansalex/binance-stoploss/notify"
	cryptoCom "github.com/giansalex/crypto-com-stoploss/crypto"
	"github.com/giansalex/crypto-com-stoploss/stoploss"
)

var (
	typePtr          = flag.String("type", "SELL", "order type: SELL or BUY")
	pairPtr          = flag.String("pair", "", "market pair, example: MCO/USDT")
	percentPtr       = flag.Float64("percent", 0.00, "percent (for trailing stop loss), example: 3.0 (3%)")
	pricePtr         = flag.Float64("price", 0.00, "price (for static stop loss), example: 9200.00 (BTC price)")
	intervalPtr      = flag.Int("interval", 30, "interval in seconds to update price, example: 30 (30 sec.)")
	amountPtr        = flag.Float64("amount", 0, "(optional) amount to order (sell or buy) on stoploss")
	notifyChangesPtr = flag.Bool("stop-change", false, "Notify on stoploss change (default: false)")
	chatPtr          = flag.Int64("telegram.chat", 0, "(optional) telegram User ID for notify")
	mailHostPtr      = flag.String("mail.host", "", "(optional) SMTP Host")
	mailPortPtr      = flag.Int("mail.port", 587, "(optional) SMTP Port")
	mailUserPtr      = flag.String("mail.user", "", "(optional) SMTP User")
	mailPassPtr      = flag.String("mail.pass", "", "(optional) SMTP Password")
	mailFromPtr      = flag.String("mail.from", "", "(optional) email sender")
	mailToPtr        = flag.String("mail.to", "", "(optional) email receptor")
)

func Execute() {
	flag.Parse()
	apiKey := os.Getenv("CRYPTO_APIKEY")
	secret := os.Getenv("CRYPTO_SECRET")

	if apiKey == "" || secret == "" {
		log.Fatal("CRYPTO_APIKEY, CRYPTO_SECRET are required")
	}

	if pairPtr == nil || *pairPtr == "" {
		log.Fatal("pair market is required")
	}

	if (percentPtr == nil || *percentPtr <= 0) && (pricePtr == nil || *pricePtr <= 0) {
		log.Fatal("a price or percent parameter is required")
	}

	api := cryptoCom.NewAPI(apiKey, secret)
	notifier := buildNotify()
	config := &stoploss.Config{
		OrderType:        strings.ToUpper(*typePtr),
		Market:           *pairPtr,
		Price:            *pricePtr,
		Quantity:         *amountPtr,
		StopFactor:       *percentPtr / 100,
		NotifyStopChange: *notifyChangesPtr,
	}
	trailing := stoploss.NewTrailing(stoploss.NewExchange(api), notifier, &notify.LogNotify{}, config)

	for {
		if trailing.RunStop() {
			break
		}

		time.Sleep(time.Duration(*intervalPtr) * time.Second)
	}
}

func buildNotify() notify.SingleNotify {
	notifiers := []notify.SingleNotify{}

	tgToken := os.Getenv("TELEGRAM_TOKEN")
	if *chatPtr != 0 && tgToken != "" {
		notifiers = append(notifiers, notify.NewTelegramNotify(tgToken, *chatPtr))
	}

	if *mailHostPtr != "" && *mailUserPtr != "" && *mailPassPtr != "" && *mailFromPtr != "" && *mailToPtr != "" {
		subject := "Binance StopLoss Bot"
		notifiers = append(notifiers, notify.NewMailNotify(*mailHostPtr, *mailPortPtr, *mailUserPtr, *mailPassPtr, subject, *mailFromPtr, *mailToPtr))
	}

	return stoploss.NewNotify(notifiers)
}
