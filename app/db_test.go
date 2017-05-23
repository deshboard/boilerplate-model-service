// +build acceptance integration

package app

import (
	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	config := &Configuration{}

	envconfig.MustProcess("", config)

	txdb.Register("txdb", "mysql", buildDSN(config))
}
