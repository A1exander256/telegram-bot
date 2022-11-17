package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) start(chatId int64) {
	textMsg := "Привет! Я очень рад, что вы пользуетесь этим сервисом!\nВыберите действие..."
	msg := tgbotapi.NewMessage(chatId, textMsg)
	msg.ReplyMarkup = keyboardMainMenu
	if _, err := h.tgbot.Send(msg); err != nil {
		h.log.Error("error sending message : ", err)
	}
}
