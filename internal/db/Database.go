package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"sort"
	"time"
	"totmapi/internal/config"
)

type User struct {
	Username string
	Password string
	Salt     string
	Role     string
}

//func (dbc *DBContext) Query[T any](sqlQuery string) T {
//	var result T
//	return result
//}

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

func UpdatePassword(username string, newpassword string) error {

	// Get the PostgreSQL connection string from your config
	connStr := *config.GetDBConfig().ConnectionString

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Build the UPDATE query
	// We only need two parameters: new password and username
	query := `
        UPDATE totm.users
        SET password = $1
        WHERE username = $2
    `

	// Execute the update with the new password and the username
	_, err = db.Exec(query, newpassword, username)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("UpdatePassword: %v", err)
	}

	return nil
}

func SelectBlogPosts(publishedParam string) ([]BlogPost, error) {

	publishedParamConstraint := "WHERE published = true"

	if publishedParam == "all" {
		publishedParamConstraint = ""
	}

	if publishedParam == "false" {
		publishedParamConstraint = "WHERE published = false"
	}

	connStr := *config.GetDBConfig().ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM totm.blogposts %s", publishedParamConstraint))
	if err != nil {
		return []BlogPost{}, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var blogpostid string
	var title string
	var markdown string
	var category string
	var image sql.NullString
	var video sql.NullString
	var date time.Time
	var published bool

	var posts []BlogPost

	for rows.Next() {

		err := rows.Scan(&blogpostid, &title, &markdown, &category, &image, &video, &date, &published)
		if err != nil {
			return []BlogPost{}, fmt.Errorf("scan error: %w", err)
		}
		posts = append(posts, BlogPost{BlogPostId: blogpostid, Title: title, Markdown: markdown, Category: category, Image: image, Video: video, Date: date})
	}

	if err := rows.Err(); err != nil {
		return []BlogPost{}, fmt.Errorf("rows error: %w", err)
	}

	sortPostsDescending(posts)

	return posts, nil

}

func sortPostsDescending(posts []BlogPost) {
	sort.Slice(posts, func(i, j int) bool {
		// Return true if posts[i] should appear before posts[j].
		// For newest first, we want the more recent date to come first.
		// So compare using Date.After(...).
		return posts[i].Date.After(posts[j].Date)
	})
}
