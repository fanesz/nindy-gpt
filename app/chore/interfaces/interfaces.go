package interfaces

import "nindy-gpt/app/chore/entity"

type NindyGPTService interface {
	Chat(req *entity.ChatRequest) (string, error)
}
