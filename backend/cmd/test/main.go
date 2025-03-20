package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"auto-messenger/pkg/callmebot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Error] (godotenv):" + err.Error())
	}

	phoneNumber := os.Getenv("CMB_NUMBER")
	apiKey := os.Getenv("CMB_API_KEY")

	if phoneNumber == "" || apiKey == "" {
		log.Fatal("[Error] (CMB_NUMBER and CMB_API_KEY must be set in .env file)")
	}

	client := callmebot.NewClient(callmebot.Config{
		PhoneNumber: phoneNumber,
		ApiKey:      apiKey,
	})

	err = client.SendMessage(context.Background(), callmebot.Message{
		Content: "This is a test message from CallMeBot",
	})

	if err != nil {
		fmt.Printf("[Error] (client.SendMessage): %v\n", err)
	} else {
		fmt.Println("[Success] (client.SendMessage)")
	}
}
