package settings

import (
	"data-miner/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MainSettingsMessage(chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keyboards.SettingsMainReplyKeyboard()

	return msg
}
