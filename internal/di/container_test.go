//go:build integration

package di

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
	"totmapi/internal/config"
	"totmapi/internal/db"
)

func TestRegisterService(t *testing.T) {

	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	sqlx, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	ctx := db.NewDbContext(sqlx)

	RegisterService[db.DbContext](ctx)

	returnedService := GetService[db.DbContext]()
	returnedService.Close()

}

func TestInitializeServices(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	InitializeServices()

	returnedService := GetService[db.DbContext]()
	returnedService.Close()

}

func TestBlogPosts(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	InitializeServices()

	returnedService := GetService[db.BlogPostRepository]()
	result := returnedService.SelectAllPublished()
	fmt.Println(result)

}
