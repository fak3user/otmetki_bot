package keyboards

import (
	"data-miner/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MainRecordKeyboard(userId int64) tgbotapi.InlineKeyboardMarkup {
	user, _ := db.GetMe(userId)
	lakes, _ := db.GetLakes(user.Lakes)

	lakesRows := LakesKeyboardRows(lakes, "record")
	keyboard := tgbotapi.NewInlineKeyboardMarkup()

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, lakesRows...)

	return keyboard
}

func AdditionalRecordInfoKeyboard(lakeId string, recordId string) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup()
	rateRow := RateRow(lakeId, recordId)
	noteButton := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Заметка", "record,"+lakeId+","+recordId+",note,1"),
	)
	closeButton := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("x", "exit"),
	)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rateRow)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, noteButton)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, closeButton)

	return keyboard
}
