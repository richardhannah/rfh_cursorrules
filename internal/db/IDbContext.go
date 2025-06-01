package db

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type IDbContext interface {
	SelectAll(destPtr interface{}) error
	Select(destPtr interface{}, modifier func(squirrel.SelectBuilder) squirrel.SelectBuilder) error
	QueryDB(query string) *sql.Rows
	Close()
}
