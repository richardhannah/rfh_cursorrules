package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"totmapi/internal/config"
	"totmapi/internal/models"
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

func (d DbContext) Select() string {
	var posts []models.Blogposts
	// this will do: rows.Columns(), map col→field by `db` tag, scan for you
	err := d.db.Select(&posts, `SELECT blogpostid, title, markdown, category, image, video, date, published
                             FROM blogposts`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf(posts[0].Title.String)
	return posts[1].Title.String
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
