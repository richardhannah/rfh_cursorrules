package mocks

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type MockDbContext struct {
	SelectAllFunc func(destPtr interface{}) error
	SelectFunc    func(destPtr interface{}, modifier func(sq.SelectBuilder) sq.SelectBuilder) error
	InsertFunc    func(destPtr interface{}, modifier func(sq.InsertBuilder) sq.InsertBuilder) error
	UpdateFunc    func(destPtr interface{}, modifier func(sq.UpdateBuilder) sq.UpdateBuilder) error
	DeleteFunc    func(destPtr interface{}, modifier func(sq.DeleteBuilder) sq.DeleteBuilder) error
	QueryDBFunc   func(query string) (*sql.Rows, error)
	CloseFunc     func() error
}

func NewMockDbContext() *MockDbContext {
	return &MockDbContext{}
}

func (mdb MockDbContext) SelectAll(destPtr interface{}) error {
	if mdb.SelectAllFunc != nil {
		return mdb.SelectAllFunc(destPtr)
	}
	return nil
}
func (mdb MockDbContext) Select(destPtr interface{}, modifier func(sq.SelectBuilder) sq.SelectBuilder) error {
	if mdb.SelectFunc != nil {
		return mdb.SelectFunc(destPtr, modifier)
	}
	return nil
}
func (mdb MockDbContext) Insert(destPtr interface{}, modifier func(sq.InsertBuilder) sq.InsertBuilder) error {
	if mdb.InsertFunc != nil {
		return mdb.InsertFunc(destPtr, modifier)
	}
	return nil
}
func (mdb MockDbContext) Update(destPtr interface{}, modifier func(sq.UpdateBuilder) sq.UpdateBuilder) error {
	if mdb.UpdateFunc != nil {
		return mdb.UpdateFunc(destPtr, modifier)
	}
	return nil
}
func (mdb MockDbContext) Delete(destPtr interface{}, modifier func(sq.DeleteBuilder) sq.DeleteBuilder) error {
	if mdb.DeleteFunc != nil {
		return mdb.DeleteFunc(destPtr, modifier)
	}
	return nil
}

func (mdb MockDbContext) QueryDB(query string) (*sql.Rows, error) {
	if mdb.QueryDBFunc != nil {
		return mdb.QueryDBFunc(query)
	}
	return nil, nil
}
func (mdb MockDbContext) Close() error {
	if mdb.CloseFunc != nil {
		return mdb.CloseFunc()
	}
	return nil
}
