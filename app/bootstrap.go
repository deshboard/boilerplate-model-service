package app

import (
	"database/sql"
	"fmt"
)

// NewDB creates a new DB connection
func NewDB(config *Configuration) (*sql.DB, error) {
	return sql.Open("mysql", buildDSN(config))
}

func buildDSN(config *Configuration) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DbUser,
		config.DbPass,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
}
