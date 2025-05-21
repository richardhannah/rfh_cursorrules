package models

import (
	"database/sql"
)

type Users struct {
	ID        string         `db:"id"`
	Username  string         `db:"username"`
	Password  string         `db:"password"`
	Salt      string         `db:"salt"`
	Ipaddress sql.NullString `db:"ipaddress"`
	Enabled   sql.NullBool   `db:"enabled"`
	Role      sql.NullString `db:"role"`
}
