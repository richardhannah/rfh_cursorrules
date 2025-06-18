package di

import (
	"reflect"
	"totmapi/internal/config"
	"totmapi/internal/db"
	"totmapi/internal/logger"

	"github.com/jmoiron/sqlx"
)

var Container map[reflect.Type]interface{}

func InitializeServices() {
	connStr := *config.GetDBConfig().ConnectionString
	sqlx, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL", err,
			logger.String("connection_string", connStr),
		)
	}

	RegisterService[db.DbContext](db.NewDbContext(sqlx))
	RegisterService[db.BlogPostRepository](db.NewBlogPostRepository(GetService[db.DbContext]()))
	RegisterService[db.UserRepository](db.NewUserRepository(GetService[db.DbContext]()))
}

func RegisterService[T any](service *T) {
	if Container == nil {
		Container = make(map[reflect.Type]interface{})
	}
	key := reflect.TypeOf((*T)(nil)).Elem()
	Container[key] = service
}

func GetService[T any]() *T {
	if Container == nil {
		logger.Fatal("Container not initialized", &containerError{message: "container not initialized"})
	}
	key := reflect.TypeOf((*T)(nil)).Elem()
	entry, ok := Container[key]
	if !ok {
		logger.Fatal("No service registered for type", &containerError{message: "no service registered for type " + key.String()},
			logger.String("type", key.String()),
		)
	}

	svcPtr, ok := entry.(*T)
	if !ok {
		logger.Fatal("Registered value has wrong type", &containerError{message: "registered value for " + key.String() + " has wrong type"},
			logger.String("expected_type", key.String()),
			logger.String("actual_type", reflect.TypeOf(entry).String()),
		)
	}
	return svcPtr
}

// containerError implements the error interface
type containerError struct {
	message string
}

func (e *containerError) Error() string {
	return e.message
}
