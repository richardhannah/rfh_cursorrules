package db

import (
	"database/sql"
	"fmt"
	"time"
)

type BlogPostRepository struct {
}

type BlogPost struct {
	BlogPostId string         `json:"blogpostid"`
	Title      string         `json:"title"`
	Markdown   string         `json:"markdown"`
	Category   string         `json:"category"`
	Image      sql.NullString `json:"image"`
	Video      sql.NullString `json:"video"`
	Date       time.Time      `json:"date"`
}

func (b BlogPostRepository) Insert() {}
func (b BlogPostRepository) Select(id string) (BlogPost, error) {

	sqlquery := fmt.Sprintf("SELECT from blogposts where id = %s", id)
	rows := Query(sqlquery)
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
			return BlogPost{}, fmt.Errorf("scan error: %w", err)
		}
		posts = append(posts, BlogPost{BlogPostId: blogpostid, Title: title, Markdown: markdown, Category: category, Image: image, Video: video, Date: date})
	}

	if err := rows.Err(); err != nil {
		return BlogPost{}, fmt.Errorf("rows error: %w", err)
	}

	if len(posts) > 1 {
		return BlogPost{}, fmt.Errorf("too many rows returned")
	}

	return posts[0], nil
}

func (b BlogPostRepository) Update() {}
func (b BlogPostRepository) Delete() {}
