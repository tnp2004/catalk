package ai

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var clientInstance *genai.Client

func TextToGemini(body *GeminiTextBody) (*GeminiResponse, error) {
	contentResp, err := sendMsgToGemini(body.Message)
	if err != nil {
		return nil, err
	}

	resp := &GeminiResponse{Message: contentResp.Candidates[0].Content.Parts[0]}

	return resp, nil
}

func sendMsgToGemini(msg string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()

	client, err := newAiClient(ctx)
	if err != nil {
		return nil, err
	}
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		log.Printf("error generate content. Err: %s", err.Error())
		return nil, fmt.Errorf("generate content error: %s", err.Error())
	}

	return resp, nil
}

func newAiClient(ctx context.Context) (*genai.Client, error) {
	if clientInstance != nil {
		return clientInstance, nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
	if err != nil {
		log.Printf("error new gemini client. Err: %s", err.Error())
		return nil, fmt.Errorf("new gemini client error: %s", err.Error())
	}
	clientInstance = client

	return clientInstance, nil
}
