package main

import (
	"fmt"

	"database/sql"

	"github.com/deshboard/boilerplate-model-service/app"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/emperror"
	fxsql "github.com/goph/fxt/database/sql"
	"github.com/goph/fxt/debug"
	fxgrpc "github.com/goph/fxt/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

// ServiceParams provides a set of dependencies for the service constructor.
type ServiceParams struct {
	dig.In

	Connection   *sql.DB
	Logger       log.Logger       `optional:"true"`
	ErrorHandler emperror.Handler `optional:"true"`
}

// NewService constructs a new service instance.
func NewService(params ServiceParams) *app.Service {
	return app.NewService(
		params.Connection,
		app.Logger(params.Logger),
		app.ErrorHandler(params.ErrorHandler),
	)
}

// NewGrpcConfig creates a grpc config.
func NewGrpcConfig(config *Config) *fxgrpc.Config {
	c := fxgrpc.NewConfig(config.GrpcAddr)
	c.ReflectionEnabled = config.GrpcEnableReflection

	return c
}

// NewStreamInterceptor creates a new gRPC server stream interceptor.
func NewStreamInterceptor(tracer opentracing.Tracer) grpc.StreamServerInterceptor {
	return grpc_middleware.ChainStreamServer(
		grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		grpc_prometheus.StreamServerInterceptor,
		grpc_recovery.StreamServerInterceptor(),
	)
}

// NewUnaryInterceptor creates a new gRPC server unary interceptor.
func NewUnaryInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	)
}

// RegisterPrometheusHandler registers the Prometheus metrics handler in the debug server.
func RegisterPrometheusHandler(handler debug.Handler) {
	handler.Handle("/metrics", promhttp.Handler())
}

// NewDatabaseConfig returns a new database connection configuration.
func NewDatabaseConfig(config *Config) *fxsql.Config {
	return fxsql.NewConfig(
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
}
