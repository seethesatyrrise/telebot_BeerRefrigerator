package main

import (
	"fmt"
	"telebot_BeerRefrigerator/internal/bot"
)

func main() {
	bot, err := bot.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()
}
