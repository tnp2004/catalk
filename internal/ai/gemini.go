package ai

import (
	"catalk/instructions"
	"catalk/utils"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var clientInstance *genai.Client

func TextToGemini(req *GeminiRequest, breed string, apiKey string) (*GeminiResponse, error) {
	breed, ok := instructions.CatBreedsMap[breed]

	if !ok {
		log.Printf("error %s breed isn't match", breed)
		return nil, fmt.Errorf("%s breed isn't match", breed)
	}

	instructions, err := utils.ReadInstructions("instructions/json/cat.json")
	if err != nil {
		return nil, err
	}

	breedIns, ok := instructions.BreedsInstruction[breed]
	if !ok {
		log.Printf("error instruction of %s breed not found", breed)
		return nil, fmt.Errorf("%s instruction not found", breed)
	}

	contentResp, err := sendMsgToGemini(req, instructions.MainInstruction, breedIns, apiKey)
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

func sendMsgToGemini(req *GeminiRequest, mainInstruction string, subInstruction string, apiKey string) ([]*MessageInfo, error) {
	ctx := context.Background()

	client, err := newAiClient(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(mainInstruction), genai.Text(subInstruction)},
	}
	cs := model.StartChat()
	msgHistoryContents := convertMessageHistoryToParts(req.MessageHistory)
	cs.History = append(cs.History, msgHistoryContents...)

	contentResp, err := cs.SendMessage(ctx, genai.Text(req.NewUserMessage))
	if err != nil {
		log.Printf("error generate content. Err: %s", err.Error())
		return nil, fmt.Errorf("generate content error: %s", err.Error())
	}
	resp := make([]*MessageInfo, 0)
	resp = append(resp, &MessageInfo{
		Message: req.NewUserMessage,
		Role:    "user",
	})
	for _, part := range contentResp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			resp = append(resp, &MessageInfo{
				Message: string(txt),
				Role:    contentResp.Candidates[0].Content.Role,
			})
		}
	}

	return resp, nil
}

func newAiClient(ctx context.Context, apiKey string) (*genai.Client, error) {
	if clientInstance != nil {
		return clientInstance, nil
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("error new gemini client. Err: %s", err.Error())
		return nil, fmt.Errorf("new gemini client error: %s", err.Error())
	}
	clientInstance = client

	return clientInstance, nil
}
