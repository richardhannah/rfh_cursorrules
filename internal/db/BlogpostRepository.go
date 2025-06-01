package db

import (
	"database/sql"
	"fmt"
	"log"
	"totmapi/internal/models"
)

type BlogPostRepository struct {
	dbContext IDbContext
}

func NewBlogPostRepository(ctx IDbContext) *BlogPostRepository {
	return &BlogPostRepository{dbContext: ctx}
}

func (b BlogPostRepository) SelectAll() []models.Blogposts {

	var posts []models.Blogposts
	if err := b.dbContext.Select(&posts); err != nil {
		log.Println(fmt.Sprintf("Error selecting all blogposts %s", err))
	}
	return posts
}

func (b BlogPostRepository) Insert() {}
func (b BlogPostRepository) Select(id string) (models.Blogposts, error) {

	sqlquery := fmt.Sprintf("SELECT from blogposts where id = %s", id)
	rows := Query(sqlquery)
	defer rows.Close()

	var blogpostid string
	var title sql.NullString
	var markdown sql.NullString
	var category sql.NullString
	var image sql.NullString
	var video sql.NullString
	var date sql.NullTime
	var published bool

	var posts []models.Blogposts

	for rows.Next() {

		err := rows.Scan(&blogpostid, &title, &markdown, &category, &image, &video, &date, &published)
		if err != nil {
			return models.Blogposts{}, fmt.Errorf("scan error: %w", err)
		}
		posts = append(posts, models.Blogposts{BlogpostID: blogpostid, Title: title, Markdown: markdown, Category: category, Image: image, Video: video, Date: date})
	}

	if err := rows.Err(); err != nil {
		return models.Blogposts{}, fmt.Errorf("rows error: %w", err)
	}

	if len(posts) > 1 {
		return models.Blogposts{}, fmt.Errorf("too many rows returned")
	}

	return posts[0], nil
}

func (b BlogPostRepository) Update() {}
func (b BlogPostRepository) Delete() {}
