package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sagikazarmark/utilz/strings"
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
		db.MapperFunc(strings.ToSnake)
	}

	return db, err
}
