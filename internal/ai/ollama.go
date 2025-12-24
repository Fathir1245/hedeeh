package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const OllamaURL = "http://localhost:11434/api/generate"
const ModelName = "qwen2.5-coder:7b"

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func GenerateCode(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  ModelName,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("gagal konek ke Ollama: %w", err)
	}
	defer resp.Body.Close()

	var result OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	cleanCode := cleanMarkdown(result.Response)
	return cleanCode, nil
}

func cleanMarkdown(raw string) string {
	raw = strings.ReplaceAll(raw, "```go", "")
	raw = strings.ReplaceAll(raw, "```", "")
	return strings.TrimSpace(raw)
}
