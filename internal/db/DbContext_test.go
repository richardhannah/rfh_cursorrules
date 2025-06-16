//go:build unit

package db

import (
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestDbConnectionMockedSqlx(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testSubject := NewDbContext(db)
	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"one"}))

	rows, err := testSubject.QueryDB("SELECT 1")
	if err != nil {
		assert.Fail(t, "failed to connect to db")
	}
	assert.NotNil(t, rows)
}
