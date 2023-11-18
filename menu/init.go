package menu

import (
	"data-miner/bot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitMenu() {
	bot := bot.GetBot()

	cfg := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "/quiz",
			Description: "Start quiz",
		},
		tgbotapi.BotCommand{
			Command:     "/settings",
			Description: "Settings",
		},
	)

	_, _ = bot.Request(cfg)
}
