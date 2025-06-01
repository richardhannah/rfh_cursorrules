package db

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type IDbContext interface {
	SelectAll(destPtr interface{}) error
	Select(destPtr interface{}, modifier func(squirrel.SelectBuilder) squirrel.SelectBuilder) error
	Insert(destPtr interface{}, modifier func(squirrel.InsertBuilder) squirrel.InsertBuilder) error
	Update(destPtr interface{}, modifier func(squirrel.UpdateBuilder) squirrel.UpdateBuilder) error
	Delete(destPtr interface{}, modifier func(squirrel.DeleteBuilder) squirrel.DeleteBuilder) error
	QueryDB(query string) *sql.Rows
	Close()
}
