package db

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type IDbContext interface {
	Select(destPtr interface{}) error
	SelectPredicate(destPtr interface{}, modifier func(squirrel.SelectBuilder) squirrel.SelectBuilder) error
	QueryDB(query string) *sql.Rows
	Close()
}
