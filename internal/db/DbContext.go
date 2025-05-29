package db

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
	"strings"
	"totmapi/internal/config"
)

type DbContext struct {
	DB *sqlx.DB
}

func NewDbContext() *DbContext {
	connStr := *config.GetDBConfig().ConnectionString
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return &DbContext{DB: db}
}

func (d *DbContext) Select(destPtr interface{}) error {
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("SELECT * FROM %s", table)
	return d.DB.Select(destPtr, q)
}

func (d *DbContext) SelectPredicate(destPtr interface{}, modifier func(sq.SelectBuilder) sq.SelectBuilder) error {
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	sb := sq.Select("*").From(table).
		PlaceholderFormat(sq.Dollar)

	sb = modifier(sb)

	sqlStr, args, err := sb.ToSql()
	if err != nil {
		return err
	}
	return d.DB.Select(destPtr, sqlStr, args...)
}

func (d DbContext) QueryDB(query string) *sql.Rows {
	rows, err := d.DB.Query(query)
	if err != nil {
		log.Printf("Failed to query Db")
		return nil
	}
	return rows
}

func (d DbContext) Close() {
	d.DB.Close()
}

func tableName(ptrToSlice interface{}) (string, error) {
	t := reflect.TypeOf(ptrToSlice)
	if t.Kind() != reflect.Ptr {
		return "", fmt.Errorf("Select: expected pointer to slice, got %T", ptrToSlice)
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		return "", fmt.Errorf("Select: expected pointer to slice, got pointer to %s", t.Kind())
	}
	elem := t.Elem()
	// If you want snake_case, you can use a small util here:
	return toSnakeCase(elem.Name()), nil
}

func toSnakeCase(str string) string {
	var buf strings.Builder
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(r)
	}
	return strings.ToLower(buf.String())
}
