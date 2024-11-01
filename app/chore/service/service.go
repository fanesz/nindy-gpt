package service

import (
	"context"
	"nindy-gpt/app/chore/entity"
	"nindy-gpt/app/chore/interfaces"
	"nindy-gpt/app/config"
	"nindy-gpt/app/database"
	"nindy-gpt/pkg/env"
	"strings"
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

func (s *nindyGPTService) Chat(req *entity.ChatRequest) (string, error) {
	threadID, err := config.GetThreadID()
	if err != nil {
		return "", err
	}

	messageRequest := openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Message,
	}

	// add metadata if sender is not empty
	if req.Sender != "" {
		messageRequest.Metadata = map[string]any{
			"user_name": req.Sender,
		}
	}

	// create message to current active thread
	_, err = s.client.CreateMessage(s.context, threadID, messageRequest)
	if err != nil {
		return "", err
	}

	// sending message to current active thread
	run, err := s.client.CreateRun(s.context, threadID, openai.RunRequest{
		AssistantID: env.AssistantIDNindy,
	})
	if err != nil {
		return "", err
	}

	// wait until the run is completed
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

	// get the response
	numMessages := 1
	messages, err := s.client.ListMessage(s.context, run.ThreadID, &numMessages, nil, nil, nil, nil)
	if err != nil {
		return "", err
	}

	response := messages.Messages[0].Content[0].Text.Value

	// insert chat history
	go database.Insert(req.Sender, req.Message, response)

	// replace placeholders
	placeholders := []string{"[user_name]", "{user_name}", "<user_name>", "(user_name)"}
	for _, placeholder := range placeholders {
		response = strings.ReplaceAll(response, placeholder, req.Sender)
	}

	return response, nil
}
