package main

import (
	"flag"
	"fmt"
	classroomauths "glassroom/classroom_auths"
	"log"
	telegrambot "glassroom/telegram_bot"
)

func main() {
	setup := flag.Bool("setup", false, "Run setup function to get token from classroomApi")
	bot := flag.Bool("run", true, "Run the telegram bot")

	flag.Parse()

	switch {
	case *setup:
		fmt.Println("Running setup for Google Classroom...")
		client := classroomauths.GetClient()
	case *bot:
		fmt.Println("Running Telegram Bot...")
		fmt.Printf("Connecting to chatID: %v", 6)
		if err := telegrambot.RunBot(); err != nil {
			log.Fatalf("Bot failed to run:")
		}
	}

}
