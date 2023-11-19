package settings

import (
	"data-miner/db"
	"data-miner/keyboards"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageMainSettings(chatId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keyboards.SettingsMainReplyKeyboard()

	return msg
}

func MessageDataSettigns(chatId int64, messageId int, userId int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")

	user, _ := db.GetMe(userId)
	lakes, err := db.GetLakes(user.Lakes)
	if err != nil {
		spew.Dump(err)
	}

	msg.ReplyMarkup = keyboards.SettingsDataListKeyboard(lakes)

	return msg
}

func MessageEnterNewLakeTitle(chatId int64, messageId int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keyboards.SettingsAddLakeTitleKeyboard()
	msg.Text = "Введи название для нового типа данных"

	return msg
}
