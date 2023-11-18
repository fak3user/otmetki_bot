package main

import (
	"data-miner/bot"
	"data-miner/db"
	"data-miner/menu"
	"data-miner/schedule"
	"data-miner/settings"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.InitDb()
	schedule.InitSchedule()

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
		msg := tgbotapi.NewMessage(update.Message.From.ID, "test")

		if update.Message.IsCommand() {
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
		}

		tgbot.Send(msg)
	}
}
