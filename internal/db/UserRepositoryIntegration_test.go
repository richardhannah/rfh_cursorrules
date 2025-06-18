//go:build integration

package db

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"testing"
	"totmapi/internal/dto"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func newUserTestRepository(t *testing.T) *UserRepository {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	ctx := NewDbContext(db)

	repository := NewUserRepository(ctx)
	t.Cleanup(func() { ctx.Close() })
	return repository
}

func TestSelectByUsername(t *testing.T) {
	testSubject := newUserTestRepository(t)
	users := testSubject.SelectByUsername("richard")

	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "richard", users[0].Username)

	// Password verification: hash password from env var with the user's salt and compare
	password := os.Getenv("TOTM_TESTADMIN_PASS")
	salt := users[0].Salt
	expectedHash := hashit(password + salt)
	assert.Equal(t, expectedHash, users[0].Password)
}

// hashit replicates the password hashing logic from the application
func hashit(saltedPass string) string {
	hash := sha256.New()
	hash.Write([]byte(saltedPass))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}

func TestSelectAll(t *testing.T) {
	testSubject := newUserTestRepository(t)
	users := testSubject.SelectAll()

	assert.NotNil(t, users)
	assert.Greater(t, len(users), 0)
}

func TestInsertAndDelete(t *testing.T) {
	testSubject := newUserTestRepository(t)

	// Create a test user
	testUser := dto.UserDTO{
		ID:        "test-user-id",
		Username:  "testuser",
		Password:  "testpass",
		Salt:      "testsalt",
		Ipaddress: "127.0.0.1",
		Enabled:   true,
		Role:      "user",
	}

	// Insert the test user
	testSubject.Insert(testUser)

	// Verify the user was inserted
	users := testSubject.SelectByUsername("testuser")
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "testuser", users[0].Username)

	// Clean up by deleting the test user
	testSubject.Delete(testUser)

	// Verify the user was deleted
	users = testSubject.SelectByUsername("testuser")
	assert.Equal(t, 0, len(users))
}

func TestUserUpdate(t *testing.T) {
	testSubject := newUserTestRepository(t)

	// Create a test user
	testUser := dto.UserDTO{
		ID:        "test-update-id",
		Username:  "testupdate",
		Password:  "testpass",
		Salt:      "testsalt",
		Ipaddress: "127.0.0.1",
		Enabled:   true,
		Role:      "user",
	}

	// Insert the test user
	testSubject.Insert(testUser)

	// Update the user
	testUser.Username = "testupdate2"
	testUser.Enabled = false
	testSubject.Update(testUser)

	// Verify the update
	users := testSubject.SelectByUsername("testupdate2")
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "testupdate2", users[0].Username)
	assert.Equal(t, false, users[0].Enabled.Bool)

	// Clean up
	testSubject.Delete(testUser)
}
