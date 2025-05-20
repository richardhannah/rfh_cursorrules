package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"totmapi/internal/config"
)

func TestSelectUser(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable"

	config.SetDBConfig(&connStr)

	user, err := SelectUser("richard")
	if err != nil {
		log.Fatal("error", err)
	}
	assert.Equal(t, "richard", user.Username)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
}
