package bost

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mozgunovdm/omp-bot/internal/app/path"
)

func (c *MdiBostCommander) CallbackList(callback *tgbotapi.CallbackQuery, msg string) {
	argsData := strings.Fields(msg)
	if len(argsData) != 2 {
		log.Printf("MdiBostCommander.CallbackList: error number of args")
		return
	}

	from, err := strconv.Atoi(argsData[0])
	if err != nil {
		log.Println("MdiBostCommander.CallbackList: wrong args - ", msg)
		return
	}

	limit, err := strconv.Atoi(argsData[1])
	if err != nil {
		log.Println("MdiBostCommander.CallbackList: wrong args - ", msg)
		return
	}

	outputMsgText := fmt.Sprintf("Products from %d count %d: \n\n", from, limit)

	var msgOut tgbotapi.MessageConfig
	products, err := c.subdomainService.List(uint64(from), uint64(limit))
	if nil != err {
		log.Printf("MdiBostCommander.CallbackList: %v", err)
		msgOut = tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("%v", err))
	} else {
		for _, p := range products {
			outputMsgText += p.Name
			outputMsgText += "\n"
		}

		msgOut = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

		from += limit
		callbackPath := path.CallbackPath{
			Domain:       "mdi",
			Subdomain:    "bost",
			CallbackName: "list",
			CallbackData: fmt.Sprintf("%d %d", from, limit),
		}

		msgOut.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msgOut)
	if err != nil {
		log.Printf("MdiBostCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
