package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type ISqlx interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

type DbContext struct {
	DB ISqlx
}

func NewDbContext(iSqlx ISqlx) *DbContext {
	//connStr := *config.GetDBConfig().ConnectionString
	//db, err := sqlx.Connect("postgres", connStr)
	//if err != nil {
	//	log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	//}
	return &DbContext{DB: iSqlx}
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

	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	ib := sq.Insert(table).PlaceholderFormat(sq.Dollar)

	ib = modifier(ib)

	sqlStr, args, err := ib.ToSql()
	if err != nil {
		return fmt.Errorf("Insert: failed to build SQL: %w", err)
	}

	_, err = d.DB.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Insert exec error: %w", err)
	}
	return nil
}

func (d *DbContext) Update(destPtr interface{}, modifier func(sq.UpdateBuilder) sq.UpdateBuilder) error {

	table, err := tableName(destPtr)
	if err != nil {
		return err
	}

	ub := sq.Update(table).PlaceholderFormat(sq.Dollar)

	ub = modifier(ub)

	sqlStr, args, err := ub.ToSql()
	if err != nil {
		return fmt.Errorf("Update: failed to build SQL: %w", err)
	}

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

	dbb := sq.Delete(table).PlaceholderFormat(sq.Dollar)

	dbb = modifier(dbb)

	sqlStr, args, err := dbb.ToSql()
	if err != nil {
		return fmt.Errorf("Delete: failed to build SQL: %w", err)
	}

	_, err = d.DB.Exec(sqlStr, args...)
	if err != nil {
		return fmt.Errorf("Delete exec error: %w", err)
	}
	return nil
}

func (d DbContext) QueryDB(query string) (*sql.Rows, error) {
	rows, err := d.DB.Query(query)
	if err != nil {
		log.Printf("Failed to query Db")
		return nil, err
	}
	return rows, nil
}

func (d DbContext) Close() error {
	err := d.DB.Close()
	if err != nil {
		return err
	}
	return nil
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
