package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SettingsMainReplyKeyboard() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Данные", "data"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выход", "exit"),
		),
	)

	return keyboard
}
