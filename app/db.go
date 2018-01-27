// +build !acceptance

package app

import (
	"fmt"

	fxsql "github.com/goph/fxt/database/sql"
)

// NewDatabaseConfig returns a new database connection configuration.
func NewDatabaseConfig(config Config) *fxsql.Config {
	return fxsql.NewConfig(
		"mysql",
		fmt.Sprintf(
			"%s?parseTime=true",
			config.Db.Dsn(),
		),
	)
}
