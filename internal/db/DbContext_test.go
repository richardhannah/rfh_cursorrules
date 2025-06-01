package db

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
	"testing"
	"totmapi/internal/config"
	"totmapi/internal/models"
)

func newTestContext(t *testing.T) *DbContext {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	ctx := NewDbContext()
	t.Cleanup(func() { ctx.Close() })
	return ctx
}

func TestDbConnection(t *testing.T) {
	testSubject := newTestContext(t)
	rows := testSubject.QueryDB("SELECT 1")
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
