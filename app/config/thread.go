package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Data struct {
	ThreadID string `json:"thread_id"`
}

var filename = "app/config/session.json"

func InitializeThread() {
	ctx := context.Background()
	client := GetClient()

	threadID, err := GetThreadID()
	if err != nil {
		log.Fatal(err)
	}
	if threadID == "" {
		thread, err := client.CreateThread(ctx, openai.ThreadRequest{})
		if err != nil {
			log.Fatal(err)
		}

		err = EditThreadID(thread.ID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		run, err := client.RetrieveThread(ctx, threadID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(run.ID)
	}
}

func GetThreadID() (string, error) {
	var data Data
	err := readJSONFile(&data)
	if err != nil {
		return "", err
	}
	return data.ThreadID, nil
}

func EditThreadID(newThreadID string) error {
	var data Data
	err := readJSONFile(&data)
	if err != nil {
		return err
	}

	data.ThreadID = newThreadID

	newByteValue, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, newByteValue, 0644)
}

func readJSONFile(data *Data) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, data)
}
