package models

type Shops struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	UserID string `db:"userid"`
}
