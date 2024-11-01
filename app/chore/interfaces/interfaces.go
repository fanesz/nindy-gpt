package interfaces

type NindyGPTService interface {
	Chat(message string) (string, error)
}
