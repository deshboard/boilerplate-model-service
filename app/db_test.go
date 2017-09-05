// +build acceptance integration

package model_test

import (
	"testing"

	"database/sql"
	"fmt"
	"os"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load("../.env.test", "../.env.dist")

	txdb.Register(
		"txdb",
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
	)
}

func TestDbConnection(t *testing.T) {
	db, err := sql.Open("txdb", "db connection check")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
