package main

import (
	"log"
	"os"
	"test-bot/db"
	"test-bot/handlers"
	"test-bot/middleware"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	// TODO: load all the configs at the same time
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramBotToken == "" {
		log.Fatal("failed to load telegram bot token")
	}

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		log.Fatal("failed to load mongo uri")
	}

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.MongoConnection(mongoUri)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("failed to get updates channel")
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch {
		case update.Message.Command() == "start":
			handlers.StartHandler(update, bot, collection)
		case update.Message.Command() == "content":
			middleware.AuthMiddleware(collection, handlers.ContentHandler)(update, bot)
		case update.Message.Contact != nil:
			handlers.ContactHandler(update, bot, collection)
		}
	}
}
