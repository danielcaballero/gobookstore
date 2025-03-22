package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gobookstore/models"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

func GenerateRandomBook(ctx context.Context) (*models.Book, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	prompt := `Generate a random book in JSON format with fields: title, author, description, genre, year.`

	var resp openai.ChatCompletionResponse
	var err error
	maxRetries := 5

	for retries := 0; retries < maxRetries; retries++ {
		resp, err = client.CreateChatCompletion(
			ctx,
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{Role: openai.ChatMessageRoleSystem, Content: "You are a helpful assistant."},
					{Role: openai.ChatMessageRoleUser, Content: prompt},
				},
			},
		)
		if err == nil {
			break
		}
		if isRateLimitError(err) {
			time.Sleep(time.Second * time.Duration(2^retries)) // Exponential backoff
			continue
		}
		return nil, fmt.Errorf("API error: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed after retries: %w", err)
	}

	var book models.Book
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &book); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &book, nil
}

func isRateLimitError(err error) bool {
	return strings.Contains(err.Error(), "429")
}
