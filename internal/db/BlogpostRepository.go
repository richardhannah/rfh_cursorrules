package db

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"log"
	"totmapi/internal/dto"
	"totmapi/internal/models"
)

type BlogPostRepository struct {
	dbContext IDbContext
}

func NewBlogPostRepository(ctx IDbContext) *BlogPostRepository {
	return &BlogPostRepository{dbContext: ctx}
}

func (b BlogPostRepository) SelectAllPublished() []models.Blogposts {

	var posts []models.Blogposts

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.
			Where(sq.Eq{"published": true}).
			OrderBy("date DESC")
	}

	if err := b.dbContext.Select(&posts, predicate); err != nil {
		log.Println(fmt.Sprintf("Error selecting all blogposts %s", err))
	}
	return posts
}

func (b BlogPostRepository) SelectById(id string) []models.Blogposts {
	var posts []models.Blogposts

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.
			Where(sq.Eq{"blogpostid": id})
	}

	if err := b.dbContext.Select(&posts, predicate); err != nil {
		log.Println(fmt.Sprintf("Error selecting all blogposts %s", err))
	}
	return posts
}

func (b BlogPostRepository) Insert(blogPost dto.BlogpostDTO) {

	predicate := func(sb sq.InsertBuilder) sq.InsertBuilder {
		return sb.
			Columns("blogpostid", "title", "markdown", "category", "image", "video", "date", "published").
			Values(uuid.New().String(), blogPost.Title, blogPost.Markdown, blogPost.Category, blogPost.Image, blogPost.Video, blogPost.Date, blogPost.Published)
	}

	b.dbContext.Insert(&[]models.Blogposts{}, predicate)
}

func (b BlogPostRepository) Update() {}
func (b BlogPostRepository) Delete() {}

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
