package app

import (
	grpcfx "github.com/deshboard/boilerplate-model-service/pkg/driver/fx"
	"github.com/goph/fxt/app/grpc"
	"github.com/goph/fxt/database/sql"
	"github.com/goph/fxt/tracing/opentracing"
	"go.uber.org/fx"
)

// Module is the collection of all modules of the application.
var Module = fx.Options(
	fxgrpcapp.Module,

	// Configuration
	fx.Provide(
		NewLoggerConfig,
		NewDebugConfig,
	),

	// gRPC server
	fx.Provide(NewGrpcConfig),

	fx.Provide(fxopentracing.NewTracer),

	// Database
	fx.Provide(
		NewDatabaseConfig,
		fxsql.NewConnection,
	),

	grpcfx.Module,
)

type Runner = fxgrpcapp.Runner
