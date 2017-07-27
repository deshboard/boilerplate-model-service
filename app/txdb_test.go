// +build acceptance integration

package app

import (
	"fmt"

	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/stdlib/os"
)

func init() {
	txdb.Register(
		"txdb",
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.MustEnv("DB_USER"),
			os.MustEnv("DB_PASS"),
			os.MustEnv("DB_HOST"),
			os.MustEnv("DB_PORT"),
			os.MustEnv("DB_NAME"),
		),
	)
}
