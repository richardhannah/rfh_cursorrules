package models

type Person struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
