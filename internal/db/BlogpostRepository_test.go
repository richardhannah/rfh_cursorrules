package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"totmapi/internal/config"
)

func newDbTestRepository(t *testing.T) *BlogPostRepository {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	ctx := NewDbContext()

	repository := NewBlogPostRepository(ctx)
	t.Cleanup(func() { ctx.Close() })
	return repository
}

func TestSelectAll(t *testing.T) {
	testSubject := newDbTestRepository(t)
	rows := testSubject.SelectAll()
	fmt.Println(rows)
	assert.NotNil(t, rows)
}
