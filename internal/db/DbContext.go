package db

import (
	"database/sql"
	"log"
	"totmapi/internal/config"
)

type DbContext struct {
	db *sql.DB
}

func NewDbContext() *DbContext {
	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	return &DbContext{db}
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
