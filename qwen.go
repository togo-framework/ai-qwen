// Package qwen is a Qwen (DashScope) driver for togo ai (OpenAI-compatible API).
// Blank-import it and set AI_DRIVER=qwen + DASHSCOPE_API_KEY.
package qwen

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/togo-framework/ai"
	"github.com/togo-framework/togo"
)

const (
	driverName   = "qwen"
	defaultBase  = "https://dashscope.aliyuncs.com/compatible-mode/v1"
	defaultModel = "qwen-plus"
	embedModel   = "text-embedding-v2"
	keyEnv       = "DASHSCOPE_API_KEY"
)

func init() {
	ai.RegisterDriver(driverName, func(k *togo.Kernel) (ai.Provider, error) {
		key := os.Getenv(keyEnv)
		if key == "" && true {
			return nil, errors.New("ai-qwen: " + keyEnv + " not set")
		}
		base := os.Getenv("DASHSCOPE_BASE_URL")
		if base == "" {
			base = defaultBase
		}
		return &provider{key: key, base: base, model: defaultModel, client: &http.Client{Timeout: 120 * time.Second}}, nil
	})
}

type provider struct {
	key, base, model string
	client           *http.Client
}

func (p *provider) Chat(ctx context.Context, req ai.ChatRequest) (ai.ChatResponse, error) {
	model := req.Model
	if model == "" {
		model = p.model
	}
	body := map[string]any{"model": model, "messages": req.Messages}
	if req.Temperature != 0 {
		body["temperature"] = req.Temperature
	}
	if req.MaxTokens != 0 {
		body["max_tokens"] = req.MaxTokens
	}
	var out struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Model string `json:"model"`
		Usage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	}
	if err := p.post(ctx, "/chat/completions", body, &out); err != nil {
		return ai.ChatResponse{}, err
	}
	res := ai.ChatResponse{Model: out.Model, Usage: ai.Usage{PromptTokens: out.Usage.PromptTokens, CompletionTokens: out.Usage.CompletionTokens, TotalTokens: out.Usage.TotalTokens}}
	if len(out.Choices) > 0 {
		res.Content = out.Choices[0].Message.Content
	}
	return res, nil
}

func (p *provider) Embed(ctx context.Context, req ai.EmbedRequest) (ai.EmbedResponse, error) {
	model := req.Model
	if model == "" {
		model = embedModel
	}
	var out struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
		Usage struct {
			PromptTokens int `json:"prompt_tokens"`
			TotalTokens  int `json:"total_tokens"`
		} `json:"usage"`
	}
	if err := p.post(ctx, "/embeddings", map[string]any{"model": model, "input": req.Inputs}, &out); err != nil {
		return ai.EmbedResponse{}, err
	}
	res := ai.EmbedResponse{Usage: ai.Usage{PromptTokens: out.Usage.PromptTokens, TotalTokens: out.Usage.TotalTokens}}
	for _, d := range out.Data {
		res.Vectors = append(res.Vectors, d.Embedding)
	}
	return res, nil
}

func (p *provider) post(ctx context.Context, path string, body, out any) error {
	buf, _ := json.Marshal(body)
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, p.base+path, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	if p.key != "" {
		r.Header.Set("Authorization", "Bearer "+p.key)
	}
	resp, err := p.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return fmt.Errorf("ai-qwen: %s: %s", resp.Status, string(data))
	}
	return json.Unmarshal(data, out)
}
