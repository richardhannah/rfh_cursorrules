//go:build integration

package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	dto2 "totmapi/internal/dto"
)

func newDbTestRepository(t *testing.T) *BlogPostRepository {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	ctx := NewDbContext(db)

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

func TestUpdateBlogPost(t *testing.T) {
	testSubject := newDbTestRepository(t)
	nowUTC := time.Now().UTC()

	dto := dto2.BlogpostDTO{
		BlogpostID: "a78fd4c1-5e30-4230-8bfd-dad27a157791",
		Title:      "testtitle",
		Markdown:   "markdown",
		Category:   "test",
		Image:      "testimage.jpg",
		Video:      "testvideo",
		Date:       &nowUTC,
		Published:  false,
	}

	testSubject.Update(dto)
}

func TestDeleteBlogPost(t *testing.T) {
	testSubject := newDbTestRepository(t)
	nowUTC := time.Now().UTC()

	dto := dto2.BlogpostDTO{
		BlogpostID: "a78fd4c1-5e30-4230-8bfd-dad27a157791",
		Title:      "testtitle",
		Markdown:   "markdown",
		Category:   "test",
		Image:      "testimage.jpg",
		Video:      "testvideo",
		Date:       &nowUTC,
		Published:  false,
	}

	testSubject.Delete(dto)
}
