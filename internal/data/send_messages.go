package data

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
)

func (data *UserData) Send(client *openai.Client, bot *tgbotapi.BotAPI, userID int64, messID int, mess string) error {

	if len(mess) > 4097 {

		msg := tgbotapi.NewMessage(userID, "this model's maximum context length is 4097 tokens...")
		msg.ReplyToMessageID = messID

		if _, err := bot.Send(msg); err != nil {
			log.Printf("send message to user error: %v\n", err)
			return err
		}

		return nil
	}

	messages := data.Get(userID)

	if len(mess) > 4097 {
		log.Printf("this model's maximum context length is 4097 tokens...")
		data.Clear()
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: mess,
	})

	var (
		resp openai.ChatCompletionResponse
		err  error
	)

	if resp, err = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	); err != nil {
		log.Printf("chat completion error: %v\n", err)
		return err
	}

	// -1001285932539
	msg := tgbotapi.NewMessage(userID, resp.Choices[0].Message.Content)
	msg.ReplyToMessageID = messID

	if _, err = bot.Send(msg); err != nil {
		log.Printf("send message to user error: %v\n", err)
		return err
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})

	data.Set(userID, messages)

	return nil
}
