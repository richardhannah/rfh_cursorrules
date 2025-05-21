package models

import (
	"database/sql"
)

type Blogposts struct {
	BlogpostID string         `db:"blogpostid"`
	Title      sql.NullString `db:"title"`
	Markdown   sql.NullString `db:"markdown"`
	Category   sql.NullString `db:"category"`
	Image      sql.NullString `db:"image"`
	Video      sql.NullString `db:"video"`
	Date       sql.NullTime   `db:"date"`
	Published  bool           `db:"published"`
}
