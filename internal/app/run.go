package app

import (
	"log"
	"regexp"
	"time"

	"github.com/MarlikAlmighty/kickHisAss/internal/data"

	"github.com/MarlikAlmighty/kickHisAss/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
)

// Run start bot
func Run() error {

	log.Println("[LOG]: Run Vasya, run...")

	// get env
	cfg := config.New()
	if err := cfg.GetEnv(); err != nil {
		return err
	}

	// new where we saved data
	users := data.New()

	// client ai
	clientAI := openai.NewClient(cfg.AiToken)

	// Start botAPI with token
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	// for limit request
	limiter := time.Tick(time.Minute / 3)

	// regexp for clear text
	reg := regexp.MustCompile(`^@ai\s+`)

	// clear data every 10 minutes
	go CleanUserData(users)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {

			// only my channel from config
			if update.Message.Chat.ID != cfg.Channel {
				continue
			}

			mess := update.Message.Text

			if matched := reg.MatchString(mess); matched {

				<-limiter

				userID := update.Message.Chat.ID
				messID := update.Message.MessageID

				mess = clearText(mess, reg)

				if err = users.Send(clientAI, bot, userID, messID, mess); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
