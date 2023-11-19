package main

import (
	"data-miner/bot"
	"data-miner/db"
	"data-miner/menu"
	"data-miner/messages"
	"data-miner/settings"
	"data-miner/types"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	db.InitDb()

	err := <-bot.Init()
	if err != nil {
		panic(err)
	}

	tgbot := bot.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	menu.InitMenu()

	updates := tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			callback.Text = ""
			tgbot.Request(callback)

			callbackData := strings.Split(update.CallbackQuery.Data, ",")

			switch callbackData[0] {
			case "record":
				if len(callbackData) == 2 {
					insertedRecord, err := db.AddRecord(callbackData[1])
					if err != nil {
						// handle error
					}

					updatedMessage := settings.UpdaterRecord(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, callbackData[1], insertedRecord.Hex())
					_, err = tgbot.Request(updatedMessage)
					spew.Dump(err)
				} else if len(callbackData) == 5 {
					switch callbackData[3] {
					case "rate":
						err := db.UpdateRecord(callbackData[1], callbackData[2], types.Record{
							Rate: callbackData[4],
						})
						if err != nil {
							spew.Dump(err)
						}
					case "note":
						bot.AddNewNoteCreator(update.CallbackQuery.From.ID, callbackData[1], callbackData[2])
						updatedMessage := settings.UpdaterEnterRecordNote(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
						tgbot.Request(updatedMessage)
					default:
						continue
					}
				}
			case "exit":
				deleteMessageRequest := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				tgbot.Request(deleteMessageRequest)
				SendDefaultRecordMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From.ID, tgbot)

				bot.RemoveNewLakeCreator(update.CallbackQuery.From.ID)
				bot.RemoveNewNoteCreator(update.CallbackQuery.From.ID)
			case "data":
				if len(callbackData) == 2 {
					switch callbackData[1] {
					case "back":
						updatedMessage := settings.UpdaterMainSettings(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
						tgbot.Request(updatedMessage)
					case "add":
						bot.AddNewLakeCreator(update.CallbackQuery.From.ID)
						if len(callbackData) == 3 {
							bot.RemoveNewLakeCreator(update.CallbackQuery.From.ID)
							switch callbackData[2] {
							case "back":
								updatedMessage := settings.UpdaterDataSettignsMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, update.CallbackQuery.From.ID)
								tgbot.Request(updatedMessage)
							default:
								continue
							}
						} else {

						}

						updatedMessage := settings.UpdaterEnterNewLakeTitle(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
						tgbot.Request(updatedMessage)
					default:
						continue
					}
				} else {
					updatedMessage := settings.UpdaterDataSettignsMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, update.CallbackQuery.From.ID)
					tgbot.Request(updatedMessage)
				}
			default:
				continue
			}
		} else if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.From.ID, "test")

			switch update.Message.Command() {
			case "start":
				ok, err := db.CreateNewUserOrCheckExist(update.Message.From)
				if err != nil {
					// handle db err
				}
				if ok {
					msg.Text = "Hello"
				} else {
					msg.Text = "Hello again"
				}
			case "settings":
				DeletePrevMessage(update)
				msg = settings.MessageMainSettings(update.Message.From.ID)
				msg.Text = "Settings"
			default:
				continue
			}
			tgbot.Send(msg)
			DeleteCurrentMessage(update)
		} else {
			if bot.CheckNewLakeCreator(update.Message.From.ID) {
				fmt.Println("from lake: ")
				DeletePrevMessage(update)
				bot.RemoveNewLakeCreator(update.Message.From.ID)
				lakeId, _ := db.AddLake(update.Message.Text, update.Message.From.ID)
				db.AddLakeToUser(update.Message.From.ID, lakeId)

				msg := tgbotapi.NewMessage(update.Message.From.ID, "✅")
				sendedMessage, _ := tgbot.Send(msg)

				go func() {
					time.Sleep(1 * time.Second)
					deleteMessageRequest := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, sendedMessage.MessageID)
					tgbot.Request(deleteMessageRequest)

					SendDefaultRecordMessage(update.Message.Chat.ID, update.Message.From.ID, tgbot)
				}()
			}
			if record := bot.CheckNewNoteCreator(update.Message.From.ID); record.UserId != 0 {
				spew.Dump(record)
				fmt.Println("from note: ")

				DeletePrevMessage(update)
				bot.RemoveNewNoteCreator(update.Message.From.ID)
				err := db.UpdateRecord(record.LakeId, record.RecordId, types.Record{
					Note: update.Message.Text,
				})

				if err != nil {
					spew.Dump(err)
				}

				msg := tgbotapi.NewMessage(update.Message.From.ID, "✅")
				sendedMessage, _ := tgbot.Send(msg)

				go func() {
					time.Sleep(1 * time.Second)
					deleteMessageRequest := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, sendedMessage.MessageID)
					tgbot.Request(deleteMessageRequest)

					SendDefaultRecordMessage(update.Message.Chat.ID, update.Message.From.ID, tgbot)
				}()
			}
			DeleteCurrentMessage(update)

		}

	}
}

func SendDefaultRecordMessage(chatId int64, userId int64, tgbot *tgbotapi.BotAPI) {
	msg := messages.MessageRecords(chatId, userId)
	tgbot.Send(msg)
}

func DeletePrevMessage(update tgbotapi.Update) error {
	tgbot := bot.GetBot()
	var deleteMessageRequest tgbotapi.DeleteMessageConfig

	if update.Message != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID-1)
	}
	if update.CallbackQuery != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID-1)
	}
	_, err := tgbot.Request(deleteMessageRequest)

	return err
}

func DeleteCurrentMessage(update tgbotapi.Update) error {
	tgbot := bot.GetBot()
	var deleteMessageRequest tgbotapi.DeleteMessageConfig

	if update.Message != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
	}
	if update.CallbackQuery != nil {
		deleteMessageRequest = tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	}
	_, err := tgbot.Request(deleteMessageRequest)

	return err
}
