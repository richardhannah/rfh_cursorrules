package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"totmapi/internal/config"
	"totmapi/internal/models"
)

func TestDbConnection(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable"
	config.SetDBConfig(&connStr)

	testSubject := NewDbContext()
	defer testSubject.Close()
	rows := testSubject.QueryDB("SELECT 1")

	assert.NotNil(t, rows)

}

func TestSelect(t *testing.T) {
	connStr := "postgres://richard:Onlyone1@localhost:5432/richard?sslmode=disable&search_path=totm"
	config.SetDBConfig(&connStr)

	testSubject := NewDbContext()
	defer testSubject.Close()
	var dest []models.Blogposts
	err := testSubject.Select(dest)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(dest)

}

func TestWorkBench(t *testing.T) {

	var dest []models.Blogposts

	fmt.Println(elementTypeName(dest))
}
func elementTypeName(slice interface{}) string {
	t := reflect.TypeOf(slice)
	// If someone passes a pointer to a slice, dereference it
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// Make sure it really is a slice
	if t.Kind() != reflect.Slice {
		return ""
	}
	elem := t.Elem()
	// If it’s a named type in a package, PkgPath() != ""
	if pkg := elem.PkgPath(); pkg != "" {
		return strings.ToLower(elem.Name())
	}
	// Built-ins (e.g. int, string) or unnamed types
	if elem.Name() != "" {
		return elem.Name()
	}

	return elem.String()
}
