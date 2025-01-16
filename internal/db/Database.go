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
	Salt     string
}

func SelectUser(username string) (User, error) {

	//connStr := buildConnectionString()
	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT password,salt FROM users WHERE username = '%s'", username))
	if err != nil {
		return User{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var password string
	var salt string
	for rows.Next() {

		err := rows.Scan(&password, &salt)
		if err != nil {
			return User{}, fmt.Errorf("scan error: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return User{}, fmt.Errorf("rows error: %w", err)
	}

	return User{Username: username, Password: password, Salt: salt}, nil
}
