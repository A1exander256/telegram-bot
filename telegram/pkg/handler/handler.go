package handler

import (
	"github.com/a1exander256/telegram-bot/telegram/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var (
	keyboardMainMenu = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Наши магазины"),
	))

	rowMainMenu = tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Главное меню"),
	)
)

type Handler struct {
	log     *logrus.Logger
	tgbot   *tgbotapi.BotAPI
	service *service.Service
}

func NewHandler(service *service.Service, tgbot *tgbotapi.BotAPI, log *logrus.Logger) *Handler {
	return &Handler{
		log:     log,
		tgbot:   tgbot,
		service: service,
	}
}

func (h *Handler) ListenUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := h.tgbot.GetUpdatesChan(u)

	for update := range updates {
		chatId := update.Message.Chat.ID
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				h.start(chatId)
			case "Главное меню":
				h.mainMenu(chatId)
			case "Наши магазины":
				h.ourShops(chatId)
			}
		}
	}
}
