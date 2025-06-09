package mocks

import (
	"database/sql"
)

type MockSqlx struct {
	SelectFunc func(dest interface{}, query string, args ...interface{}) error
	ExecFunc   func(query string, args ...interface{}) (sql.Result, error)
	QueryFunc  func(query string, args ...interface{}) (*sql.Rows, error)
	CloseFunc  func() error
}

func NewMockSqlx() *MockSqlx {
	return &MockSqlx{}
}

func (msqlx MockSqlx) Select(dest interface{}, query string, args ...interface{}) error {
	if msqlx.SelectFunc != nil {
		return msqlx.SelectFunc(dest, query, args)
	}
	return nil
}
func (msqlx MockSqlx) Exec(query string, args ...interface{}) (sql.Result, error) {
	if msqlx.ExecFunc != nil {
		return msqlx.ExecFunc(query, args)
	}
	return nil, nil
}
func (msqlx MockSqlx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if msqlx.QueryFunc != nil {
		return msqlx.QueryFunc(query, args)
	}
	return nil, nil
}
func (msqlx MockSqlx) Close() error {
	if msqlx.CloseFunc != nil {
		return msqlx.CloseFunc()
	}
	return nil
}
