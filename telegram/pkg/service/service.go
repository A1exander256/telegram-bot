package service

import (
	"github.com/a1exander256/telegram-bot/models"
	"github.com/a1exander256/telegram-bot/telegram/pkg/storage"
	"github.com/sirupsen/logrus"
)

type Product interface {
}

type Shop interface {
	GetAll() ([]models.Shop, error)
}

type Service struct {
	Product
	Shop
}

func NewService(storage *storage.Storage, log *logrus.Logger) *Service {
	return &Service{
		Shop: NewShopService(storage.Shop, log),
	}
}
