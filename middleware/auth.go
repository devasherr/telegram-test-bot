package middleware

import (
	"context"
	"test-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type HandlerFunc func(update tgbotapi.Update, bot *tgbotapi.BotAPI)

func AuthMiddleware(collection *mongo.Collection, next HandlerFunc) HandlerFunc {
	return func(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
		userID := update.Message.From.ID

		var user models.User
		err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
		if err != nil || !user.Registerd {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "you are not registered"))
			return
		}

		next(update, bot)
	}
}
