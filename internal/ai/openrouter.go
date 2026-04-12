package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	apiURL = "https://openrouter.ai/api/v1/chat/completions"
	model  = "openrouter/free"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func GetAPIKey() (string, error) {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY bulunamadı. Kurmak için: pilot setup")
	}
	return key, nil
}

func Ask(apiKey, system, prompt string) (string, error) {
	req := request{
		Model: model,
		Messages: []message{
			{Role: "system", Content: system},
			{Role: "user", Content: prompt},
		},
	}

	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequest("POST", apiURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("bağlantı hatası: %v", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	var r response
	if err := json.Unmarshal(raw, &r); err != nil {
		return "", fmt.Errorf("yanıt parse hatası: %s", string(raw))
	}
	if r.Error != nil {
		return "", fmt.Errorf("AI hatası: %s", r.Error.Message)
	}
	if len(r.Choices) == 0 {
		return "", fmt.Errorf("boş yanıt alındı")
	}
	return r.Choices[0].Message.Content, nil
}