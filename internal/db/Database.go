package db

import (
	"database/sql"
	"fmt"
	"log"
	"totmapi/internal/config"
)

type User struct {
	Username string
	Password string
}

func SelectUser(username string) (User, error) {

	connStr := buildConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT password FROM users WHERE username = '%s'", username))
	if err != nil {
		return User{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var password string
	for rows.Next() {

		err := rows.Scan(&password)
		if err != nil {
			return User{}, fmt.Errorf("scan error: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return User{}, fmt.Errorf("rows error: %w", err)
	}

	return User{Username: username, Password: password}, nil
}

func buildConnectionString() string {
	conf := config.GetDBConfig()

	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", conf.Username, conf.Password, conf.Host, conf.Database)
}
