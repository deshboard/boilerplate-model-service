package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/healthz"
	"github.com/goph/serverz"
	"github.com/goph/serverz/grpc"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

// newGrpcServer creates the main server instance for the service.
func newGrpcServer(appCtx *application) serverz.Server {
	serviceChecker := healthz.NewTCPChecker(appCtx.config.GrpcAddr, healthz.WithTCPTimeout(2*time.Second))
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

	return &serverz.AppServer{
		Server: &grpc.Server{Server: server},
		Name:   "grpc",
		Addr:   serverz.NewAddr("tcp", appCtx.config.GrpcAddr),
		Logger: appCtx.logger,
		Closer: db,
	}
}
