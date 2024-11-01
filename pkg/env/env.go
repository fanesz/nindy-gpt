package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	// Server
	BEHost string
	BEPort string

	// Openai
	OpenAPIKey       string
	AssistantIDNindy string
)

func InitializeEnv() {
	fmt.Println("===== Initialize .env =====")

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	BEHost = os.Getenv("BE_HOST")
	BEPort = os.Getenv("BE_PORT")
	OpenAPIKey = os.Getenv("open_api_key")
	AssistantIDNindy = os.Getenv("assistant_id_nindy")

	fmt.Println("✓ .env initialized")
}
