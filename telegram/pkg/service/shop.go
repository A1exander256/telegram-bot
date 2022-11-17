package service

import (
	"github.com/a1exander256/telegram-bot/models"
	"github.com/a1exander256/telegram-bot/telegram/pkg/storage"
	"github.com/sirupsen/logrus"
)

type ShopService struct {
	storage storage.Shop
	log     *logrus.Logger
}

func NewShopService(storage storage.Shop, log *logrus.Logger) *ShopService {
	return &ShopService{
		storage: storage,
		log:     log,
	}
}

func (s *ShopService) GetAll() ([]models.Shop, error) {
	return s.storage.GetAll()
}
