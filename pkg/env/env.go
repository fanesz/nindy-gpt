package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	OpenAPIKey       string
	AssistantIDNindy string
)

func InitializeEnv() {
	fmt.Println("===== Initialize .env =====")

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	OpenAPIKey = os.Getenv("open_api_key")
	AssistantIDNindy = os.Getenv("assistant_id_nindy")

	fmt.Println("✓ .env initialized")
}
