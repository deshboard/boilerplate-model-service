// +build acceptance integration

package app_test

import (
	"fmt"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/stdlib/os"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load("../.env.test", "../.env.dist")

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