// +build integration

package app_test

import (
	"os"
	"testing"

	"github.com/deshboard/boilerplate-model-service/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	integrationSetUp()

	result := m.Run()

	integrationTearDown()

	os.Exit(result)
}

// Integration test initialization
func integrationSetUp() {
	setupDatabase()
}

// Cleanup after integration tests
func integrationTearDown() {
	teardownDatabase()
}

func setupDatabase() {
	config := &app.Configuration{}

	envconfig.MustProcess("app", config)

	// Necessary to avoid shadowing global "db" variable
	var err error

	db, err = app.NewDB(config)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	cleanupDatabase()
}

func teardownDatabase() {
	cleanupDatabase()

	db.Close()
}

func cleanupDatabase() {
	db.MustExec("DELETE FROM boilerplate")
}
