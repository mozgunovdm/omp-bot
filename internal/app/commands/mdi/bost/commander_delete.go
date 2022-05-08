package bost

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *MdiBostCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("MdiBostCommander.Delete: wrong args", args)
		return
	}

	var outputMsgText string

	res, err := c.subdomainService.Remove(uint64(idx))
	if err != nil {
		outputMsgText = fmt.Sprintf("%v", err)
	} else if res {
		outputMsgText = fmt.Sprintf("Product %d deleted", idx)
	} else {
		outputMsgText = fmt.Sprintf("Couldn't deleted product %d", idx)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("MdiBostCommander.Delete: error sending reply message to chat - %v", err)
	}
}
