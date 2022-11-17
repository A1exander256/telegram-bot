package storage

import (
	"fmt"

	"github.com/a1exander256/telegram-bot/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ShopStorage struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewShopStorage(db *sqlx.DB, log *logrus.Logger) *ShopStorage {
	return &ShopStorage{
		db:  db,
		log: log,
	}
}
func (s *ShopStorage) GetAll() ([]models.Shop, error) {
	var shops []models.Shop

	/*query := fmt.Sprintf(`SELECT shops.name as shop_name, products.name as product_name, products_count.count FROM %s products_count
	LEFT JOIN %s shops ON products_count.shop_id=shops.id
	LEFT JOIN %s products ON products_count.product_id=products.id
	WHERE shops.name = $1`,
		tableProductsCount, tableShops, tableProducts)
	*/
	query := fmt.Sprintf(`SELECT * FROM %s`, tableShops)
	err := s.db.Select(&shops, query)
	return shops, err
}
