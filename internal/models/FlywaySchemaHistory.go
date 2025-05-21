package models

import (
	"database/sql"
	"time"
)

type FlywaySchemaHistory struct {
	InstalledRank int            `db:"installed_rank"`
	Version       sql.NullString `db:"version"`
	Description   string         `db:"description"`
	Type          string         `db:"type"`
	Script        string         `db:"script"`
	Checksum      sql.NullInt64  `db:"checksum"`
	InstalledBy   string         `db:"installed_by"`
	InstalledOn   time.Time      `db:"installed_on"`
	ExecutionTime int            `db:"execution_time"`
	Success       bool           `db:"success"`
}
