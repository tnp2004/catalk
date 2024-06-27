package ai

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var clientInstance *genai.Client

func TextGemini(msg string) *genai.GenerateContentResponse {
	ctx := context.Background()

	client := newAiClient(ctx)
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		log.Printf("error generate content. Err: %s", err.Error())
	}

	return resp
}

func newAiClient(ctx context.Context) *genai.Client {
	if clientInstance != nil {
		return clientInstance
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Printf("error text gemini. Err: %s", err.Error())
	}
	clientInstance = client

	return clientInstance
}
