package bost

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//model "github.com/mozgunovdm/omp-bot/internal/model/mdi"
	service "github.com/mozgunovdm/omp-bot/internal/service/mdi/bost"
)

type BostCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
	CallbackList(callback *tgbotapi.CallbackQuery, arg string)
}

type MdiBostCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService service.BostService
}

func NewBostCommander(bot *tgbotapi.BotAPI) BostCommander {
	return &MdiBostCommander{
		bot:              bot,
		subdomainService: service.NewDummyBostService(),
	}
}
