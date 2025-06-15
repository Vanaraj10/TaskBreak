package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Vanaraj10/taskmorph-backend/models"
)

type Step struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"` // Completed indicates if the step is done
}

func AskGemini(taskTitle string) ([]models.Step, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY is missing")
	}

	prompt := fmt.Sprintf(`Break the task "%s" into 5 short steps. 
Respond ONLY as raw JSON array: 
[{"title": "Step 1", "description": "..."}, ...]`, taskTitle)

	reqBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
				"role": "user",
			},
		},
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key="+apiKey, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	// 🔎 Print raw response
	fmt.Println("🧪 Gemini Raw:", string(bodyBytes))

	var parsed map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &parsed); err != nil {
		return nil, err
	}

	candidates, ok := parsed["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return nil, fmt.Errorf("No candidates returned")
	}

	content := candidates[0].(map[string]interface{})["content"]
	parts := content.(map[string]interface{})["parts"].([]interface{})
	text := parts[0].(map[string]interface{})["text"].(string)

	// 🧹 Clean Markdown wrapper
	clean := strings.TrimSpace(text)
	clean = strings.TrimPrefix(clean, "```json")
	clean = strings.TrimSuffix(clean, "```")
	clean = strings.TrimSpace(clean)

	// ✅ Parse the cleaned JSON string
	var steps []models.Step
	if err := json.Unmarshal([]byte(clean), &steps); err != nil {
		return nil, fmt.Errorf("Failed to parse steps: %v", err)
	}

	return steps, nil
}