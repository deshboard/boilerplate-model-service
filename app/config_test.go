package app

import (
	"fmt"
	"os"
	"strconv"

	"github.com/goph/fxt/database/sql"
	"github.com/goph/fxt/test/nettest"
)

func newConfig() Config {
	debugPort, _ := nettest.GetFreePort()
	grpcAddr, _ := nettest.GetFreePort()
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return Config{
		Environment: "test",
		LogFormat:   "logfmt",
		DebugAddr:   fmt.Sprintf("127.0.0.1:%d", debugPort),
		GrpcAddr:    fmt.Sprintf("127.0.0.1:%d", grpcAddr),
		Db: sql.AppConfig{
			Host: os.Getenv("DB_HOST"),
			Port: dbPort,
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
