package app

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

func groupChat(client *openai.Client, bot *tgbotapi.BotAPI, userID int64, mess string) (int64, string, error) {

	if len(mess) > 4097 {
		return 0, "", errors.New("This model's maximum context length is 4097 tokens.")
	}

	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: mess,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		log.Printf("chat completion error: %v\n", err)
		return 0, "", err
	}

	msg := tgbotapi.NewMessage(-1001285932539, resp.Choices[0].Message.Content)
	//msg.ReplyToMessageID = messID

	if _, err = bot.Send(msg); err != nil {
		log.Printf("send message to user error: %v\n", err)
		return 0, "", err
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: resp.Choices[0].Message.Content,
	})

	return userID, resp.Choices[0].Message.Content, nil
}
