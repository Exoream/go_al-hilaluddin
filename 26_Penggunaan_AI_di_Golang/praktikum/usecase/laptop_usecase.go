package usecase

import (
	"context"
	"laptop/entity"

	openai "github.com/sashabaranov/go-openai"
)

type LaptopUsecase struct {
	userData entity.LaptopUsecase
	openAIKey string
}

func NewUserService(userData entity.LaptopUsecase, openAIKey string) *LaptopUsecase {
	return &LaptopUsecase{
		userData:   userData,
		openAIKey: openAIKey,
	}
}

func (u *LaptopUsecase) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (u *LaptopUsecase) RecommendLaptop(params entity.LaptopParams) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(u.openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Selamat datang! Saya adalah sistem rekomendasi laptop.",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: params.UserInput,
		},
	}

	resp, err := u.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

