package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

// Config struct
type Config struct {
	BotToken string `required:"true" split_words:"true"`
	//AiToken  string `required:"true" split_words:"true"`
	WebHook string `required:"true" split_words:"true"`
	Channel int64  `required:"true"`
}

// Run start bot
func Run() error {

	// get env
	cfg := Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}

	// Start botAPI with token
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	wh, _ := tgbotapi.NewWebhook(cfg.WebHook + cfg.BotToken)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook(cfg.WebHook + bot.Token)
	go http.ListenAndServe("/", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

	return nil
}
