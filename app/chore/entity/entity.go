package entity

type ChatRequest struct {
	Message string `json:"message" validate:"required"`
	Sender  string `json:"sender"`
}
