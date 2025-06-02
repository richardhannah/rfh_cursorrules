package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"totmapi/internal/config"
	dto2 "totmapi/internal/dto"
)

func newDbTestRepository(t *testing.T) *BlogPostRepository {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	ctx := NewDbContext()

	repository := NewBlogPostRepository(ctx)
	t.Cleanup(func() { ctx.Close() })
	return repository
}

func TestSelectAllPublished(t *testing.T) {
	testSubject := newDbTestRepository(t)
	rows := testSubject.SelectAllPublished()
	fmt.Println(rows)
	assert.NotNil(t, rows)
}

func TestSelectById(t *testing.T) {
	testSubject := newDbTestRepository(t)
	rows := testSubject.SelectById("471b7545-56d4-4f99-9914-c284c9dfe4ba")
	fmt.Println(rows)
	assert.NotNil(t, rows)
}

func TestInsertBlogPost(t *testing.T) {
	testSubject := newDbTestRepository(t)
	nowUTC := time.Now().UTC()

	dto := dto2.BlogpostDTO{
		BlogpostID: "",
		Title:      "testtitle",
		Markdown:   "markdown",
		Category:   "test",
		Image:      "testimage.jpg",
		Video:      "testvideo",
		Date:       &nowUTC,
		Published:  false,
	}

	testSubject.Insert(dto)
}
