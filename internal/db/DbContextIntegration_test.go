//go:build integration

package db

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	"totmapi/internal/models"
)

func newTestContext(t *testing.T) *DbContext {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	sqlxDb, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	ctx := NewDbContext(sqlxDb)
	t.Cleanup(func() { ctx.Close() })
	return ctx
}

func TestDbConnection(t *testing.T) {
	testSubject := newTestContext(t)

	rows, err := testSubject.QueryDB("SELECT 1")
	if err != nil {
		assert.Fail(t, "failed to connect to db")
	}
	assert.NotNil(t, rows)
}

func TestSelectBlogPosts(t *testing.T) {

	testSubject := newTestContext(t)

	var posts []models.Blogposts
	if err := testSubject.SelectAll(&posts); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Loaded %d posts\n", len(posts))
}

func TestSelectUsers(t *testing.T) {

	testSubject := newTestContext(t)

	var posts []models.Users
	if err := testSubject.SelectAll(&posts); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Loaded %d users\n", len(posts))
}

func TestSelectUsersWithPredicate(t *testing.T) {

	testSubject := newTestContext(t)

	var users []models.Users

	predicate := func(sb sq.SelectBuilder) sq.SelectBuilder {
		return sb.
			Where(sq.Eq{"enabled": true})
	}

	if err := testSubject.Select(&users, predicate); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Loaded %d users\n", len(users))
}

func TestInsert(t *testing.T) {

	testSubject := newTestContext(t)

	newPost := models.Blogposts{
		BlogpostID: "some-uuid",
		Title:      sql.NullString{String: "New Title", Valid: true},
		Markdown:   sql.NullString{String: "Markdown body", Valid: true},
		Category:   sql.NullString{String: "test", Valid: true},
		Image:      sql.NullString{Valid: false}, // will insert NULL
		Video:      sql.NullString{String: "xyz123", Valid: true},
		Date:       sql.NullTime{Time: time.Now(), Valid: true},
		Published:  false,
	}

	newPosts := make([]models.Blogposts, 0)
	newPosts = append(newPosts, newPost)

	predicate := func(ib sq.InsertBuilder) sq.InsertBuilder {
		return ib.
			Columns(
				"blogpostid",
				"title",
				"markdown",
				"category",
				"image",
				"video",
				"date",
				"published",
			).
			Values(
				newPost.BlogpostID,
				newPost.Title,
				newPost.Markdown,
				newPost.Category,
				newPost.Image,
				newPost.Video,
				newPost.Date,
				newPost.Published,
			)
	}

	err := testSubject.Insert(&newPosts, predicate)
	if err != nil {
		fmt.Println(err)
	}
}

func TestUpdate(t *testing.T) {

	testSubject := newTestContext(t)

	newPost := models.Blogposts{
		BlogpostID: "some-uuid",
		Title:      sql.NullString{String: "New Title", Valid: true},
		Markdown:   sql.NullString{String: "Markdown body", Valid: true},
		Category:   sql.NullString{String: "test", Valid: true},
		Image:      sql.NullString{Valid: false}, // will insert NULL
		Video:      sql.NullString{String: "xyz123", Valid: true},
		Date:       sql.NullTime{Time: time.Now(), Valid: true},
		Published:  false,
	}

	newPosts := make([]models.Blogposts, 0)
	newPosts = append(newPosts, newPost)

	predicate := func(ub sq.UpdateBuilder) sq.UpdateBuilder {
		return ub.
			Set("title", "updated Title").
			Where(
				sq.Eq{"blogpostid": "some-uuid"},
			)
	}

	err := testSubject.Update(&newPosts, predicate)
	if err != nil {
		fmt.Println(err)
	}
}

func TestDelete(t *testing.T) {

	testSubject := newTestContext(t)

	newPost := models.Blogposts{
		BlogpostID: "some-uuid",
		Title:      sql.NullString{String: "New Title", Valid: true},
		Markdown:   sql.NullString{String: "Markdown body", Valid: true},
		Category:   sql.NullString{String: "test", Valid: true},
		Image:      sql.NullString{Valid: false}, // will insert NULL
		Video:      sql.NullString{String: "xyz123", Valid: true},
		Date:       sql.NullTime{Time: time.Now(), Valid: true},
		Published:  false,
	}

	newPosts := make([]models.Blogposts, 0)
	newPosts = append(newPosts, newPost)

	predicate := func(db sq.DeleteBuilder) sq.DeleteBuilder {
		return db.Where(
			sq.Eq{"blogpostid": "some-uuid"},
		)
	}

	err := testSubject.Delete(&newPosts, predicate)
	if err != nil {
		fmt.Println(err)
	}
}
