package di

import (
	"log"
	"reflect"
	"totmapi/internal/db"
)

var Container map[reflect.Type]interface{}

func InitializeServices() {

	RegisterService[db.DbContext](db.NewDbContext())
	RegisterService[db.BlogPostRepository](db.NewBlogPostRepository(GetService[db.DbContext]()))
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
		log.Fatal("container not initialized")
	}
	key := reflect.TypeOf((*T)(nil)).Elem()
	entry, ok := Container[key]
	if !ok {
		log.Fatalf("no service registered for type %v", key)
	}

	svcPtr, ok := entry.(*T)
	if !ok {
		log.Fatalf("registered value for %v is %T, not %T", key, entry, new(T))
	}
	return svcPtr
}
