package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	endpoint string
	apiKey   string
	model    string
}

func (c *Client) Embed(message string) ([]float64, error) {
	reqBody, err := json.Marshal(map[string]interface{}{
		"input":           message,
		"model":           c.model,
		"encoding_format": "float",
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get embeddings: %s", string(bodyBytes))
	}

	var respBody struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if len(respBody.Data) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	return respBody.Data[0].Embedding, nil
}
