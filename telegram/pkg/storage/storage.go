package storage

import (
	"github.com/a1exander256/telegram-bot/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Product interface {
}

type Shop interface {
	GetAll() ([]models.Shop, error)
}
type Storage struct {
	Product
	Shop
}

func NewStorage(db *sqlx.DB, log *logrus.Logger) *Storage {
	return &Storage{
		Shop: NewShopStorage(db, log),
	}
}
