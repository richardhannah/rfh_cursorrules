package db

import (
	"totmapi/internal/dto"
	"totmapi/internal/logger"
	"totmapi/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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
		logger.Error("Error selecting all published blogposts", err)
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
		logger.Error("Error selecting blogpost by id", err,
			logger.String("blogpost_id", id),
		)
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

func (b BlogPostRepository) Update(blogPost dto.BlogpostDTO) {

	predicate := func(ub sq.UpdateBuilder) sq.UpdateBuilder {
		return ub.
			Set("title", "updated Title").
			Where(
				sq.Eq{"blogpostid": blogPost.BlogpostID},
			)
	}

	b.dbContext.Update(&[]models.Blogposts{}, predicate)

}

func (b BlogPostRepository) Delete(blogPost dto.BlogpostDTO) {

	predicate := func(db sq.DeleteBuilder) sq.DeleteBuilder {
		return db.Where(
			sq.Eq{"blogpostid": blogPost.BlogpostID},
		)
	}

	err := b.dbContext.Delete(&[]models.Blogposts{}, predicate)
	if err != nil {
		logger.Error("Error deleting blogpost", err,
			logger.String("blogpost_id", blogPost.BlogpostID),
		)
	}

}
