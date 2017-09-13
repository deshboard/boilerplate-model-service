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
func newGrpcServer(a *application) serverz.Server {
	serviceChecker := healthz.NewTCPChecker(a.config.GrpcAddr, healthz.WithTCPTimeout(2*time.Second))
	a.healthCollector.RegisterChecker(healthz.LivenessCheck, serviceChecker)

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			a.config.DbUser,
			a.config.DbPass,
			a.config.DbHost,
			a.config.DbPort,
			a.config.DbName,
		),
	)

	if err != nil {
		panic(err)
	}

	a.healthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))

	server := createGrpcServer(a)

	// Register servers here

	grpc_prometheus.Register(server)

	return &serverz.AppServer{
		Server: &grpc.Server{Server: server},
		Name:   "grpc",
		Addr:   serverz.NewAddr("tcp", a.config.GrpcAddr),
		Logger: a.logger,
		Closer: db,
	}
}
