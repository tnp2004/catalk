package ai

type GeminiTextBody struct {
	Message string `json:"message"`
}

type GeminiResponse struct {
	Message any `json:"message"`
}
