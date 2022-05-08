package bost

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *MdiBostCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get id - get product by id\n"+
			"/delete id - delete product by id\n"+
			"/new name - add new product\n"+
			"/edit id name - set new name by id",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("MdiBostCommander.Help: error sending reply message to chat - %v", err)
	}
}
