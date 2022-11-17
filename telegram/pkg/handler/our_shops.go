package handler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) ourShops(chatId int64) {
	shops, err := h.service.Shop.GetAll()
	if err != nil {
		h.log.Error("error getting our stores from postgres : ", err)
		h.errorResponse(chatId, "Ой, что-то пошло не по плану(")
		return
	}

	rowShop := tgbotapi.NewKeyboardButtonRow()
	textMsg := "Наши магазины :"
	for _, shop := range shops {
		rowShop = append(rowShop, tgbotapi.NewKeyboardButton(shop.Name))
		textMsg += fmt.Sprintf("\n%s (%s)", shop.Name, shop.Description)
	}

	keybord := tgbotapi.NewReplyKeyboard(rowShop, rowMainMenu)
	msg := tgbotapi.NewMessage(chatId, textMsg)
	msg.ReplyMarkup = keybord
	if _, err := h.tgbot.Send(msg); err != nil {
		h.log.Error("error sending message : ", err)
	}
}
