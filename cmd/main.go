package main

import (
	"kickHisAss/bot"
	"log"
	"os"
)

func main() {

	cfg := bot.Configuration{}

	cfg.WebHook = os.Getenv("webHook")

	cfg.BotToken = os.Getenv("botToken")

	cfg.Host = "0.0.0.0"

	cfg.Port = os.Getenv("PORT")

	if err := bot.Run(&cfg); err != nil {
		log.Printf("Error bot run: %s\n", err)
	}
}
