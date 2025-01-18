package services

import (
	"chatbot-backend/app/common/request"
	"chatbot-backend/global"
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type openAI struct {
}

var OpenAI = new(openAI)

func createOpenAIClient() *openai.Client {
	apiKey := global.App.Config.OpenAI.SecretKey

	// Create a new client
	return openai.NewClient(apiKey)
}

type MessageRes struct {
	content string `json:"content"`
}

func (OpenAI *openAI) SendMessage(params request.Message) (response MessageRes, err error) {
	client := createOpenAIClient()

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: params.Prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	} else {
		response = MessageRes{
			content: resp.Choices[0].Message.Content,
		}
	}

	return
}
