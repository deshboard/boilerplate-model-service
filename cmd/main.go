package main

import (
	"context"
	"os"

	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror"
	"github.com/goph/fxt"
	"github.com/goph/nest"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

func main() {
	c := new(Context)
	app := fx.New(
		fx.NopLogger,
		fxt.Bootstrap,
		fx.Provide(NewConfig, NewApplicationInfo),
		AppModule,
		fx.Populate(c),
	)

	err := app.Err()
	if dig.RootCause(err) == nest.ErrFlagHelp {
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	// Close resources when the application stops running
	defer c.Close()

	// Register error handler to recover from panics
	defer emperror.HandleRecover(c.ErrorHandler)

	level.Info(c.Logger).Log(
		"msg", "starting",
		"version", Version,
		"commit_hash", CommitHash,
		"build_date", BuildDate,
	)

	err = app.Start(context.Background())
	if err != nil {
		panic(err)
	}

	c.Wait(app)

	ctx, cancel := context.WithTimeout(context.Background(), c.Config.ShutdownTimeout)
	defer cancel()

	err = app.Stop(ctx)
	if err != nil {
		c.ErrorHandler.Handle(err)
	}
}
