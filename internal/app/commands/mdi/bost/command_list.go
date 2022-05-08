package bost

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mozgunovdm/omp-bot/internal/app/path"
)

func (c *MdiBostCommander) List(inputMsg *tgbotapi.Message) {
	fmt.Println(inputMsg)
	args := inputMsg.CommandArguments()

	argsData := strings.Fields(args)
	if len(argsData) != 2 {
		log.Printf("MdiBostCommander.List: error number of args")
		return
	}

	from, err := strconv.Atoi(argsData[0])
	if err != nil {
		log.Println("MdiBostCommander.List: wrong args - ", args)
		return
	}

	limit, err := strconv.Atoi(argsData[1])
	if err != nil {
		log.Println("MdiBostCommander.List: wrong args - ", args)
		return
	}

	outputMsgText := fmt.Sprintf("Products from %d count %d: \n\n", from, limit)

	var msg tgbotapi.MessageConfig
	products, err := c.subdomainService.List(uint64(from), uint64(limit))
	if nil != err {
		log.Printf("MdiBostCommander.List: %v", err)
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%v", err))
		return
	} else {
		for _, p := range products {
			outputMsgText += p.Name
			outputMsgText += "\n"
		}

		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

		from += limit
		callbackPath := path.CallbackPath{
			Domain:       "mdi",
			Subdomain:    "bost",
			CallbackName: "list",
			CallbackData: fmt.Sprintf("%d %d", from, limit),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("MdiBostCommander.List: error sending reply message to chat - %v", err)
	}
}
