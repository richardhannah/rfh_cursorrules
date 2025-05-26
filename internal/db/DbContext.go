package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
	"strings"
	"totmapi/internal/config"
)

type DbContext struct {
	db *sqlx.DB
}

func NewDbContext() *DbContext {
	connStr := *config.GetDBConfig().ConnectionString
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return &DbContext{db}
}

func (d DbContext) Select(model interface{}) error {
	//var dest []models.Blogposts
	err := d.db.Select(&model, fmt.Sprintf(`SELECT * FROM  %s`, modelTypeName(model)))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d DbContext) QueryDB(query string) *sql.Rows {
	rows, err := d.db.Query(query)
	if err != nil {
		log.Printf("Failed to query Db")
		return nil
	}
	return rows
}

func (d DbContext) Close() {
	d.db.Close()
}

func modelTypeName(model interface{}) string {
	t := reflect.TypeOf(model)
	// If someone passes a pointer to a model, dereference it
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// Make sure it really is a model
	if t.Kind() != reflect.Slice {
		return ""
	}
	elem := t.Elem()
	// If it’s a named type in a package, PkgPath() != ""
	if pkg := elem.PkgPath(); pkg != "" {
		return strings.ToLower(elem.Name())
	}
	// Built-ins (e.g. int, string) or unnamed types
	if elem.Name() != "" {
		return elem.Name()
	}

	return elem.String()
}
