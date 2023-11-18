package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SettingsMainReplyKeyboard() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", "next,yes"),
			tgbotapi.NewInlineKeyboardButtonData("Stop", "next,no"),
		),
	)

	return keyboard
}
