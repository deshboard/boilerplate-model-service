// +build integration

package app

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

var DB *sqlx.DB

func TestMain(m *testing.M) {
	config := &Configuration{}

	envconfig.MustProcess("", config)

	setUp(config)

	result := m.Run()

	tearDown(config)

	os.Exit(result)
}

// Integration test initialization
func setUp(config *Configuration) {
	// Necessary to avoid shadowing global "DB" variable
	var err error

	DB, err = NewDB(config)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	cleanupDatabase()
}

// Cleanup after integration tests
func tearDown(config *Configuration) {
	cleanupDatabase()

	DB.Close()
}

func cleanupDatabase() {
	DB.MustExec("DELETE FROM boilerplate")
}
