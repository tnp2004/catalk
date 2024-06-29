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

func TextToGemini(req *GeminiRequest) (*GeminiResponse, error) {
	contentResp, err := sendMsgToGemini(req)

	if err != nil {
		return nil, err
	}

	newMessageHistory := append(req.MessageHistory, contentResp...)
	resp := &GeminiResponse{
		NewMessageHistory: newMessageHistory,
		ResponseMessage:   contentResp[len(contentResp)-1].Message,
	}

	return resp, nil
}

func sendMsgToGemini(req *GeminiRequest) ([]*MessageInfo, error) {
	ctx := context.Background()

	client, err := newAiClient(ctx)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	cs := model.StartChat()
	msgHistoryContents := convertMessageHistoryToParts(req.MessageHistory)
	cs.History = append(cs.History, msgHistoryContents...)

	resp, err := cs.SendMessage(ctx, genai.Text(req.NewUserMessage))
	if err != nil {
		log.Printf("error generate content. Err: %s", err.Error())
		return nil, fmt.Errorf("generate content error: %s", err.Error())
	}
	r := make([]*MessageInfo, 0)
	r = append(r, &MessageInfo{
		Message: req.NewUserMessage,
		Role:    "user",
	})
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			r = append(r, &MessageInfo{
				Message: string(txt),
				Role:    resp.Candidates[0].Content.Role,
			})
		}
	}

	return r, nil
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
