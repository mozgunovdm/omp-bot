package bost

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mozgunovdm/omp-bot/internal/model/mdi"
)

func (c *MdiBostCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	argsData := strings.Fields(args)
	if len(argsData) != 2 {
		log.Printf("MdiBostCommander.Edit: error number of args")
		return
	}

	idx, err := strconv.Atoi(argsData[0])
	if err != nil {
		log.Println("MdiBostCommander.Edit: wrong args - ", args)
		return
	}

	var outputMsgText string

	err = c.subdomainService.Update(uint64(idx), mdi.Bost{Name: argsData[1]})
	if err != nil {
		outputMsgText = fmt.Sprintf("%v", err)
	} else {
		outputMsgText = fmt.Sprintf("Product %d changed to %s", idx, argsData[1])
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("MdiBostCommander.Edit: error sending reply message to chat - %v", err)
	}
}
