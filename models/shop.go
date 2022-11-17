package models

type Shop struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
