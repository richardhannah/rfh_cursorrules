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

func (d *DbContext) SelectAll(destPtr interface{}) error {
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("SELECT * FROM %s", table)
	fmt.Println(q)
	return d.DB.Select(destPtr, q)
}

func (d *DbContext) Select(destPtr interface{}, modifier func(sq.SelectBuilder) sq.SelectBuilder) error {
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	sb := sq.Select("*").From(table).
		PlaceholderFormat(sq.Dollar)

	sb = modifier(sb)

	sqlStr, args, err := sb.ToSql()
	fmt.Println(sqlStr)
	if err != nil {
		return err
	}
	return d.DB.Select(destPtr, sqlStr, args...)
}

func (d *DbContext) Insert(destPtr interface{}, modifier func(sq.InsertBuilder) sq.InsertBuilder) error {
	// 1) Figure out the table name from destPtr
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	// 2) Start an InsertBuilder for that table, using PostgreSQL-style placeholders
	ib := sq.Insert(table).PlaceholderFormat(sq.Dollar)

	// 3) Let the caller add Columns(...) / Values(...) / Returning(...) etc.
	ib = modifier(ib)

	// 4) Generate SQL + args
	sqlStr, args, err := ib.ToSql()
	if err != nil {
		return fmt.Errorf("Insert: failed to build SQL: %w", err)
	}

	// 5) Exec the INSERT
	_, err = d.DB.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Insert exec error: %w", err)
	}
	return nil
}

func (d *DbContext) Update(destPtr interface{}, modifier func(sq.UpdateBuilder) sq.UpdateBuilder) error {
	// 1) Infer the table name
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	// 2) Start an UpdateBuilder for that table
	ub := sq.Update(table).PlaceholderFormat(sq.Dollar)

	// 3) Let the caller add .Set(...) / .Where(...) / possibly .Suffix(...) or .Returning(...)
	ub = modifier(ub)

	// 4) Build SQL + args
	sqlStr, args, err := ub.ToSql()
	if err != nil {
		return fmt.Errorf("Update: failed to build SQL: %w", err)
	}

	// 5) Exec the UPDATE
	_, err = d.DB.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Update exec error: %w", err)
	}
	return nil
}

func (d *DbContext) Delete(destPtr interface{}, modifier func(sq.DeleteBuilder) sq.DeleteBuilder) error {
	// 1) Infer the table name
	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	// 2) Start a DeleteBuilder for that table
	dbb := sq.Delete(table).PlaceholderFormat(sq.Dollar)

	// 3) Let the caller add .Where(...) / .Suffix(...) / .Returning(...)
	dbb = modifier(dbb)

	// 4) Build SQL + args
	sqlStr, args, err := dbb.ToSql()
	if err != nil {
		return fmt.Errorf("Delete: failed to build SQL: %w", err)
	}

	// 5) Exec the DELETE
	_, err = d.DB.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Delete exec error: %w", err)
	}
	return nil
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
		return "", fmt.Errorf("SelectAllPublished: expected pointer to slice, got %T", ptrToSlice)
	}
	t = t.Elem()
	if t.Kind() != reflect.Slice {
		return "", fmt.Errorf("SelectAllPublished: expected pointer to slice, got pointer to %s", t.Kind())
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
