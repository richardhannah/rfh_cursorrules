package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
	"totmapi/internal/config"
)

type User struct {
	Username string
	Password string
	Salt     string
	Role     string
}

type BlogPost struct {
	BlogPostId string
	Title      string
	Markdown   string
	Category   string
	Image      string
	Video      sql.NullString
	Date       time.Time
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

	rows, err := db.Query(fmt.Sprintf("SELECT password,salt,role FROM totm.users WHERE username = '%s'", username))
	if err != nil {
		return User{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var password string
	var salt string
	var role string

	for rows.Next() {

		err := rows.Scan(&password, &salt, &role)
		if err != nil {
			return User{}, fmt.Errorf("scan error: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return User{}, fmt.Errorf("rows error: %w", err)
	}

	return User{Username: username, Password: password, Salt: salt, Role: role}, nil
}

func SelectBlogPosts() ([]BlogPost, error) {

	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM totm.blogposts"))
	if err != nil {
		return []BlogPost{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var blogpostid string
	var title string
	var markdown string
	var category string
	var image string
	var video sql.NullString
	var date time.Time

	var posts []BlogPost

	for rows.Next() {

		err := rows.Scan(&blogpostid, &title, &markdown, &category, &image, &video, &date)
		if err != nil {
			return []BlogPost{}, fmt.Errorf("scan error: %w", err)
		}
		posts = append(posts, BlogPost{BlogPostId: blogpostid, Title: title, Markdown: markdown, Category: category, Image: image, Video: video, Date: date})
	}

	if err := rows.Err(); err != nil {
		return []BlogPost{}, fmt.Errorf("rows error: %w", err)
	}

	return posts, nil

}
