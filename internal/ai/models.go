package ai

import (
	"github.com/google/generative-ai-go/genai"
)

type GeminiRequest struct {
	MessageHistory []*MessageInfo `json:"messageHistory"`
	NewUserMessage string         `json:"newUserMessage"`
}

type MessageInfo struct {
	Message string `json:"message"`
	Role    string `json:"role"`
}

type GeminiResponse struct {
	NewMessageHistory []*MessageInfo `json:"newMessageHistory"`
	ResponseMessage   any            `json:"responseMessage"`
}

func convertMessageHistoryToParts(msgInfo []*MessageInfo) []*genai.Content {
	contents := make([]*genai.Content, 0)

	for _, v := range msgInfo {
		c := &genai.Content{
			Parts: []genai.Part{
				genai.Text(v.Message),
			},
			Role: v.Role,
		}

		contents = append(contents, c)
	}

	return contents
}
