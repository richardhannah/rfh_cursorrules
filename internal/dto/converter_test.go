//go:build unit

package dto

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
	"totmapi/internal/models"
)

func TestBlogPostConversion(t *testing.T) {

	blogpost := models.Blogposts{
		BlogpostID: "test",
		Title:      sql.NullString{String: "test", Valid: true},
		Markdown:   sql.NullString{String: "test", Valid: true},
		Category:   sql.NullString{String: "test", Valid: true},
		Image:      sql.NullString{String: "test", Valid: true},
		Video:      sql.NullString{String: "test", Valid: true},
		Date:       sql.NullTime{Time: time.Time{}, Valid: true},
		Published:  false,
	}

	result, err := ConvertToDTO[models.Blogposts, BlogpostDTO](blogpost)
	if err != nil {
		t.Errorf("Error converting models.Blogposts to BlogpostDTO: %v", err)
	}

	fmt.Println(result)

}

func TestBlogPostSliceConversion(t *testing.T) {

	blogpost := models.Blogposts{
		BlogpostID: "test",
		Title:      sql.NullString{String: "test", Valid: true},
		Markdown:   sql.NullString{String: "test", Valid: true},
		Category:   sql.NullString{String: "test", Valid: true},
		Image:      sql.NullString{String: "test", Valid: true},
		Video:      sql.NullString{String: "test", Valid: true},
		Date:       sql.NullTime{Time: time.Time{}, Valid: true},
		Published:  false,
	}

	blogPostsSlice := make([]models.Blogposts, 0)
	blogPostsSlice = append(blogPostsSlice, blogpost)

	result, err := ConvertSlice[models.Blogposts, BlogpostDTO](blogPostsSlice)
	if err != nil {
		t.Errorf("Error converting models.Blogposts to BlogpostDTO: %v", err)
	}

	fmt.Println(result)

}
