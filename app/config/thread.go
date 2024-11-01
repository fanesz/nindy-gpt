package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Data struct {
	ThreadID string `json:"thread_id"`
}

var filename = "session.json"

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
		_, err := client.RetrieveThread(ctx, threadID)
		if err != nil {
			log.Fatal(err)
		}
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
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("session.json not found, creating a new one...")
			file, err = os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()

			defaultData := Data{ThreadID: ""}
			data = &defaultData
			jsonData, err := json.MarshalIndent(defaultData, "", "  ")
			if err != nil {
				return err
			}
			_, err = file.Write(jsonData)
			if err != nil {
				return err
			}

			return nil
		}
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, data)
}
