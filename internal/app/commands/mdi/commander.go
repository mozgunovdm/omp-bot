package mdi

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mozgunovdm/omp-bot/internal/app/commands/mdi/bost"
	"github.com/mozgunovdm/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type MdiCommander struct {
	bot *tgbotapi.BotAPI
	com bost.BostCommander
}

func NewMdiCommander(
	bot *tgbotapi.BotAPI,
) *MdiCommander {
	return &MdiCommander{
		bot: bot,
		com: bost.NewBostCommander(bot),
	}
}

func (c *MdiCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	c.com.CallbackList(callback, callbackPath.CallbackData)
}

func (c *MdiCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("MdiCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *MdiCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.com.Help(msg)
	case "list":
		c.com.List(msg)
	case "get":
		c.com.Get(msg)
	case "delete":
		c.com.Delete(msg)
	case "new":
		c.com.New(msg)
	case "edit":
		c.com.Edit(msg)
	default:
	}
}
