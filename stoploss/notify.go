package stoploss

import "fmt"

// Notify notify stoploss
type Notify struct {
	tlgToken string
	chanelID int
}

// Send send message
func (notify *Notify) Send(message string) {
	fmt.Println(message)
}
