package keyboards

import (
	"data-miner/types"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SettingsMainReplyKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Данные", "data"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выход", "exit"),
		),
	)
}

func SettingsDataListKeyboard(lakes []types.Lake) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup()
	lakesRows := LakesKeyboardRows(lakes, "data")

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, lakesRows...)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, ControlKeyboardRow("data", true))

	return keyboard
}

func SettingsAddLakeTitleKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		ControlKeyboardRow("data,add", false),
	)
}
