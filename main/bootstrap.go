package main

import (
	"database/sql"
	"fmt"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sagikazarmark/healthz"
	"github.com/sagikazarmark/serverz"
)

func bootstrap() serverz.Server {
	serviceChecker := healthz.NewTCPChecker(config.ServiceAddr, healthz.WithTCPTimeout(2*time.Second))
	checkerCollector.RegisterChecker(healthz.LivenessCheck, serviceChecker)

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.DbUser,
			config.DbPass,
			config.DbHost,
			config.DbPort,
			config.DbName,
		),
	)

	if err != nil {
		logger.Panic(err)
	}

	checkerCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewPingChecker(db))

	server := createGrpcServer()

	// Register servers here

	grpc_prometheus.Register(server)

	return &serverz.NamedServer{
		Server: &serverz.GrpcServer{server},
		Name:   "grpc",
	}
}
