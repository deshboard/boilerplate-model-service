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
func newGrpcServer(app *application) serverz.Server {
	serviceChecker := healthz.NewTCPChecker(app.config.GrpcAddr, healthz.WithTCPTimeout(2*time.Second))
	app.healthCollector.RegisterChecker(healthz.LivenessCheck, serviceChecker)

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			app.config.DbUser,
			app.config.DbPass,
			app.config.DbHost,
			app.config.DbPort,
			app.config.DbName,
		),
	)

	if err != nil {
		panic(err)
	}

	app.healthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))

	server := createGrpcServer(app)

	// Register servers here

	grpc_prometheus.Register(server)

	return &serverz.AppServer{
		Server: &grpc.Server{Server: server},
		Name:   "grpc",
		Addr:   serverz.NewAddr("tcp", app.config.GrpcAddr),
		Logger: app.logger,
		Closer: db,
	}
}
