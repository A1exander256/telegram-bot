package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handler) mainMenu(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, "Вы вернулись в главное меню!")
	msg.ReplyMarkup = keyboardMainMenu

	if _, err := h.tgbot.Send(msg); err != nil {
		h.log.Error("error sending message : ", err)
	}
}
