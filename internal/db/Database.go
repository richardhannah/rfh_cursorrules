package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"totmapi/internal/config"
)

type User struct {
	Username string
	Password string
	Salt     string
}

func InsertUser(username string, password string, salt string) error {

	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Example SQL statement using placeholders
	query := `
        INSERT INTO totm.users (id,username, password, salt)
        VALUES ($1, $2, $3, $4)
    `

	id := uuid.New().String()
	// Execute the INSERT
	_, err = db.Exec(query, id, username, password, salt)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("insertUser: %v", err)
	}

	return nil

}

func SelectUser(username string) (User, error) {

	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT password,salt FROM totm.users WHERE username = '%s'", username))
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
