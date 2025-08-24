package handlers

import (
	"context"
	"fmt"
	"test-bot/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const Group_ID int64 = 4739390162

func StartHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI, collection *mongo.Collection) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please share your to register")
	button := tgbotapi.NewKeyboardButtonContact("Share Contact")
	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button})

	// TODO: check the reply keyboard render
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup(keyboard)
	bot.Send(msg)
}

func ContactHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI, collection *mongo.Collection) {
	if update.Message.Contact == nil {
		return
	}
	user := update.Message.From
	contact := update.Message.Contact

	newUser := models.User{
		ID:          int64(user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserName:    user.UserName,
		PhoneNumber: contact.PhoneNumber,
		Registerd:   true,
	}

	var searchUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"_id": newUser.ID}).Decode(&searchUser)
	if err == nil && searchUser.Registerd {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "already registered")
		bot.Send(msg)
		return
	}

	collection.InsertOne(context.TODO(), newUser)

	// log mesage to the group
	logMsg := fmt.Sprintf("New Registration:\nUserID: %d\nName: %s %s\nPhone: %s\n", newUser.ID, newUser.FirstName, newUser.LastName, newUser.PhoneNumber)

	bot.Send(tgbotapi.NewMessage(Group_ID, logMsg))
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "registration sucessfull"))
}
