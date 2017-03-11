package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sagikazarmark/utilz/str"
)

// NewDB creates a new DB connection
func NewDB(config *Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			config.DbUser,
			config.DbPass,
			config.DbHost,
			config.DbPort,
			config.DbName,
		),
	)
	if err == nil {
		db.MapperFunc(str.ToSnake)
	}

	return db, err
}
