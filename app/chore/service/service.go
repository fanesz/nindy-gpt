package service

import (
	"context"
	"nindy-gpt/app/chore/interfaces"
	"nindy-gpt/app/config"
	"nindy-gpt/pkg/env"
	"time"

	"github.com/sashabaranov/go-openai"
)

var _ interfaces.NindyGPTService = &nindyGPTService{}

type nindyGPTService struct {
	client  *openai.Client
	context context.Context
}

func NewNindyGPTService(client *openai.Client, ctx context.Context) interfaces.NindyGPTService {
	return &nindyGPTService{
		client:  client,
		context: ctx,
	}
}

func (s *nindyGPTService) Chat(message string) (string, error) {
	threadID, err := config.GetThreadID()
	if err != nil {
		return "", err
	}

	_, err = s.client.CreateMessage(s.context, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	if err != nil {
		return "", err
	}

	run, err := s.client.CreateRun(s.context, threadID, openai.RunRequest{
		AssistantID: env.AssistantIDNindy,
	})
	if err != nil {
		return "", err
	}

	for run.Status == openai.RunStatusQueued || run.Status == openai.RunStatusInProgress {
		run, err = s.client.RetrieveRun(s.context, run.ThreadID, run.ID)
		if err != nil {
			return "", err
		}
		time.Sleep(100 * time.Millisecond)
	}
	if run.Status != openai.RunStatusCompleted {
		return "", err
	}

	numMessages := 1
	messages, err := s.client.ListMessage(s.context, run.ThreadID, &numMessages, nil, nil, nil, nil)
	if err != nil {
		return "", err
	}

	return messages.Messages[0].Content[0].Text.Value, nil
}
