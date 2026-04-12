package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const apiURL = "https://openrouter.ai/api/v1/chat/completions"

// models is the fallback list — tried in order until one succeeds.
var models = []string{
	"deepseek/deepseek-chat-v3.1:free",
	"meta-llama/llama-4-maverick:free",
	"qwen/qwen3-235b-a22b:free",
	"google/gemma-3-27b-it:free",
	"openrouter/free",
}

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

// GetAPIKey reads OPENROUTER_API_KEY from the environment.
func GetAPIKey() (string, error) {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY not found. Run: pilot setup")
	}
	return key, nil
}

// SetPreferredModel prepends a user-configured model to the top of the fallback list.
func SetPreferredModel(model string) {
	if model == "" {
		return
	}
	// Remove if already present, then prepend
	filtered := make([]string, 0, len(models))
	for _, m := range models {
		if m != model {
			filtered = append(filtered, m)
		}
	}
	models = append([]string{model}, filtered...)
}

// ask sends a single request to the given model.
func ask(apiKey, model, system, prompt string) (string, error) {
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
		return "", fmt.Errorf("connection error: %v", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	var r response
	if err := json.Unmarshal(raw, &r); err != nil {
		return "", fmt.Errorf("parse error: %s", string(raw))
	}
	if r.Error != nil {
		return "", fmt.Errorf("no endpoints")
	}
	if len(r.Choices) == 0 {
		return "", fmt.Errorf("empty response")
	}
	return r.Choices[0].Message.Content, nil
}

// isValid checks that the response contains expected output markers.
func isValid(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return false
	}
	return strings.Contains(s, "```") ||
		strings.Contains(s, "📌") ||
		strings.Contains(s, "🔍") ||
		strings.Contains(s, "📦")
}

// Ask tries each model in the fallback list until a valid response is returned.
func Ask(apiKey, system, prompt string) (string, error) {
	var lastErr error
	for _, m := range models {
		result, err := ask(apiKey, m, system, prompt)
		if err != nil && strings.Contains(err.Error(), "no endpoints") {
			lastErr = err
			continue
		}
		if err != nil {
			return "", err
		}
		if !isValid(result) {
			lastErr = fmt.Errorf("invalid response from %s", m)
			continue
		}
		return result, nil
	}
	return "", fmt.Errorf("all models failed: %v", lastErr)
}