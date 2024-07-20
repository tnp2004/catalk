package ai

import (
	"catalk/config"
	"catalk/instructions"
	"catalk/utils"
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ai struct {
	googleConfig *config.Google
}

type AiService interface {
	TextToGemini(req *GeminiRequest, breed string) (*GeminiResponse, error)
}

func NewAi(config *config.Google) AiService {
	return &ai{googleConfig: config}
}

var (
	clientInstance *genai.Client
	once           sync.Once
)

func (a *ai) TextToGemini(req *GeminiRequest, breed string) (*GeminiResponse, error) {
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

	contentResp, err := a.sendMsgToGemini(req, instructions.MainInstruction, breedIns)
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

func (a *ai) sendMsgToGemini(req *GeminiRequest, mainInstruction string, subInstruction string) ([]*MessageInfo, error) {
	ctx := context.Background()

	client, err := a.newAiClient(ctx)
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

func (a *ai) newAiClient(ctx context.Context) (*genai.Client, error) {
	var err error
	once.Do(func() {
		client, clientErr := genai.NewClient(ctx, option.WithAPIKey(a.googleConfig.ApiKey))
		if err != nil {
			log.Printf("error new gemini client. Err: %s", clientErr.Error())
			err = fmt.Errorf("new gemini client failed")
			return
		}
		clientInstance = client
	})

	if err != nil {
		return nil, err
	}

	return clientInstance, nil
}
