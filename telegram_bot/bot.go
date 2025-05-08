package telegrambot

import (
	"encoding/json"
	classroomauths "glassroom/classroom_auths"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	APIKey string `json:"api_key"` // structs tags
	ChatID int64  `json:"chat_id"`
}

func RunBot() {
	config, err := getConfig("telegram_api_token.json")
	if err != nil {
		log.Fatalf("Couldn't read token from local file: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.APIKey) // or just `token` if not using a pointer
	if err != nil {
		log.Fatalf("Couldn't initialize bot: %v", err)
	}

	client := classroomauths.GetClient()
	srv, err := classroomauths.CreateServiceToClassroom(client)
	if err != nil {
		log.Fatalf("Couldn't create a service to classroom: %v", err)
	}
	data, err := classroomauths.ListCourses(srv)
	if err != nil {
		log.Fatalf("Couldn't fetch data: %v", err)
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
