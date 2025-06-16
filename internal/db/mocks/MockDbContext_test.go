package mocks

import (
	"database/sql"
	"fmt"
	"testing"
	"totmapi/internal/models"
)

func TestExampleUsage(t *testing.T) {

	testSubject := NewMockDbContext()
	testSubject.SelectAllFunc = func(destPtr interface{}) error {
		slicePtr := destPtr.(*[]models.Blogposts)
		*slicePtr = []models.Blogposts{
			{BlogpostID: "id1", Title: sql.NullString{String: "T1", Valid: true}},
			{BlogpostID: "id2", Title: sql.NullString{String: "T2", Valid: true}},
		}
		return nil
	}

	var posts []models.Blogposts
	if err := testSubject.SelectAll(&posts); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Loaded %d posts\n", len(posts))
}
