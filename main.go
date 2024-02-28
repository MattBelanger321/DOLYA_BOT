package main

import (
	"github.com/mattbelanger321/dolya-bot/utilities"
)

func main() {

	msgData := utilities.CreateMessage(true, "DOLYA BOT ACTIVATED")
	req := utilities.CreateRequest("POST", msgData)
	utilities.SendMessage(req)
}
