package messages

import (
	"data-miner/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageRecords(chatId int64, userId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "Что ты хочешь отметить сейчас?")
	msg.ReplyMarkup = keyboards.MainRecordKeyboard(userId)

	return msg
}
