package models

type ShopStock struct {
	ShopID            string `db:"shopid"`
	Name              string `db:"name"`
	Description       string `db:"description"`
	Encumbrance       int    `db:"encumbrance"`
	Unit              string `db:"unit"`
	QuantityAvailable int    `db:"quantity_available"`
	Category          string `db:"category"`
	Cost              int    `db:"cost"`
}
