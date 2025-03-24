package db

import (
	"database/sql"
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
func (b BlogPostRepository) Select() {}
func (b BlogPostRepository) Update() {}
func (b BlogPostRepository) Delete() {}
