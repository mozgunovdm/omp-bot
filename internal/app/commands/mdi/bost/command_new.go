package bost

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mozgunovdm/omp-bot/internal/model/mdi"
)

func (c *MdiBostCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	outputMsgText := "Added new product №: \n"

	idx, err := c.subdomainService.Create(mdi.Bost{Name: args})
	if err != nil {
		log.Printf("MdiBostCommander.Create: error create new product - %v", err)
		return
	}

	outputMsgText += strconv.Itoa(int(idx))

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("MdiBostCommander.Create: error sending reply message to chat - %v", err)
	}
}
