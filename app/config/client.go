package config

import (
	"nindy-gpt/pkg/env"

	"github.com/sashabaranov/go-openai"
)

var client *openai.Client

func InitializeClient() {
	client = openai.NewClient(env.OpenAPIKey)
	if client == nil {
		panic("Failed to initialize OpenAI client")
	}
}

func GetClient() *openai.Client {
	if client == nil {
		InitializeClient()
	}

	return client
}
