package di

import (
	"fmt"
	"testing"
	"totmapi/internal/config"
	"totmapi/internal/db"
)

func TestRegisterService(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)
	ctx := db.NewDbContext()

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
	result := returnedService.SelectAll()
	fmt.Println(result)

}
