package keyboards

import (
	"data-miner/types"
	"data-miner/utils"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ControlKeyboardRow(callbackPrefix string, withAdd bool) []tgbotapi.InlineKeyboardButton {
	if withAdd {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<<<", callbackPrefix+",back"),
			tgbotapi.NewInlineKeyboardButtonData("x", "exit"),
			tgbotapi.NewInlineKeyboardButtonData("+", callbackPrefix+",add"),
		)
	} else {
		return tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<<<", callbackPrefix+",back"),
			tgbotapi.NewInlineKeyboardButtonData("x", "exit"),
		)
	}
}

func LakesKeyboardRows(lakes []types.Lake, callbackPrefix string) [][]tgbotapi.InlineKeyboardButton {
	var keyboardRows [][]tgbotapi.InlineKeyboardButton

	for i := range lakes {
		keyboardRows = append(keyboardRows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lakes[i].Title, callbackPrefix+","+lakes[i].ID.Hex()),
		))
	}

	return keyboardRows
}

func RateRow(lakeId string, recordId string) []tgbotapi.InlineKeyboardButton {
	var rateRow []tgbotapi.InlineKeyboardButton

	for i := 1; i <= 3; i++ {
		rateRow = append(rateRow, tgbotapi.NewInlineKeyboardButtonData(utils.NumToColor(i), "record,"+lakeId+","+recordId+",rate,"+strconv.Itoa(i)))
	}

	return rateRow
}
