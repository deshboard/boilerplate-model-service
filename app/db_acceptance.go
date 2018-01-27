// +build acceptance

package app

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/DATA-DOG/go-txdb"
	fxsql "github.com/goph/fxt/database/sql"
)

// NewDatabaseConfig returns a new database connection configuration.
func NewDatabaseConfig(config Config) *fxsql.Config {
	var registered bool
	drivers := sql.Drivers()

	for _, driver := range drivers {
		if driver == "txdb" {
			registered = true
			break
		}
	}

	if !registered {
		txdb.Register(
			"txdb",
			"mysql",
			fmt.Sprintf(
				"%s?parseTime=true",
				config.Db.Dsn(),
			),
		)
	}

	return fxsql.NewConfig(
		"txdb",
		strconv.FormatInt(time.Now().UnixNano(), 10),
	)
}
