// +build acceptance integration

package app

import (
	"fmt"

	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	config := &testConfiguration{}

	envconfig.MustProcess("", config)

	txdb.Register(
		"txdb",
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.DbUser,
			config.DbPass,
			config.DbHost,
			config.DbPort,
			config.DbName,
		),
	)
}

// testConfiguration is a subset of main.configuration necessary for testing.
type testConfiguration struct {
	DbHost string `split_words:"true" required:"true"`
	DbPort int    `split_words:"true" default:"3306"`
	DbUser string `split_words:"true" required:"true"`
	DbPass string `split_words:"true" required:"true"`
	DbName string `split_words:"true" required:"true"`
}
