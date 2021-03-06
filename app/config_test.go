package app

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/goph/fxt/database/sql"
	"github.com/goph/fxt/dev"
	"github.com/goph/fxt/log"
	"github.com/goph/fxt/testing/nettest"
	"github.com/goph/nest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	fxdev.LoadEnvFromFile("../.env.test")
	fxdev.LoadEnvFromFile("../.env.dist")
}

func newConfig() (Config, error) {
	debugPort, _ := nettest.GetFreePort()
	grpcAddr, _ := nettest.GetFreePort()

	config := Config{
		Environment: "test",
		LogFormat:   "logfmt",
		DebugAddr:   fmt.Sprintf("127.0.0.1:%d", debugPort),
		GrpcAddr:    fmt.Sprintf("127.0.0.1:%d", grpcAddr),
	}

	configurator := nest.NewConfigurator()
	configurator.SetName(FriendlyServiceName)
	configurator.SetArgs([]string{})

	err := configurator.Load(&config)

	return config, err
}

func TestConfig(t *testing.T) {
	defer func() {
		os.Clearenv()
		fxdev.LoadEnvFromFile("../.env.test")
		fxdev.LoadEnvFromFile("../.env.dist")
	}()

	tests := map[string]struct {
		env      map[string]string
		args     []string
		actual   Config
		expected Config
	}{
		"full config": {
			map[string]string{
				"ENVIRONMENT":            "test",
				"DEBUG":                  "false",
				"LOG_FORMAT":             "logfmt",
				"GRPC_ENABLE_REFLECTION": "true",
				"DB_HOST":                "127.0.0.1",
				"DB_PORT":                "3336",
				"DB_USER":                "root",
				"DB_PASS":                "toor",
				"DB_NAME":                "service",
			},
			[]string{"service", "--debug-addr", ":10001", "--shutdown-timeout", "10s", "--grpc-addr", ":8001"},
			Config{},
			Config{
				Environment:          "test",
				Debug:                false,
				LogFormat:            fxlog.LogfmtFormat.String(),
				DebugAddr:            ":10001",
				ShutdownTimeout:      10 * time.Second,
				GrpcAddr:             ":8001",
				GrpcEnableReflection: true,
				Db: fxsql.AppConfig{
					Host: "127.0.0.1",
					Port: 3336,
					User: "root",
					Pass: "toor",
					Name: "service",
				},
			},
		},
		"defaults": {
			map[string]string{},
			[]string{},
			Config{
				Db: fxsql.AppConfig{
					Host: "localhost",
					User: "root",
					Name: "db",
				},
			},
			Config{
				Environment:          "production",
				Debug:                false,
				LogFormat:            fxlog.JsonFormat.String(),
				DebugAddr:            ":10000",
				ShutdownTimeout:      15 * time.Second,
				GrpcAddr:             ":8000",
				GrpcEnableReflection: false,
				Db: fxsql.AppConfig{
					Host: "localhost",
					Port: 3306,
					User: "root",
					Pass: "",
					Name: "db",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			os.Clearenv()

			for key, value := range test.env {
				os.Setenv(key, value)
			}

			configurator := nest.NewConfigurator()
			configurator.SetName(FriendlyServiceName)
			configurator.SetArgs(test.args)

			err := configurator.Load(&test.actual)
			require.NoError(t, err)
			assert.Equal(t, test.expected, test.actual)
		})
	}
}
