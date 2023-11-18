package main

import (
	"data-miner/bot"
	"data-miner/db"
	"data-miner/menu"
	"data-miner/settings"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.InitDb()

	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	tgbot := bot.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	menu.InitMenu()

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			callback.Text = ""
			tgbot.Request(callback)

			callbackData := strings.Split(update.CallbackQuery.Data, ",")

			switch callbackData[0] {
			case "exit":
				deleteMessageRequest := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				tgbot.Request(deleteMessageRequest)
			case "data":

			default:
				continue
			}
		} else if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.From.ID, "test")

			switch update.Message.Command() {
			case "start":
				ok, err := db.CreateNewUserOrCheckExist(update.Message.From)
				if err != nil {
					// handle db err
				}
				if ok {
					msg.Text = "Hello"
				} else {
					msg.Text = "Hello again"
				}
			case "settings":
				msg = settings.MainSettingsMessage(update.Message.From.ID)
				msg.Text = "Settings"
			default:
				continue
			}
			tgbot.Send(msg)
			DeleteCurrentMessage(update)
		} else {
			DeleteCurrentMessage(update)

		}

	}
}

func DeletePrevMessage(update tgbotapi.Update) error {
	tgbot := bot.GetBot()
	var deleteMessageRequest tgbotapi.DeleteMessageConfig

	if update.Message != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID-1)
	}
	if update.CallbackQuery != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID-1)
	}
	_, err := tgbot.Request(deleteMessageRequest)

	return err
}

func DeleteCurrentMessage(update tgbotapi.Update) error {
	tgbot := bot.GetBot()
	var deleteMessageRequest tgbotapi.DeleteMessageConfig

	if update.Message != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
	}
	if update.CallbackQuery != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	}
	_, err := tgbot.Request(deleteMessageRequest)

	return err
}
