//go:build unit

package db

import (
	"testing"
	"totmapi/internal/db/mocks"
	"totmapi/internal/dto"
	"totmapi/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_SelectAll(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.SelectFunc = func(result interface{}, predicate func(sq.SelectBuilder) sq.SelectBuilder) error {
		if users, ok := result.(*[]models.Users); ok {
			*users = append(*users, models.Users{ID: "1", Username: "user1"})
		}
		return nil
	}
	repo := NewUserRepository(mockCtx)

	users := repo.SelectAll()

	assert.NotNil(t, users)
}

func TestUserRepository_SelectById(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.SelectFunc = func(result interface{}, predicate func(sq.SelectBuilder) sq.SelectBuilder) error {
		if users, ok := result.(*[]models.Users); ok {
			*users = append(*users, models.Users{ID: "1", Username: "user1"})
		}
		return nil
	}
	repo := NewUserRepository(mockCtx)

	users := repo.SelectById("1")

	assert.NotNil(t, users)
}

func TestUserRepository_SelectByUsername(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.SelectFunc = func(result interface{}, predicate func(sq.SelectBuilder) sq.SelectBuilder) error {
		if users, ok := result.(*[]models.Users); ok {
			*users = append(*users, models.Users{ID: "1", Username: "user1"})
		}
		return nil
	}
	repo := NewUserRepository(mockCtx)

	users := repo.SelectByUsername("user1")

	assert.NotNil(t, users)
}

func TestUserRepository_Insert(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.InsertFunc = func(result interface{}, predicate func(sq.InsertBuilder) sq.InsertBuilder) error {
		return nil
	}
	repo := NewUserRepository(mockCtx)

	testUser := dto.UserDTO{
		ID:        "test-id",
		Username:  "testuser",
		Password:  "testpass",
		Salt:      "testsalt",
		Ipaddress: "127.0.0.1",
		Enabled:   true,
		Role:      "user",
	}

	repo.Insert(testUser)
}

func TestUserRepository_Update(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.UpdateFunc = func(result interface{}, predicate func(sq.UpdateBuilder) sq.UpdateBuilder) error {
		return nil
	}
	repo := NewUserRepository(mockCtx)

	testUser := dto.UserDTO{
		ID:        "test-id",
		Username:  "testuser",
		Password:  "testpass",
		Salt:      "testsalt",
		Ipaddress: "127.0.0.1",
		Enabled:   true,
		Role:      "user",
	}

	repo.Update(testUser)
}

func TestUserRepository_Delete(t *testing.T) {
	mockCtx := mocks.NewMockDbContext()
	mockCtx.DeleteFunc = func(result interface{}, predicate func(sq.DeleteBuilder) sq.DeleteBuilder) error {
		return nil
	}
	repo := NewUserRepository(mockCtx)

	testUser := dto.UserDTO{
		ID:        "test-id",
		Username:  "testuser",
		Password:  "testpass",
		Salt:      "testsalt",
		Ipaddress: "127.0.0.1",
		Enabled:   true,
		Role:      "user",
	}

	repo.Delete(testUser)
}
