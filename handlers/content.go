package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func ContentHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "here is the private content")
	bot.Send(msg)
}
