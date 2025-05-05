package main

import (
	"encoding/json"
	"fmt"
	"glassroom"
	"log"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	APIKey string `json:"api_key"` // structs tags
	ChatID int64  `json:"chat_id"`
}

func main() {
	config, err := getConfig("src/token.json")
	if err != nil {
		log.Fatalf("Couldn't read token from local file: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.APIKey) // or just `token` if not using a pointer
	if err != nil {
		log.Fatalf("Couldn't initialize bot: %v", err)
	}

	data, err := glassroom.FetchClassInfo()
	if err != nil {
		log.Fatalf("Couldn't fetch data from Google Classroom: %v", err)
	}

	for _, c := range data.Courses {
		msg := tgbotapi.NewMessage(int64(config.ChatID), c.Name) // use course name or any relevant field
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func getConfig(file string) (*Config, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	config := &Config{}
	// decoder := json.NewDecoder(f)
	// if err := decoder.Decode(&config); err != nil {
	// 	return nil, err
	// }
	err = json.NewDecoder(f).Decode(config)
	return config, nil

}
