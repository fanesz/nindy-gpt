package database

import (
	"log"
)

func Insert(sender, message, response string) {
	db := GetDatabase()

	_, err := db.Exec(
		"INSERT INTO chat_history (sender, message, response) VALUES (?, ?, ?)",
		sender,
		message,
		response,
	)
	if err != nil {
		log.Fatal(err)
	}
}
