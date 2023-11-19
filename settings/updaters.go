package settings

import (
	"data-miner/db"
	"data-miner/keyboards"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UpdaterMainSettings(chatId int64, messageId int) tgbotapi.EditMessageReplyMarkupConfig {
	newKeyboard := keyboards.SettingsMainReplyKeyboard()

	return tgbotapi.EditMessageReplyMarkupConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   messageId,
			ReplyMarkup: &newKeyboard,
		},
	}
}

func UpdaterDataSettignsMessage(chatId int64, messageId int, userId int64) tgbotapi.EditMessageReplyMarkupConfig {
	user, _ := db.GetMe(userId)
	lakes, err := db.GetLakes(user.Lakes)
	if err != nil {
		spew.Dump(err)
	}

	newKeyboard := keyboards.SettingsDataListKeyboard(lakes)

	return tgbotapi.EditMessageReplyMarkupConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   messageId,
			ReplyMarkup: &newKeyboard,
		},
	}
}

func UpdaterEnterNewLakeTitle(chatId int64, messageId int) tgbotapi.EditMessageTextConfig {
	newKeyboard := keyboards.SettingsAddLakeTitleKeyboard()

	return tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   messageId,
			ReplyMarkup: &newKeyboard,
		},
		Text: "Введи название для нового типа данных",
	}
}

func UpdaterEnterRecordNote(chatId int64, messageId int) tgbotapi.EditMessageTextConfig {
	return tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   messageId,
			ReplyMarkup: nil,
		},
		Text: "Введи текст заметки",
	}
}

func UpdaterRecord(chatId int64, messageId int, lakeId string, recordId string) tgbotapi.EditMessageTextConfig {
	newKeyboard := keyboards.AdditionalRecordInfoKeyboard(lakeId, recordId)

	return tgbotapi.EditMessageTextConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      chatId,
			MessageID:   messageId,
			ReplyMarkup: &newKeyboard,
		},
		Text: "Дополнительная информация",
	}
}
