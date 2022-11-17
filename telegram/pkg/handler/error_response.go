package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) errorResponse(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)
	if _, err := h.tgbot.Send(msg); err != nil {
		h.log.Error("error sending message : ", err)
	}
}
