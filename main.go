package main

import (
	"github.com/mattbelanger321/dolya-bot/utilities"
)

func main() {
	utilities.CreateAndSendMessage(true, utilities.To, "DOLYA BOT ACTIVATED.", "POST")
}
