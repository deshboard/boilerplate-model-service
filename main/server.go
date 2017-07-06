package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/healthz"
	"github.com/goph/serverz"
	"github.com/goph/serverz/aio"
	_grpc "github.com/goph/serverz/grpc"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

// newServer creates the main server instance for the service.
func newServer(appCtx *application) serverz.Server {
	serviceChecker := healthz.NewTCPChecker(appCtx.config.ServiceAddr, healthz.WithTCPTimeout(2*time.Second))
	appCtx.healthCollector.RegisterChecker(healthz.LivenessCheck, serviceChecker)

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			appCtx.config.DbUser,
			appCtx.config.DbPass,
			appCtx.config.DbHost,
			appCtx.config.DbPort,
			appCtx.config.DbName,
		),
	)

	if err != nil {
		panic(err)
	}

	appCtx.healthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))

	server := createGrpcServer(appCtx)

	// Register servers here

	grpc_prometheus.Register(server)

	return &aio.Server{
		Server:     &_grpc.Server{Server: server},
		ServerName: "grpc",
		Closer:     db,
	}
}
