package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"totmapi/internal/config"
)

func TestDbConnection(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable"
	config.SetDBConfig(&connStr)

	testSubject := NewDbContext()
	defer testSubject.Close()
	rows := testSubject.QueryDB("SELECT 1")

	assert.NotNil(t, rows)

}

func TestSelect(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	testSubject := NewDbContext()
	defer testSubject.Close()
	result := testSubject.Select()
	fmt.Println(result)

}
